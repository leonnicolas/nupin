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

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// NewAPIKeyAdvancedResourcePutPutParams creates a new APIKeyAdvancedResourcePutPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIKeyAdvancedResourcePutPutParams() *APIKeyAdvancedResourcePutPutParams {
	return &APIKeyAdvancedResourcePutPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIKeyAdvancedResourcePutPutParamsWithTimeout creates a new APIKeyAdvancedResourcePutPutParams object
// with the ability to set a timeout on a request.
func NewAPIKeyAdvancedResourcePutPutParamsWithTimeout(timeout time.Duration) *APIKeyAdvancedResourcePutPutParams {
	return &APIKeyAdvancedResourcePutPutParams{
		timeout: timeout,
	}
}

// NewAPIKeyAdvancedResourcePutPutParamsWithContext creates a new APIKeyAdvancedResourcePutPutParams object
// with the ability to set a context for a request.
func NewAPIKeyAdvancedResourcePutPutParamsWithContext(ctx context.Context) *APIKeyAdvancedResourcePutPutParams {
	return &APIKeyAdvancedResourcePutPutParams{
		Context: ctx,
	}
}

// NewAPIKeyAdvancedResourcePutPutParamsWithHTTPClient creates a new APIKeyAdvancedResourcePutPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIKeyAdvancedResourcePutPutParamsWithHTTPClient(client *http.Client) *APIKeyAdvancedResourcePutPutParams {
	return &APIKeyAdvancedResourcePutPutParams{
		HTTPClient: client,
	}
}

/*
APIKeyAdvancedResourcePutPutParams contains all the parameters to send to the API endpoint

	for the Api key advanced resource put put operation.

	Typically these are written to a http.Request.
*/
type APIKeyAdvancedResourcePutPutParams struct {

	/* APIKeyID.

	   The api key id
	*/
	APIKeyID int64

	/* Body.

	   Apply for advaced api key representation
	*/
	Body *models.AdvancedAPIKeyCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Api key advanced resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourcePutPutParams) WithDefaults() *APIKeyAdvancedResourcePutPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Api key advanced resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIKeyAdvancedResourcePutPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) WithTimeout(timeout time.Duration) *APIKeyAdvancedResourcePutPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) WithContext(ctx context.Context) *APIKeyAdvancedResourcePutPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) WithHTTPClient(client *http.Client) *APIKeyAdvancedResourcePutPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) WithAPIKeyID(aPIKeyID int64) *APIKeyAdvancedResourcePutPutParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) SetAPIKeyID(aPIKeyID int64) {
	o.APIKeyID = aPIKeyID
}

// WithBody adds the body to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) WithBody(body *models.AdvancedAPIKeyCreate) *APIKeyAdvancedResourcePutPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Api key advanced resource put put params
func (o *APIKeyAdvancedResourcePutPutParams) SetBody(body *models.AdvancedAPIKeyCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *APIKeyAdvancedResourcePutPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apiKeyId
	if err := r.SetPathParam("apiKeyId", swag.FormatInt64(o.APIKeyID)); err != nil {
		return err
	}
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
