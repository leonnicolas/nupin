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

// NewSmartlockAdminPinResourcePostPostParams creates a new SmartlockAdminPinResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockAdminPinResourcePostPostParams() *SmartlockAdminPinResourcePostPostParams {
	return &SmartlockAdminPinResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockAdminPinResourcePostPostParamsWithTimeout creates a new SmartlockAdminPinResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockAdminPinResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockAdminPinResourcePostPostParams {
	return &SmartlockAdminPinResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockAdminPinResourcePostPostParamsWithContext creates a new SmartlockAdminPinResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockAdminPinResourcePostPostParamsWithContext(ctx context.Context) *SmartlockAdminPinResourcePostPostParams {
	return &SmartlockAdminPinResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockAdminPinResourcePostPostParamsWithHTTPClient creates a new SmartlockAdminPinResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockAdminPinResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockAdminPinResourcePostPostParams {
	return &SmartlockAdminPinResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockAdminPinResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock admin pin resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockAdminPinResourcePostPostParams struct {

	/* Body.

	   Smartlock admin pin update representation
	*/
	Body *models.SmartlockAdminPinUpdate

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock admin pin resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAdminPinResourcePostPostParams) WithDefaults() *SmartlockAdminPinResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock admin pin resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAdminPinResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockAdminPinResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) WithContext(ctx context.Context) *SmartlockAdminPinResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockAdminPinResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) WithBody(body *models.SmartlockAdminPinUpdate) *SmartlockAdminPinResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) SetBody(body *models.SmartlockAdminPinUpdate) {
	o.Body = body
}

// WithSmartlockID adds the smartlockID to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) WithSmartlockID(smartlockID int64) *SmartlockAdminPinResourcePostPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock admin pin resource post post params
func (o *SmartlockAdminPinResourcePostPostParams) SetSmartlockID(smartlockID int64) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockAdminPinResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
