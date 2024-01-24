package main

import (
	"context"
	"crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"slices"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-openapi/runtime/client"
	flag "github.com/spf13/pflag"
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
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}
}

func authMiddleware(next http.HandlerFunc, verifier *oidc.IDTokenVerifier, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			fmt.Println("no cookie")
			http.Redirect(w, r, "/login", http.StatusFound)

			return
		}

		idToken, err := verifier.Verify(r.Context(), c.Value)
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

		if idToken.Expiry.Before(time.Now()) {
			http.Redirect(w, r, "/login", http.StatusFound)

			return
		}

		var claims claims
		if err := idToken.Claims(&claims); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "oauth callback failed: %s", err.Error())

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
	flag.StringVarP(&ConfigPath, "Config", "c", "", "path to Config file")
}

const appName = "nupin"

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

	authRes, err := s.c.SmartlockAuth.SmartlockAuthsResourceGetGet(
		&smartlock_auth.SmartlockAuthsResourceGetGetParams{
			Context: ctx, Types: ptr("13"),
			SmartlockID: s.cfg.Nuki.SmartLockDevice,
		},
		httptransport.BearerToken(s.cfg.Nuki.APIKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get auths %w", err)
	}
	index := slices.IndexFunc(authRes.Payload, func(a *models.SmartlockAuth) bool {
		return a.Name != nil && *a.Name == c.Email
	})
	if index == -1 {
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
			return v0.UpdatePindefaultJSONResponse{
				StatusCode: http.StatusInternalServerError,
				Body: v0.ApiError{
					DisplayMessage: ptr("please try another Pin"),
					Error:          err.Error(),
				},
			}, nil
		}

		return &v0.UpdatePin200Response{}, nil
	}

	log.Println("found smartlock auth", *authRes.Payload[index].ID, "claim name", c.Email, "auth name", *authRes.Payload[index].Name)

	_, err = s.c.SmartlockAuth.SmartlockAuthResourcePostPost(
		&smartlock_auth.SmartlockAuthResourcePostPostParams{
			SmartlockID: s.cfg.Nuki.SmartLockDevice,
			Context:     ctx,
			Body: &models.SmartlockAuthUpdate{
				Code: int32(request.Body.Code),
			},
			ID: *authRes.Payload[index].ID,
		},
		httptransport.BearerToken(s.cfg.Nuki.APIKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update pin %w", err)
	}

	return &v0.UpdatePin200Response{}, nil
}

var _ v0.StrictServerInterface = &server{}

func main() {
	flag.Parse()

	ctx := context.Background()

	var cfg config.Config
	cfg.Web.Address = ":9999"

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

	u, err := url.Parse("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	client := nukiclient.NewHTTPClientWithConfig(nil, &nukiclient.TransportConfig{
		Host:    "api.nuki.io",
		Schemes: []string{"https"},
	})

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
	mux.HandleFunc("/auth/callback", callbackHandler(&oauth2Config, &cfg, verifier, s))
	mux.HandleFunc("/a", authMiddleware(homeHandler(), verifier, &cfg))
	mux.Mount("/api",
		authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				v0.HandlerWithOptions(v0.NewStrictHandlerWithOptions(&server{c: client, cfg: &cfg}, nil,
					v0.StrictHTTPServerOptions{
						RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
							w.Header().Set("Content-Type", "application/json")
							w.WriteHeader(http.StatusUnprocessableEntity)
							json.NewEncoder(w).Encode(v0.ErrorResponse{
								Error:          err.Error(),
								DisplayMessage: ptr("unexpected error"),
							})
						},
						ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
							w.Header().Set("Content-Type", "application/json")
							w.WriteHeader(http.StatusInternalServerError)
							json.NewEncoder(w).Encode(v0.ErrorResponse{
								Error:          err.Error(),
								DisplayMessage: ptr("unexpected error"),
							})
						},
					}),
					v0.ChiServerOptions{
						ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
							log.Println("chi error handler", err.Error())
							return
						},
					}).ServeHTTP(w, r)
			},
			verifier, &cfg))
	if v, ok := os.LookupEnv("NUPIN_DEVELOPMENT"); ok && v == "true" {
		log.Println("proxy /* to localhost:3000")
		mux.Handle("/*", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
			}, verifier, &cfg))
	} else {
		dir, err := fs.Sub(public, "fe/build")
		if err != nil {
			panic(err)
		}
		mux.Handle("/*", authMiddleware(
			func(w http.ResponseWriter, r *http.Request) {
				http.FileServer(http.FS(dir)).ServeHTTP(w, r)
			}, verifier, &cfg))
	}
	http.ListenAndServe(cfg.Web.Address, mux)
}