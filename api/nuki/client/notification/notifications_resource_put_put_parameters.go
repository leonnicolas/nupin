// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewNotificationsResourcePutPutParams creates a new NotificationsResourcePutPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationsResourcePutPutParams() *NotificationsResourcePutPutParams {
	return &NotificationsResourcePutPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsResourcePutPutParamsWithTimeout creates a new NotificationsResourcePutPutParams object
// with the ability to set a timeout on a request.
func NewNotificationsResourcePutPutParamsWithTimeout(timeout time.Duration) *NotificationsResourcePutPutParams {
	return &NotificationsResourcePutPutParams{
		timeout: timeout,
	}
}

// NewNotificationsResourcePutPutParamsWithContext creates a new NotificationsResourcePutPutParams object
// with the ability to set a context for a request.
func NewNotificationsResourcePutPutParamsWithContext(ctx context.Context) *NotificationsResourcePutPutParams {
	return &NotificationsResourcePutPutParams{
		Context: ctx,
	}
}

// NewNotificationsResourcePutPutParamsWithHTTPClient creates a new NotificationsResourcePutPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationsResourcePutPutParamsWithHTTPClient(client *http.Client) *NotificationsResourcePutPutParams {
	return &NotificationsResourcePutPutParams{
		HTTPClient: client,
	}
}

/*
NotificationsResourcePutPutParams contains all the parameters to send to the API endpoint

	for the notifications resource put put operation.

	Typically these are written to a http.Request.
*/
type NotificationsResourcePutPutParams struct {

	/* Body.

	   Notification representation
	*/
	Body *models.Notification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notifications resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsResourcePutPutParams) WithDefaults() *NotificationsResourcePutPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifications resource put put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsResourcePutPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) WithTimeout(timeout time.Duration) *NotificationsResourcePutPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) WithContext(ctx context.Context) *NotificationsResourcePutPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) WithHTTPClient(client *http.Client) *NotificationsResourcePutPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) WithBody(body *models.Notification) *NotificationsResourcePutPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notifications resource put put params
func (o *NotificationsResourcePutPutParams) SetBody(body *models.Notification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsResourcePutPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
