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

// NewAddressReservationRevokeResourcePostPostParams creates a new AddressReservationRevokeResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressReservationRevokeResourcePostPostParams() *AddressReservationRevokeResourcePostPostParams {
	return &AddressReservationRevokeResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressReservationRevokeResourcePostPostParamsWithTimeout creates a new AddressReservationRevokeResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewAddressReservationRevokeResourcePostPostParamsWithTimeout(timeout time.Duration) *AddressReservationRevokeResourcePostPostParams {
	return &AddressReservationRevokeResourcePostPostParams{
		timeout: timeout,
	}
}

// NewAddressReservationRevokeResourcePostPostParamsWithContext creates a new AddressReservationRevokeResourcePostPostParams object
// with the ability to set a context for a request.
func NewAddressReservationRevokeResourcePostPostParamsWithContext(ctx context.Context) *AddressReservationRevokeResourcePostPostParams {
	return &AddressReservationRevokeResourcePostPostParams{
		Context: ctx,
	}
}

// NewAddressReservationRevokeResourcePostPostParamsWithHTTPClient creates a new AddressReservationRevokeResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressReservationRevokeResourcePostPostParamsWithHTTPClient(client *http.Client) *AddressReservationRevokeResourcePostPostParams {
	return &AddressReservationRevokeResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
AddressReservationRevokeResourcePostPostParams contains all the parameters to send to the API endpoint

	for the address reservation revoke resource post post operation.

	Typically these are written to a http.Request.
*/
type AddressReservationRevokeResourcePostPostParams struct {

	/* AddressID.

	   The address id
	*/
	AddressID int64

	/* ID.

	   The address reservation id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address reservation revoke resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressReservationRevokeResourcePostPostParams) WithDefaults() *AddressReservationRevokeResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address reservation revoke resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressReservationRevokeResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) WithTimeout(timeout time.Duration) *AddressReservationRevokeResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) WithContext(ctx context.Context) *AddressReservationRevokeResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) WithHTTPClient(client *http.Client) *AddressReservationRevokeResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) WithAddressID(addressID int64) *AddressReservationRevokeResourcePostPostParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) SetAddressID(addressID int64) {
	o.AddressID = addressID
}

// WithID adds the id to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) WithID(id string) *AddressReservationRevokeResourcePostPostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address reservation revoke resource post post params
func (o *AddressReservationRevokeResourcePostPostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressReservationRevokeResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", swag.FormatInt64(o.AddressID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
