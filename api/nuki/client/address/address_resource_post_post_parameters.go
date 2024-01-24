// Code generated by go-swagger; DO NOT EDIT.

package address

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

// NewAddressResourcePostPostParams creates a new AddressResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressResourcePostPostParams() *AddressResourcePostPostParams {
	return &AddressResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressResourcePostPostParamsWithTimeout creates a new AddressResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAddressResourcePostPostParamsWithTimeout(timeout time.Duration) *AddressResourcePostPostParams {
	return &AddressResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAddressResourcePostPostParamsWithContext creates a new AddressResourcePostPostParams object
// with the ability to set a context for a request.
func NewAddressResourcePostPostParamsWithContext(ctx context.Context) *AddressResourcePostPostParams {
	return &AddressResourcePostPostParams{
		Context: ctx,
	}
}

// NewAddressResourcePostPostParamsWithHTTPClient creates a new AddressResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressResourcePostPostParamsWithHTTPClient(client *http.Client) *AddressResourcePostPostParams {
	return &AddressResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
AddressResourcePostPostParams contains all the parameters to send to the API endpoint

	for the address resource post post operation.

	Typically these are written to a http.Request.
*/
type AddressResourcePostPostParams struct {

	/* AddressID.

	   The address id
	*/
	AddressID int64

	/* Body.

	   Address update representation
	*/
	Body *models.AddressUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressResourcePostPostParams) WithDefaults() *AddressResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address resource post post params
func (o *AddressResourcePostPostParams) WithTimeout(timeout time.Duration) *AddressResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address resource post post params
func (o *AddressResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address resource post post params
func (o *AddressResourcePostPostParams) WithContext(ctx context.Context) *AddressResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address resource post post params
func (o *AddressResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address resource post post params
func (o *AddressResourcePostPostParams) WithHTTPClient(client *http.Client) *AddressResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address resource post post params
func (o *AddressResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the address resource post post params
func (o *AddressResourcePostPostParams) WithAddressID(addressID int64) *AddressResourcePostPostParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the address resource post post params
func (o *AddressResourcePostPostParams) SetAddressID(addressID int64) {
	o.AddressID = addressID
}

// WithBody adds the body to the address resource post post params
func (o *AddressResourcePostPostParams) WithBody(body *models.AddressUpdate) *AddressResourcePostPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the address resource post post params
func (o *AddressResourcePostPostParams) SetBody(body *models.AddressUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddressResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", swag.FormatInt64(o.AddressID)); err != nil {
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