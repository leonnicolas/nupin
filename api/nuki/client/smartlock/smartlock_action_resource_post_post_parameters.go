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

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// NewSmartlockActionResourcePostPostParams creates a new SmartlockActionResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockActionResourcePostPostParams() *SmartlockActionResourcePostPostParams {
	return &SmartlockActionResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockActionResourcePostPostParamsWithTimeout creates a new SmartlockActionResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockActionResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockActionResourcePostPostParams {
	return &SmartlockActionResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockActionResourcePostPostParamsWithContext creates a new SmartlockActionResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockActionResourcePostPostParamsWithContext(ctx context.Context) *SmartlockActionResourcePostPostParams {
	return &SmartlockActionResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockActionResourcePostPostParamsWithHTTPClient creates a new SmartlockActionResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockActionResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockActionResourcePostPostParams {
	return &SmartlockActionResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockActionResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock action resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockActionResourcePostPostParams struct {

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

// WithDefaults hydrates default values in the smartlock action resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockActionResourcePostPostParams) WithDefaults() *SmartlockActionResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock action resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockActionResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockActionResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) WithContext(ctx context.Context) *SmartlockActionResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockActionResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) WithBody(body *models.SmartlockAction) *SmartlockActionResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) SetBody(body *models.SmartlockAction) {
	o.Body = body
}

// WithSmartlockID adds the smartlockID to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) WithSmartlockID(smartlockID string) *SmartlockActionResourcePostPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock action resource post post params
func (o *SmartlockActionResourcePostPostParams) SetSmartlockID(smartlockID string) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockActionResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
