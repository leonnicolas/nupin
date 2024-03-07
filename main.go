package main

import (
	"context"
	"crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"slices"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/metalmatze/signal/internalserver"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	flag "github.com/spf13/pflag"
	passwdv "github.com/wagslane/go-password-validator"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"

	nukiclient "github.com/leonnicolas/nupin/api/nuki/client"
	"github.com/leonnicolas/nupin/api/nuki/client/account_user"
	"github.com/leonnicolas/nupin/api/nuki/client/smartlock_auth"
	"github.com/leonnicolas/nupin/api/nuki/models"
	v0 "github.com/leonnicolas/nupin/api/v0"
	"github.com/leonnicolas/nupin/config"
	"github.com/leonnicolas/nupin/store"
)

//go:embed fe/build
var public embed.FS

type Rand interface {
	Next() string
}

type randomGenerator struct{}

func (r *randomGenerator) Next() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

type claims struct {
	Email    string `json:"email"`
	Verified bool   `json:"email_verified"`
	Name     string `json:"name,omitempty"`
}

func redirectHandler(oauth2Config *oauth2.Config, s store.Store, randG Rand) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := randG.Next()
		s.Set(str)
		http.Redirect(w, r, oauth2Config.AuthCodeURL(str), http.StatusFound)
	}
}

func callbackHandler(oauth2Config *oauth2.Config, cfg *config.Config, verifier *oidc.IDTokenVerifier, s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		if state == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no state")

			return
		}
		if ok, err := s.Exist(state); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to check state: %s", err.Error())
			return
		} else if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "invalid state")

			return
		}
		if err := s.Delete(state); err != nil {
			log.Println("failed to delete state:", err.Error())
		}

		oauth2Token, err := oauth2Config.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "token exchange failed: %s", err.Error())
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to extract id token: %s", err.Error())

			return
		}

		idToken, err := verifier.Verify(r.Context(), rawIDToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "failed to verify token: %s", err.Error())

			return
		}
		if !slices.Contains(idToken.Audience, cfg.Oidc.ClientID) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "wrong Audience: %s", err.Error())

			return
		}

		var claims claims
		if err := idToken.Claims(&claims); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "oauth callback failed: %s", err.Error())

			return
		}

		if idToken.Expiry.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "token expired")

			return
		}

		http.SetCookie(w, &http.Cookie{
			SameSite: http.SameSiteStrictMode,
			Name:     "jwt",
			Value:    rawIDToken,
			Domain:   cfg.Oidc.CookieDomain,
			HttpOnly: true,
			Path:     "/",
		})

		// Firefox and Safari don't set the jwt cookie when using SameSite=strict right after the redirect.
		// Resulting in an endless sign in loop. This can be fixed with SameSite=lax (not the same security)
		// or using client side redirect. See below!
		fmt.Fprint(w,
			`<!DOCTYPE html>
			<html>
				<head><meta http-equiv="refresh" content="0; url='/'"></head>
				<body></body>
			</html>`)

		return
	}
}

func redirectOrError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	if header := r.Header.Get("Accept"); header == "application/json" {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(v0.ErrorResponse{
			Error: msg,
		})
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func authMiddleware(next http.HandlerFunc, verifier *oidc.IDTokenVerifier, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			log.Println("no cookie")
			redirectOrError(w, r, http.StatusUnauthorized, "missing cookie")

			return
		}

		idToken, err := verifier.Verify(r.Context(), c.Value)
		if err != nil {
			redirectOrError(w, r, http.StatusUnauthorized, fmt.Sprintf("failed to verify token: %s", err.Error()))

			return
		}

		if !slices.Contains(idToken.Audience, cfg.Oidc.ClientID) {
			redirectOrError(w, r, http.StatusUnauthorized, fmt.Sprintf("wrong Audience: %s", err.Error()))

			return
		}

		if idToken.Expiry.Before(time.Now()) {
			redirectOrError(w, r, http.StatusUnauthorized, fmt.Sprintf("session expired"))

			return
		}

		var claims claims
		if err := idToken.Claims(&claims); err != nil {
			redirectOrError(w, r, http.StatusBadRequest, fmt.Sprintf("oauth callback failed: %s", err.Error()))

			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
		r = r.WithContext(context.WithValue(r.Context(), "idToken", idToken))

		next(w, r)
	}
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Context().Value("idToken").(*oidc.IDToken)

		fmt.Fprintf(w, "hello loged in for %v", idToken.Expiry.Sub(time.Now()).String())
	}
}

