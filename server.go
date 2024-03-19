package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/prometheus/client_golang/prometheus"
	passwdv "github.com/wagslane/go-password-validator"

	nukiclient "github.com/leonnicolas/nupin/api/nuki/client"
	"github.com/leonnicolas/nupin/api/nuki/client/account_user"
	"github.com/leonnicolas/nupin/api/nuki/client/smartlock_auth"
	"github.com/leonnicolas/nupin/api/nuki/models"
	v0 "github.com/leonnicolas/nupin/api/v0"
	"github.com/leonnicolas/nupin/config"
)

type server struct {
	c   *nukiclient.NukiAPI
	cfg *config.Config
}

func v0Handler(serverImpl v0.StrictServerInterface, reg prometheus.Registerer) http.Handler {
	return v0.HandlerWithOptions(v0.NewInstrumentedServerInterface(v0.NewStrictHandlerWithOptions(serverImpl, nil,
		v0.StrictHTTPServerOptions{
			RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println("request error:", err.Error())
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(v0.ErrorResponse{
					Error:          err.Error(),
					DisplayMessage: ptr("unexpected error"),
				})
			},
			ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
				log.Println("response error:", err.Error())
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
				log.Println("chi error handler:", err.Error())
				return
			},
		})
}

func (s *server) GetUser(ctx context.Context, request v0.GetUserRequestObject) (v0.GetUserResponseObject, error) {
	c, ok := ctx.Value("claims").(claims)
	if !ok {
		return v0.GetUserdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: v0.ApiError{
				DisplayMessage: ptr("unexpected error"),
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
					Email: ptr(c.Email),
					Name:  ptr(c.Name),
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
		return nil, fmt.Errorf("failed to get auths: %w", err)
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
					Name:                ptr(c.Name),
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
