package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/strfmt"
	"github.com/metalmatze/signal/internalserver"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	flag "github.com/spf13/pflag"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"

	nukiclient "github.com/leonnicolas/nupin/api/nuki/client"
	"github.com/leonnicolas/nupin/api/nuki/client/account_user"
	"github.com/leonnicolas/nupin/api/nuki/client/smartlock_auth"
	"github.com/leonnicolas/nupin/config"
	"github.com/leonnicolas/nupin/store"
)

//go:embed fe/build
var public embed.FS

var ConfigPath string

func init() {
	flag.StringVarP(&ConfigPath, "config", "c", "", "path to Config file")
	flag.BoolP("help", "h", false, "print help text and exit")
}

func main() {
	flag.Parse()

	if help, _ := flag.CommandLine.GetBool("help"); help {
		flag.Usage()
		return
	}

	ctx := context.Background()

	cfg := &config.Config{}
	cfg.Web.Address = ":9999"
	cfg.Nuki.MinimumPinEntropy = 10
	cfg.Web.InternalAddress = ":9090"

	if ConfigPath != "" {
		log.Println("reading Config from", ConfigPath)
		if f, err := os.Open(ConfigPath); err != nil {
			log.Fatal(err)
		} else {
			if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
				log.Fatal(err)
			}
		}
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	provider, err := oidc.NewProvider(ctx, cfg.Oidc.IssuerURL)
	if err != nil {
		log.Fatal(err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.Oidc.ClientID})

	oauth2Config := oauth2.Config{
		ClientID:     cfg.Oidc.ClientID,
		ClientSecret: cfg.Oidc.ClientSecret,
		RedirectURL:  cfg.Oidc.RedirectUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	client := nukiclient.NewHTTPClientWithConfig(strfmt.Default, &nukiclient.TransportConfig{
		Host:    "api.nuki.io",
		Schemes: []string{"https"},
	})

	client.AccountUser = account_user.NewInstrumentedClientService(client.AccountUser, prometheus.WrapRegistererWith(prometheus.Labels{"api": "account_user"}, reg))
	client.SmartlockAuth = smartlock_auth.NewInstrumentedClientService(client.SmartlockAuth, prometheus.WrapRegistererWith(prometheus.Labels{"api": "smartlock_auth"}, reg))

	serverImpl := &server{c: client, cfg: cfg}

	mux := chi.NewRouter()
	rG := &randomGenerator{}

	var s store.Store
	switch cfg.StorageType {
	case "memory", "":
		s = store.NewMemoryStore()
	case "redis":
		s = store.NewRedisStore(cfg)
	default:
		log.Fatal("unknown storage type")
	}

	mux.HandleFunc("/login", redirectHandler(&oauth2Config, s, rG))
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			SameSite: http.SameSiteStrictMode,
			Name:     "jwt",
			Domain:   cfg.Oidc.CookieDomain,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Time{},
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})
	mux.HandleFunc("/auth/callback", callbackHandler(&oauth2Config, cfg, verifier, s))
	v0Handler := v0Handler(serverImpl, reg)
	mux.Mount("/api",
		authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				v0Handler.ServeHTTP(w, r)
			},
			verifier, cfg))

	if v, ok := os.LookupEnv("NUPIN_DEVELOPMENT"); ok && v == "true" {
		u, err := url.Parse("http://localhost:3000")
		if err != nil {
			panic(err)
		}

		log.Printf("proxy /* to %s\n", u.String())
		mux.Handle("/", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
			}, verifier, cfg))
		mux.Handle("/*", httputil.NewSingleHostReverseProxy(u))
	} else {
		dir, err := fs.Sub(public, "fe/build")
		if err != nil {
			panic(err)
		}
		mux.HandleFunc("/", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				http.FileServer(http.FS(dir)).ServeHTTP(w, r)
			}, verifier, cfg))
		mux.Handle("/*", http.FileServer(http.FS(dir)))
	}

	g := run.Group{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g.Add(run.SignalHandler(ctx, os.Interrupt))
	{
		l, err := net.Listen("tcp", cfg.Web.Address)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			return http.Serve(l, mux)
		},
			func(err error) {
				l.Close()
			})
	}
	{
		mux := http.NewServeMux()
		mux.Handle("/", internalserver.NewHandler(
			internalserver.WithName("internal"),
			internalserver.WithPrometheusRegistry(reg),
			internalserver.WithPProf(),
		))
		l, err := net.Listen("tcp", cfg.Web.InternalAddress)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			return http.Serve(l, mux)
		},
			func(err error) {
				l.Close()
			})
	}
	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}

func ptr[E any](e E) *E {
	return &e
}
