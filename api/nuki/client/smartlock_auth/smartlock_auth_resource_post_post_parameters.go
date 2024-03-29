// Code generated by go-swagger; DO NOT EDIT.

package smartlock_auth

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

// NewSmartlockAuthResourcePostPostParams creates a new SmartlockAuthResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockAuthResourcePostPostParams() *SmartlockAuthResourcePostPostParams {
	return &SmartlockAuthResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockAuthResourcePostPostParamsWithTimeout creates a new SmartlockAuthResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockAuthResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockAuthResourcePostPostParams {
	return &SmartlockAuthResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockAuthResourcePostPostParamsWithContext creates a new SmartlockAuthResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockAuthResourcePostPostParamsWithContext(ctx context.Context) *SmartlockAuthResourcePostPostParams {
	return &SmartlockAuthResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockAuthResourcePostPostParamsWithHTTPClient creates a new SmartlockAuthResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockAuthResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockAuthResourcePostPostParams {
	return &SmartlockAuthResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockAuthResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock auth resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockAuthResourcePostPostParams struct {

	/* Body.

	   Smartlock authorization update representation
	*/
	Body *models.SmartlockAuthUpdate

	/* ID.

	   The smartlock authorization unique id
	*/
	ID string

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock auth resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAuthResourcePostPostParams) WithDefaults() *SmartlockAuthResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock auth resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAuthResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockAuthResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithContext(ctx context.Context) *SmartlockAuthResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockAuthResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithBody(body *models.SmartlockAuthUpdate) *SmartlockAuthResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetBody(body *models.SmartlockAuthUpdate) {
	o.Body = body
}

// WithID adds the id to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithID(id string) *SmartlockAuthResourcePostPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetID(id string) {
	o.ID = id
}

// WithSmartlockID adds the smartlockID to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) WithSmartlockID(smartlockID int64) *SmartlockAuthResourcePostPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock auth resource post post params
func (o *SmartlockAuthResourcePostPostParams) SetSmartlockID(smartlockID int64) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockAuthResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
