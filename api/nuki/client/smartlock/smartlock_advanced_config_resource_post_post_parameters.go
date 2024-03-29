// Code generated by go-swagger; DO NOT EDIT.

package smartlock

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

// NewSmartlockAdvancedConfigResourcePostPostParams creates a new SmartlockAdvancedConfigResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockAdvancedConfigResourcePostPostParams() *SmartlockAdvancedConfigResourcePostPostParams {
	return &SmartlockAdvancedConfigResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockAdvancedConfigResourcePostPostParamsWithTimeout creates a new SmartlockAdvancedConfigResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockAdvancedConfigResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockAdvancedConfigResourcePostPostParams {
	return &SmartlockAdvancedConfigResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockAdvancedConfigResourcePostPostParamsWithContext creates a new SmartlockAdvancedConfigResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockAdvancedConfigResourcePostPostParamsWithContext(ctx context.Context) *SmartlockAdvancedConfigResourcePostPostParams {
	return &SmartlockAdvancedConfigResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockAdvancedConfigResourcePostPostParamsWithHTTPClient creates a new SmartlockAdvancedConfigResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockAdvancedConfigResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockAdvancedConfigResourcePostPostParams {
	return &SmartlockAdvancedConfigResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockAdvancedConfigResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock advanced config resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockAdvancedConfigResourcePostPostParams struct {

	/* Body.

	   Smartlock config update representation
	*/
	Body *models.SmartlockAdvancedConfig

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock advanced config resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithDefaults() *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock advanced config resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithContext(ctx context.Context) *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithBody(body *models.SmartlockAdvancedConfig) *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetBody(body *models.SmartlockAdvancedConfig) {
	o.Body = body
}

// WithSmartlockID adds the smartlockID to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) WithSmartlockID(smartlockID int64) *SmartlockAdvancedConfigResourcePostPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock advanced config resource post post params
func (o *SmartlockAdvancedConfigResourcePostPostParams) SetSmartlockID(smartlockID int64) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockAdvancedConfigResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param smartlockId
	if err := r.SetPathParam("smartlockId", swag.FormatInt64(o.SmartlockID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
