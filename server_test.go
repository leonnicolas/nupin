package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	nukiclient "github.com/leonnicolas/nupin/api/nuki/client"
	"github.com/leonnicolas/nupin/api/nuki/client/account_user"
	"github.com/leonnicolas/nupin/api/nuki/client/smartlock_auth"
	"github.com/leonnicolas/nupin/api/nuki/models"
	v0 "github.com/leonnicolas/nupin/api/v0"
	"github.com/leonnicolas/nupin/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/mock"
)

//go:generate  go run github.com/vektra/mockery/v2  --name ClientService --all=false --dir api/nuki/client/account_user --inpackage
//go:generate  go run github.com/vektra/mockery/v2  --name ClientService --all=false --dir api/nuki/client/smartlock_auth --inpackage

func TestUpdatePin(t *testing.T) {
	cfg := &config.Config{
		Nuki: struct {
			APIKey          string `yaml:"apiKey"`
			SmartLockDevice int64  `yaml:"smartLockDevice"`
			// MinimumPinEntropy is the mimimum entropy needed by a pin to be accepted (default: 10)
			MinimumPinEntropy  float64 `yaml:"minimumPinEntropy,omitempty"`
			AllowMonotonicPins bool    `yaml:"allowMonotonicPins,omitempty"`
			AllowedFromTime *int `yaml:"allowedFromTime,omitempty"`
			AllowedUntilTime *int `yaml:"allowedUntilTime,omitempty"`
			AllowedWeekDays int `yaml:"allowedWeekDays,omitempty"`
		}{
			MinimumPinEntropy: 10,
		},
	}

	for _, scenario := range []struct {
		name   string
		claims claims
		code   int
		before func(*testing.T, *account_user.MockClientService, *smartlock_auth.MockClientService)
		after  func(*testing.T, *http.Response)
	}{
		{
			name: "no email",
			claims: claims{
				Verified: true,
				Name:     "Harry Potter",
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusInternalServerError {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if getResponse.Error != "missing email" {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "weak pin",
			code: 123456,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusBadRequest {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "code's entropy is too low") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "AccountUsersResourceGetGet networkerror",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("network error"))
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusInternalServerError {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "network error") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "AccountUsersResourceGetGet more than one account found",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{Payload: []*models.AccountUser{{}, {}}}, nil)
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusInternalServerError {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "unexpected number of accounts found") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "account not found, creation fails",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{Payload: []*models.AccountUser{}}, nil)
				au.On("AccountUsersResourcePutPut", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("network error"))
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusInternalServerError {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "network error") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "account not found, creation succeeds, smartlockAuth not found, creation succeeds",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{Payload: []*models.AccountUser{}}, nil)
				au.On("AccountUsersResourcePutPut", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourcePutPutOK{
					Payload: &models.AccountUser{
						AccountUserID: ptr(int32(1)),
						Name:          ptr("no 1"),
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(&smartlock_auth.SmartlockAuthsResourceGetGetOK{
					Payload: []*models.SmartlockAuth{
						{
							ID:            ptr("2"),
							AccountUserID: 2,
							Name:          ptr("no 2"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourcePutPut", mock.Anything, mock.Anything).Return(nil, nil)
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusOK {
					t.Errorf("unexpected status code: %d", code)
				}
			},
		},
		{
			name: "account exists, SmartlockAuthsResourceGetGet error",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{
					Payload: []*models.AccountUser{
						{
							AccountUserID: ptr(int32(1)),
							Email:         ptr("harry.potter@hogwards.magic"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("network error"))
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusInternalServerError {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "network error") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "account exists, SmartlockAuths exists, updating fails",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{
					Payload: []*models.AccountUser{
						{
							AccountUserID: ptr(int32(1)),
							Email:         ptr("harry.potter@hogwards.magic"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(&smartlock_auth.SmartlockAuthsResourceGetGetOK{
					Payload: []*models.SmartlockAuth{
						{
							ID:            ptr("1"),
							AccountUserID: 1,
							Name:          ptr("no 1"),
						},
						{
							ID:            ptr("2"),
							AccountUserID: 2,
							Name:          ptr("no 2"),
						},
					},
				}, nil)
				// At the moment we can't tell from the client response if the pin is invalid or something else went wrong.
				// The most likely error is an invalid pin, so we take a guess here and return a 400.
				sa.On("SmartlockAuthResourcePostPost", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("bad pin"))
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusBadRequest {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "bad pin") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "account exists, SmartlockAuths exists, updating succeeds",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{
					Payload: []*models.AccountUser{
						{
							AccountUserID: ptr(int32(1)),
							Email:         ptr("harry.potter@hogwards.magic"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(&smartlock_auth.SmartlockAuthsResourceGetGetOK{
					Payload: []*models.SmartlockAuth{
						{
							ID:            ptr("1"),
							AccountUserID: 1,
							Name:          ptr("no 1"),
						},
						{
							ID:            ptr("2"),
							AccountUserID: 2,
							Name:          ptr("no 2"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthResourcePostPost", mock.Anything, mock.Anything).Return(nil, nil)
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusOK {
					t.Errorf("unexpected status code: %d", code)
				}
			},
		},
		{
			name: "account exists, SmartlockAuths does not, creation fails",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{
					Payload: []*models.AccountUser{
						{
							AccountUserID: ptr(int32(1)),
							Email:         ptr("harry.potter@hogwards.magic"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(&smartlock_auth.SmartlockAuthsResourceGetGetOK{
					Payload: []*models.SmartlockAuth{
						{
							ID:            ptr("2"),
							AccountUserID: 2,
							Name:          ptr("no 2"),
						},
					},
				}, nil)
				// At the moment we can't tell from the client response if the pin is invalid or something else went wrong.
				// The most likely error is an invalid pin, so we take a guess here and return a 400.
				sa.On("SmartlockAuthsResourcePutPut", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("bad pin"))
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusBadRequest {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.ErrorResponse
				jsonDecode(t, r.Body, &getResponse)
				if !strings.Contains(getResponse.Error, "bad pin") {
					t.Error("error mismatch")
				}
			},
		},
		{
			name: "account exists, SmartlockAuths does not, creation succeeds",
			code: 520313,
			claims: claims{
				Verified: true,
				Email:    "harry.potter@hogwards.magic",
				Name:     "Harry Potter",
			},
			before: func(t *testing.T, au *account_user.MockClientService, sa *smartlock_auth.MockClientService) {
				au.On("AccountUsersResourceGetGet", mock.Anything, mock.Anything).Return(&account_user.AccountUsersResourceGetGetOK{
					Payload: []*models.AccountUser{
						{
							AccountUserID: ptr(int32(1)),
							Email:         ptr("harry.potter@hogwards.magic"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourceGetGet", mock.Anything, mock.Anything).Return(&smartlock_auth.SmartlockAuthsResourceGetGetOK{
					Payload: []*models.SmartlockAuth{
						{
							ID:            ptr("2"),
							AccountUserID: 2,
							Name:          ptr("no 2"),
						},
					},
				}, nil)
				sa.On("SmartlockAuthsResourcePutPut", mock.Anything, mock.Anything).Return(nil, nil)
			},
			after: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusOK {
					t.Errorf("unexpected status code: %d", code)
				}
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			au := account_user.NewMockClientService(t)
			sa := smartlock_auth.NewMockClientService(t)
			nukiAPI := &nukiclient.NukiAPI{
				AccountUser:   au,
				SmartlockAuth: sa,
			}

			s := &server{
				cfg: cfg,
				c:   nukiAPI,
			}

			h := v0Handler(s, prometheus.NewPedanticRegistry())

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/v0/code", jsonEncode(t, v0.UpdatePinJSONRequestBody{Code: scenario.code}))
			req = req.WithContext(context.WithValue(req.Context(), "claims", scenario.claims))

			if scenario.before != nil {
				scenario.before(t, au, sa)
			}

			h.ServeHTTP(recorder, req)
			result := recorder.Result()
			buf := bytes.NewBuffer(nil)
			result.Body = io.NopCloser(io.TeeReader(result.Body, buf))
			scenario.after(t, result)
			if t.Failed() {
				t.Logf("body=%q\n", buf.String())
			}
		})
	}
}

func jsonDecode(t *testing.T, r io.Reader, v any) {
	t.Helper()

	if err := json.NewDecoder(r).Decode(v); err != nil {
		t.Error(err.Error())
	}
}

func jsonEncode(t *testing.T, v any) io.Reader {
	t.Helper()

	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		t.Errorf("failed to encode to json: %s", err.Error())
	}
	return buf
}

func TestGetUser(t *testing.T) {
	cfg := &config.Config{}
	nukiAPI := &nukiclient.NukiAPI{}

	s := &server{
		cfg: cfg,
		c:   nukiAPI,
	}

	si := v0.NewStrictHandler(s, nil)
	h := v0.Handler(si)

	for _, scenario := range []struct {
		name     string
		claims   claims
		expecter func(*testing.T, *http.Response)
	}{
		{
			name: "happy path",
			claims: claims{
				Email:    "harry.potter@hogwards.magic",
				Verified: true,
				Name:     "Harry Potter",
			},
			expecter: func(t *testing.T, r *http.Response) {
				if code := r.StatusCode; code != http.StatusOK {
					t.Errorf("unexpected status code: %d", code)
				}
				var getResponse v0.GetUser200JSONResponse
				if err := json.NewDecoder(r.Body).Decode(&getResponse); err != nil {
					t.Error(err)
				}
				if getResponse.Email != "harry.potter@hogwards.magic" {
					t.Error("email mismatch")
				}
				if getResponse.Name != "Harry Potter" {
					t.Error("name mismatch")
				}
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/v0/user", nil)
			req = req.WithContext(context.WithValue(req.Context(), "claims", scenario.claims))
			h.ServeHTTP(recorder, req)
			scenario.expecter(t, recorder.Result())
		})
	}

	t.Run("missing claims", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/v0/user", nil)
		h.ServeHTTP(recorder, req)
		result := recorder.Result()
		if code := result.StatusCode; code != http.StatusInternalServerError {
			t.Errorf("unexpected status code: %d", code)
		}
		var getResponse v0.ErrorResponseJSONResponse
		if err := json.NewDecoder(result.Body).Decode(&getResponse); err != nil {
			t.Error(err)
		}
		if *getResponse.DisplayMessage != "unexpected error" {
			t.Error("display message mismatch")
		}
		if getResponse.Error != "failed to get claims" {
			t.Error("error mismatch")
		}
	})
}