var ConfigPath string

func init() {
	flag.StringVarP(&ConfigPath, "config", "c", "", "path to Config file")
}

type server struct {
	c   *nukiclient.NukiAPI
	cfg *config.Config
}

func strPtr(s string) *string {
	return &s
}

func ptr[E any](e E) *E {
	return &e
}

func (s *server) GetUser(ctx context.Context, request v0.GetUserRequestObject) (v0.GetUserResponseObject, error) {
	c, ok := ctx.Value("claims").(claims)
	if !ok {
		return v0.GetUserdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpeced error"),
				Error:          "failed to get claims",
			},
		}, nil
	}

	u := v0.User{
		Email: c.Email,
		Name:  c.Name,
	}

	return &v0.GetUser200JSONResponse{UserGetResponseJSONResponse: v0.UserGetResponseJSONResponse(u)}, nil
}

func (s *server) UpdatePin(ctx context.Context, request v0.UpdatePinRequestObject) (v0.UpdatePinResponseObject, error) {
	c, ok := ctx.Value("claims").(claims)
	if !ok {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpeced error"),
				Error:          "failed to get claims",
			},
		}, nil
	}

	if c.Email == "" {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpeced error"),
				Error:          "missing email",
			},
		}, nil
	}
	if e := passwdv.GetEntropy(fmt.Sprintf("%d", request.Body.Code)); e < s.cfg.Nuki.MinimumPinEntropy {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: v0.ApiError{
				DisplayMessage: ptr("Please select a stronger Pin"),
				Error:          fmt.Sprintf("code's entropy is too low (%f < %f)", e, s.cfg.Nuki.MinimumPinEntropy),
			},
		}, nil
	}

	usersResponse, err := s.c.AccountUser.AccountUsersResourceGetGet(
		&account_user.AccountUsersResourceGetGetParams{
			Context: ctx,
			Email:   &c.Email,
		},
		httptransport.BearerToken(s.cfg.Nuki.APIKey),
	)
	if err != nil {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpeced error"),
				Error:          err.Error(),
			},
		}, nil
	}
	if len(usersResponse.Payload) > 1 {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpected number of accounts found"),
				Error:          "unexpected number of accounts found",
			},
		}, nil
	}
	var user *models.AccountUser
	if len(usersResponse.Payload) == 0 {
		log.Println("creating account user", "email", c.Email, "name", c.Name)
		createResponse, err := s.c.AccountUser.AccountUsersResourcePutPut(
			&account_user.AccountUsersResourcePutPutParams{
				Context: ctx,
				Body: &models.AccountUserCreate{
					Email: strPtr(c.Email),
					Name:  strPtr(c.Name),
				},
			},
			httptransport.BearerToken(s.cfg.Nuki.APIKey),
		)
		if err != nil {
			return nil, err
		}
		user = createResponse.Payload
	} else {
		user = usersResponse.Payload[0]
	}

	opt := func(opt *runtime.ClientOperation) {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.DisableKeepAlives = true
		httpClient := &http.Client{Transport: t}
		opt.Client = httpClient
	}

	authRes, err := s.c.SmartlockAuth.SmartlockAuthsResourceGetGet(
		&smartlock_auth.SmartlockAuthsResourceGetGetParams{
			Context: ctx, Types: ptr("13"),
			SmartlockID: s.cfg.Nuki.SmartLockDevice,
		},
		httptransport.BearerToken(s.cfg.Nuki.APIKey),
		opt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get auths %w", err)
	}
	index := slices.IndexFunc(authRes.Payload, func(a *models.SmartlockAuth) bool {
		return a.AccountUserID == *user.AccountUserID
	})
	if index == -1 {
		log.Println("creating auth user", "accountID", *user.AccountUserID, "name", c.Name)
		_, err := s.c.SmartlockAuth.SmartlockAuthsResourcePutPut(
			&smartlock_auth.SmartlockAuthsResourcePutPutParams{
				SmartlockID: s.cfg.Nuki.SmartLockDevice,
				Context:     ctx,
				Body: &models.SmartlockAuthCreate{
					AccountUserID:       *user.AccountUserID,
					SmartActionsEnabled: false,
					Name:                strPtr(c.Name),
					Type:                13,
					Code:                int32(request.Body.Code),
				},
			},
			httptransport.BearerToken(s.cfg.Nuki.APIKey),
		)
		if err != nil {
			log.Println("failed to create auth user", "accountID", *user.AccountUserID, "name", c.Name, err.Error())
			return v0.UpdatePindefaultJSONResponse{
				StatusCode: http.StatusBadRequest,
				Body: v0.ApiError{
					DisplayMessage: ptr("please try another Pin"),
					Error:          err.Error(),
				},
			}, nil
		}

		return &v0.UpdatePin200Response{}, nil
	}

	log.Println("found smartlock auth", *authRes.Payload[index].ID, "claim name", c.Email, "auth ID", *user.AccountUserID, "auth name", *authRes.Payload[index].Name)

	_, err = s.c.SmartlockAuth.SmartlockAuthResourcePostPost(
		&smartlock_auth.SmartlockAuthResourcePostPostParams{
			SmartlockID: s.cfg.Nuki.SmartLockDevice,
			Context:     ctx,
			Body: &models.SmartlockAuthUpdate{
				Code:          int32(request.Body.Code),
				Enabled:       true,
				AccountUserID: *user.AccountUserID,
			},
			ID: *authRes.Payload[index].ID,
		},
		httptransport.BearerToken(s.cfg.Nuki.APIKey),
	)
	if err != nil {
		return v0.UpdatePindefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: v0.ApiError{
				DisplayMessage: ptr("please try another Pin"),
				Error:          err.Error(),
			},
		}, nil
	}

	log.Println("updated smartlock auth", *authRes.Payload[index].ID, "claim name", c.Email, "auth name", *authRes.Payload[index].Name)

	return &v0.UpdatePin200Response{}, nil
}

