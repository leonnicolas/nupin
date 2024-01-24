// Code generated by go-swagger; DO NOT EDIT.

package account_subscription

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

// NewAccountSubscriptionPayResourcePostPostParams creates a new AccountSubscriptionPayResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountSubscriptionPayResourcePostPostParams() *AccountSubscriptionPayResourcePostPostParams {
	return &AccountSubscriptionPayResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountSubscriptionPayResourcePostPostParamsWithTimeout creates a new AccountSubscriptionPayResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAccountSubscriptionPayResourcePostPostParamsWithTimeout(timeout time.Duration) *AccountSubscriptionPayResourcePostPostParams {
	return &AccountSubscriptionPayResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAccountSubscriptionPayResourcePostPostParamsWithContext creates a new AccountSubscriptionPayResourcePostPostParams object
// with the ability to set a context for a request.
func NewAccountSubscriptionPayResourcePostPostParamsWithContext(ctx context.Context) *AccountSubscriptionPayResourcePostPostParams {
	return &AccountSubscriptionPayResourcePostPostParams{
		Context: ctx,
	}
}

// NewAccountSubscriptionPayResourcePostPostParamsWithHTTPClient creates a new AccountSubscriptionPayResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountSubscriptionPayResourcePostPostParamsWithHTTPClient(client *http.Client) *AccountSubscriptionPayResourcePostPostParams {
	return &AccountSubscriptionPayResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
AccountSubscriptionPayResourcePostPostParams contains all the parameters to send to the API endpoint

	for the account subscription pay resource post post operation.

	Typically these are written to a http.Request.
*/
type AccountSubscriptionPayResourcePostPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account subscription pay resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountSubscriptionPayResourcePostPostParams) WithDefaults() *AccountSubscriptionPayResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account subscription pay resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountSubscriptionPayResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) WithTimeout(timeout time.Duration) *AccountSubscriptionPayResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) WithContext(ctx context.Context) *AccountSubscriptionPayResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) WithHTTPClient(client *http.Client) *AccountSubscriptionPayResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account subscription pay resource post post params
func (o *AccountSubscriptionPayResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AccountSubscriptionPayResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
