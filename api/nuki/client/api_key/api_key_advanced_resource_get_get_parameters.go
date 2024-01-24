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

// NewAPIKeyAdvancedResourceGetGetParams creates a new APIKeyAdvancedResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyAdvancedResourceGetGetParams() *APIKeyAdvancedResourceGetGetParams {
	return &APIKeyAdvancedResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyAdvancedResourceGetGetParamsWithTimeout creates a new APIKeyAdvancedResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewAPIKeyAdvancedResourceGetGetParamsWithTimeout(timeout time.Duration) *APIKeyAdvancedResourceGetGetParams {
	return &APIKeyAdvancedResourceGetGetParams{
		timeout: timeout,
	}
}

// NewAPIKeyAdvancedResourceGetGetParamsWithContext creates a new APIKeyAdvancedResourceGetGetParams object
// with the ability to set a context for a request.
func NewAPIKeyAdvancedResourceGetGetParamsWithContext(ctx context.Context) *APIKeyAdvancedResourceGetGetParams {
	return &APIKeyAdvancedResourceGetGetParams{
		Context: ctx,
	}
}

// NewAPIKeyAdvancedResourceGetGetParamsWithHTTPClient creates a new APIKeyAdvancedResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyAdvancedResourceGetGetParamsWithHTTPClient(client *http.Client) *APIKeyAdvancedResourceGetGetParams {
	return &APIKeyAdvancedResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
APIKeyAdvancedResourceGetGetParams contains all the parameters to send to the API endpoint

	for the Api key advanced resource get get operation.

	Typically these are written to a http.Request.
*/
type APIKeyAdvancedResourceGetGetParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key advanced resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourceGetGetParams) WithDefaults() *APIKeyAdvancedResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key advanced resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) WithTimeout(timeout time.Duration) *APIKeyAdvancedResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) WithContext(ctx context.Context) *APIKeyAdvancedResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) WithHTTPClient(client *http.Client) *APIKeyAdvancedResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) WithAPIKeyID(aPIKeyID int64) *APIKeyAdvancedResourceGetGetParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key advanced resource get get params
func (o *APIKeyAdvancedResourceGetGetParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyAdvancedResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
