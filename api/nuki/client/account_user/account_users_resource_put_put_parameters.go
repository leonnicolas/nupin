// Code generated by go-swagger; DO NOT EDIT.

package account_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// NewAccountUsersResourcePutPutParams creates a new AccountUsersResourcePutPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountUsersResourcePutPutParams() *AccountUsersResourcePutPutParams {
	return &AccountUsersResourcePutPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountUsersResourcePutPutParamsWithTimeout creates a new AccountUsersResourcePutPutParams object
// with the ability to set a timeout on a request.
func NewAccountUsersResourcePutPutParamsWithTimeout(timeout time.Duration) *AccountUsersResourcePutPutParams {
	return &AccountUsersResourcePutPutParams{
		timeout: timeout,
	}
}

// NewAccountUsersResourcePutPutParamsWithContext creates a new AccountUsersResourcePutPutParams object
// with the ability to set a context for a request.
func NewAccountUsersResourcePutPutParamsWithContext(ctx context.Context) *AccountUsersResourcePutPutParams {
	return &AccountUsersResourcePutPutParams{
		Context: ctx,
	}
}

// NewAccountUsersResourcePutPutParamsWithHTTPClient creates a new AccountUsersResourcePutPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountUsersResourcePutPutParamsWithHTTPClient(client *http.Client) *AccountUsersResourcePutPutParams {
	return &AccountUsersResourcePutPutParams{
		HTTPClient: client,
	}
}

/*
AccountUsersResourcePutPutParams contains all the parameters to send to the API endpoint

	for the account users resource put put operation.

	Typically these are written to a http.Request.
*/
type AccountUsersResourcePutPutParams struct {

	/* Body.

	   Account sub create representation
	*/
	Body *models.AccountUserCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account users resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUsersResourcePutPutParams) WithDefaults() *AccountUsersResourcePutPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account users resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUsersResourcePutPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) WithTimeout(timeout time.Duration) *AccountUsersResourcePutPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) WithContext(ctx context.Context) *AccountUsersResourcePutPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) WithHTTPClient(client *http.Client) *AccountUsersResourcePutPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) WithBody(body *models.AccountUserCreate) *AccountUsersResourcePutPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the account users resource put put params
func (o *AccountUsersResourcePutPutParams) SetBody(body *models.AccountUserCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AccountUsersResourcePutPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
