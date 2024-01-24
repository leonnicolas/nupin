// Code generated by go-swagger; DO NOT EDIT.

package address_reservation

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

// NewAddressReservationsResourceGetGetParams creates a new AddressReservationsResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressReservationsResourceGetGetParams() *AddressReservationsResourceGetGetParams {
	return &AddressReservationsResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressReservationsResourceGetGetParamsWithTimeout creates a new AddressReservationsResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewAddressReservationsResourceGetGetParamsWithTimeout(timeout time.Duration) *AddressReservationsResourceGetGetParams {
	return &AddressReservationsResourceGetGetParams{
		timeout: timeout,
	}
}

// NewAddressReservationsResourceGetGetParamsWithContext creates a new AddressReservationsResourceGetGetParams object
// with the ability to set a context for a request.
func NewAddressReservationsResourceGetGetParamsWithContext(ctx context.Context) *AddressReservationsResourceGetGetParams {
	return &AddressReservationsResourceGetGetParams{
		Context: ctx,
	}
}

// NewAddressReservationsResourceGetGetParamsWithHTTPClient creates a new AddressReservationsResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressReservationsResourceGetGetParamsWithHTTPClient(client *http.Client) *AddressReservationsResourceGetGetParams {
	return &AddressReservationsResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
AddressReservationsResourceGetGetParams contains all the parameters to send to the API endpoint

	for the address reservations resource get get operation.

	Typically these are written to a http.Request.
*/
type AddressReservationsResourceGetGetParams struct {

	/* AddressID.

	   The address id
	*/
	AddressID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address reservations resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressReservationsResourceGetGetParams) WithDefaults() *AddressReservationsResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address reservations resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressReservationsResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) WithTimeout(timeout time.Duration) *AddressReservationsResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) WithContext(ctx context.Context) *AddressReservationsResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) WithHTTPClient(client *http.Client) *AddressReservationsResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) WithAddressID(addressID int64) *AddressReservationsResourceGetGetParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the address reservations resource get get params
func (o *AddressReservationsResourceGetGetParams) SetAddressID(addressID int64) {
	o.AddressID = addressID
}

// WriteToRequest writes these params to a swagger request
func (o *AddressReservationsResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
