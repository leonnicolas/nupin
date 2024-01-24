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

// NewAPIKeyAdvancedResourceDeleteDeleteParams creates a new APIKeyAdvancedResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyAdvancedResourceDeleteDeleteParams() *APIKeyAdvancedResourceDeleteDeleteParams {
	return &APIKeyAdvancedResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyAdvancedResourceDeleteDeleteParamsWithTimeout creates a new APIKeyAdvancedResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAPIKeyAdvancedResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *APIKeyAdvancedResourceDeleteDeleteParams {
	return &APIKeyAdvancedResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAPIKeyAdvancedResourceDeleteDeleteParamsWithContext creates a new APIKeyAdvancedResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAPIKeyAdvancedResourceDeleteDeleteParamsWithContext(ctx context.Context) *APIKeyAdvancedResourceDeleteDeleteParams {
	return &APIKeyAdvancedResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAPIKeyAdvancedResourceDeleteDeleteParamsWithHTTPClient creates a new APIKeyAdvancedResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyAdvancedResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *APIKeyAdvancedResourceDeleteDeleteParams {
	return &APIKeyAdvancedResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
APIKeyAdvancedResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the Api key advanced resource delete delete operation.

	Typically these are written to a http.Request.
*/
type APIKeyAdvancedResourceDeleteDeleteParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key advanced resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WithDefaults() *APIKeyAdvancedResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key advanced resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *APIKeyAdvancedResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WithContext(ctx context.Context) *APIKeyAdvancedResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *APIKeyAdvancedResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WithAPIKeyID(aPIKeyID int64) *APIKeyAdvancedResourceDeleteDeleteParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key advanced resource delete delete params
func (o *APIKeyAdvancedResourceDeleteDeleteParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyAdvancedResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
