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
)

// NewSmartlocksResourceGetGetParams creates a new SmartlocksResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlocksResourceGetGetParams() *SmartlocksResourceGetGetParams {
	return &SmartlocksResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlocksResourceGetGetParamsWithTimeout creates a new SmartlocksResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewSmartlocksResourceGetGetParamsWithTimeout(timeout time.Duration) *SmartlocksResourceGetGetParams {
	return &SmartlocksResourceGetGetParams{
		timeout: timeout,
	}
}

// NewSmartlocksResourceGetGetParamsWithContext creates a new SmartlocksResourceGetGetParams object
// with the ability to set a context for a request.
func NewSmartlocksResourceGetGetParamsWithContext(ctx context.Context) *SmartlocksResourceGetGetParams {
	return &SmartlocksResourceGetGetParams{
		Context: ctx,
	}
}

// NewSmartlocksResourceGetGetParamsWithHTTPClient creates a new SmartlocksResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlocksResourceGetGetParamsWithHTTPClient(client *http.Client) *SmartlocksResourceGetGetParams {
	return &SmartlocksResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
SmartlocksResourceGetGetParams contains all the parameters to send to the API endpoint

	for the smartlocks resource get get operation.

	Typically these are written to a http.Request.
*/
type SmartlocksResourceGetGetParams struct {

	/* AuthID.

	   Filter for authId
	*/
	AuthID *int64

	/* Type.

	   Filter for type
	*/
	Type *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlocks resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlocksResourceGetGetParams) WithDefaults() *SmartlocksResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlocks resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlocksResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) WithTimeout(timeout time.Duration) *SmartlocksResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) WithContext(ctx context.Context) *SmartlocksResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) WithHTTPClient(client *http.Client) *SmartlocksResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthID adds the authID to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) WithAuthID(authID *int64) *SmartlocksResourceGetGetParams {
	o.SetAuthID(authID)
	return o
}

// SetAuthID adds the authId to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) SetAuthID(authID *int64) {
	o.AuthID = authID
}

// WithType adds the typeVar to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) WithType(typeVar *int64) *SmartlocksResourceGetGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the smartlocks resource get get params
func (o *SmartlocksResourceGetGetParams) SetType(typeVar *int64) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlocksResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthID != nil {

		// query param authId
		var qrAuthID int64

		if o.AuthID != nil {
			qrAuthID = *o.AuthID
		}
		qAuthID := swag.FormatInt64(qrAuthID)
		if qAuthID != "" {

			if err := r.SetQueryParam("authId", qAuthID); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType int64

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := swag.FormatInt64(qrType)
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
