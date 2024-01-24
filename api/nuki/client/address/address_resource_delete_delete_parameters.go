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
)

// NewAddressResourceDeleteDeleteParams creates a new AddressResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressResourceDeleteDeleteParams() *AddressResourceDeleteDeleteParams {
	return &AddressResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressResourceDeleteDeleteParamsWithTimeout creates a new AddressResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAddressResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *AddressResourceDeleteDeleteParams {
	return &AddressResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAddressResourceDeleteDeleteParamsWithContext creates a new AddressResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAddressResourceDeleteDeleteParamsWithContext(ctx context.Context) *AddressResourceDeleteDeleteParams {
	return &AddressResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAddressResourceDeleteDeleteParamsWithHTTPClient creates a new AddressResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *AddressResourceDeleteDeleteParams {
	return &AddressResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
AddressResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the address resource delete delete operation.

	Typically these are written to a http.Request.
*/
type AddressResourceDeleteDeleteParams struct {

	/* AddressID.

	   The address id
	*/
	AddressID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressResourceDeleteDeleteParams) WithDefaults() *AddressResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *AddressResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) WithContext(ctx context.Context) *AddressResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *AddressResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) WithAddressID(addressID int64) *AddressResourceDeleteDeleteParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the address resource delete delete params
func (o *AddressResourceDeleteDeleteParams) SetAddressID(addressID int64) {
	o.AddressID = addressID
}

// WriteToRequest writes these params to a swagger request
func (o *AddressResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", swag.FormatInt64(o.AddressID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