var _ v0.StrictServerInterface = &server{}

func main() {
	flag.Parse()

	ctx := context.Background()

	var cfg config.Config
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

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	client := nukiclient.NewHTTPClientWithConfig(strfmt.Default, &nukiclient.TransportConfig{
		Host:    "api.nuki.io",
		Schemes: []string{"https"},
	})

	client.AccountUser = account_user.NewInstrumentedClientService(client.AccountUser, prometheus.WrapRegistererWith(prometheus.Labels{"api": "account_user"}, reg))
	client.SmartlockAuth = smartlock_auth.NewInstrumentedClientService(client.SmartlockAuth, prometheus.WrapRegistererWith(prometheus.Labels{"api": "smartlock_auth"}, reg))

	serverImpl := &server{c: client, cfg: &cfg}

	mux := chi.NewRouter()
	rG := &randomGenerator{}

	var s store.Store
	switch cfg.StorageType {
	case "memory", "":
		s = store.NewMemoryStore()
	case "redis":
		s = store.NewRedisStore(&cfg)
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
	mux.HandleFunc("/auth/callback", callbackHandler(&oauth2Config, &cfg, verifier, s))
	v0Handler := v0.HandlerWithOptions(v0.NewInstrumentedServerInterface(v0.NewStrictHandlerWithOptions(serverImpl, nil,
		v0.StrictHTTPServerOptions{
			RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println("request error", err.Error())
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(v0.ErrorResponse{
					Error:          err.Error(),
					DisplayMessage: ptr("unexpected error"),
				})
			},
			ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println("response error", err.Error())
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(v0.ErrorResponse{
					Error:          err.Error(),
					DisplayMessage: ptr("unexpected error"),
				})
			},
		}), reg),
		v0.ChiServerOptions{
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println("chi error handler", err.Error())
				return
			},
		})
	mux.Mount("/api",
		authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				v0Handler.ServeHTTP(w, r)
			},
			verifier, &cfg))

	if v, ok := os.LookupEnv("NUPIN_DEVELOPMENT"); ok && v == "true" {
		u, err := url.Parse("http://localhost:3000")
		if err != nil {
			panic(err)
		}

		log.Printf("proxy /* to %s\n", u.String())
		mux.Handle("/", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
			}, verifier, &cfg))
		mux.Handle("/*", httputil.NewSingleHostReverseProxy(u))
	} else {
		dir, err := fs.Sub(public, "fe/build")
		if err != nil {
			panic(err)
		}
		mux.HandleFunc("/", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				http.FileServer(http.FS(dir)).ServeHTTP(w, r)
			}, verifier, &cfg))
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
