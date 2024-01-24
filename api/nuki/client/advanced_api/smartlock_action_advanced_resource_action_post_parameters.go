// Code generated by go-swagger; DO NOT EDIT.

package advanced_api

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

// NewSmartlockActionAdvancedResourceActionPostParams creates a new SmartlockActionAdvancedResourceActionPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockActionAdvancedResourceActionPostParams() *SmartlockActionAdvancedResourceActionPostParams {
	return &SmartlockActionAdvancedResourceActionPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockActionAdvancedResourceActionPostParamsWithTimeout creates a new SmartlockActionAdvancedResourceActionPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockActionAdvancedResourceActionPostParamsWithTimeout(timeout time.Duration) *SmartlockActionAdvancedResourceActionPostParams {
	return &SmartlockActionAdvancedResourceActionPostParams{
		timeout: timeout,
	}
}

// NewSmartlockActionAdvancedResourceActionPostParamsWithContext creates a new SmartlockActionAdvancedResourceActionPostParams object
// with the ability to set a context for a request.
func NewSmartlockActionAdvancedResourceActionPostParamsWithContext(ctx context.Context) *SmartlockActionAdvancedResourceActionPostParams {
	return &SmartlockActionAdvancedResourceActionPostParams{
		Context: ctx,
	}
}

// NewSmartlockActionAdvancedResourceActionPostParamsWithHTTPClient creates a new SmartlockActionAdvancedResourceActionPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockActionAdvancedResourceActionPostParamsWithHTTPClient(client *http.Client) *SmartlockActionAdvancedResourceActionPostParams {
	return &SmartlockActionAdvancedResourceActionPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockActionAdvancedResourceActionPostParams contains all the parameters to send to the API endpoint

	for the smartlock action advanced resource action post operation.

	Typically these are written to a http.Request.
*/
type SmartlockActionAdvancedResourceActionPostParams struct {

	/* Body.

	   Smartlock action representation
	*/
	Body *models.SmartlockAction

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock action advanced resource action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockActionAdvancedResourceActionPostParams) WithDefaults() *SmartlockActionAdvancedResourceActionPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock action advanced resource action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockActionAdvancedResourceActionPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) WithTimeout(timeout time.Duration) *SmartlockActionAdvancedResourceActionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) WithContext(ctx context.Context) *SmartlockActionAdvancedResourceActionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) WithHTTPClient(client *http.Client) *SmartlockActionAdvancedResourceActionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) WithBody(body *models.SmartlockAction) *SmartlockActionAdvancedResourceActionPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) SetBody(body *models.SmartlockAction) {
	o.Body = body
}

// WithSmartlockID adds the smartlockID to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) WithSmartlockID(smartlockID string) *SmartlockActionAdvancedResourceActionPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock action advanced resource action post params
func (o *SmartlockActionAdvancedResourceActionPostParams) SetSmartlockID(smartlockID string) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockActionAdvancedResourceActionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("smartlockId", o.SmartlockID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
