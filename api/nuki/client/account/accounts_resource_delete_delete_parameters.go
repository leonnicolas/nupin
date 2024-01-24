// Code generated by go-swagger; DO NOT EDIT.

package account

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
)

// NewAccountsResourceDeleteDeleteParams creates a new AccountsResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountsResourceDeleteDeleteParams() *AccountsResourceDeleteDeleteParams {
	return &AccountsResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsResourceDeleteDeleteParamsWithTimeout creates a new AccountsResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAccountsResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *AccountsResourceDeleteDeleteParams {
	return &AccountsResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAccountsResourceDeleteDeleteParamsWithContext creates a new AccountsResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAccountsResourceDeleteDeleteParamsWithContext(ctx context.Context) *AccountsResourceDeleteDeleteParams {
	return &AccountsResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAccountsResourceDeleteDeleteParamsWithHTTPClient creates a new AccountsResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountsResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *AccountsResourceDeleteDeleteParams {
	return &AccountsResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
AccountsResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the accounts resource delete delete operation.

	Typically these are written to a http.Request.
*/
type AccountsResourceDeleteDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accounts resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsResourceDeleteDeleteParams) WithDefaults() *AccountsResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accounts resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *AccountsResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) WithContext(ctx context.Context) *AccountsResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *AccountsResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts resource delete delete params
func (o *AccountsResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
