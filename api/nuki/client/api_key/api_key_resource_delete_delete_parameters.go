// Code generated by go-swagger; DO NOT EDIT.

package api_key

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
	"github.com/go-openapi/swag"
)

// NewAPIKeyResourceDeleteDeleteParams creates a new APIKeyResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyResourceDeleteDeleteParams() *APIKeyResourceDeleteDeleteParams {
	return &APIKeyResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyResourceDeleteDeleteParamsWithTimeout creates a new APIKeyResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAPIKeyResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *APIKeyResourceDeleteDeleteParams {
	return &APIKeyResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAPIKeyResourceDeleteDeleteParamsWithContext creates a new APIKeyResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAPIKeyResourceDeleteDeleteParamsWithContext(ctx context.Context) *APIKeyResourceDeleteDeleteParams {
	return &APIKeyResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAPIKeyResourceDeleteDeleteParamsWithHTTPClient creates a new APIKeyResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *APIKeyResourceDeleteDeleteParams {
	return &APIKeyResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
APIKeyResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the Api key resource delete delete operation.

	Typically these are written to a http.Request.
*/
type APIKeyResourceDeleteDeleteParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyResourceDeleteDeleteParams) WithDefaults() *APIKeyResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *APIKeyResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) WithContext(ctx context.Context) *APIKeyResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *APIKeyResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) WithAPIKeyID(aPIKeyID int64) *APIKeyResourceDeleteDeleteParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key resource delete delete params
func (o *APIKeyResourceDeleteDeleteParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKeyId
	if err := r.SetPathParam("apiKeyId", swag.FormatInt64(o.APIKeyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}