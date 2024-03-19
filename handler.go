package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"

	v0 "github.com/leonnicolas/nupin/api/v0"
	"github.com/leonnicolas/nupin/config"
	"github.com/leonnicolas/nupin/store"
)

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
