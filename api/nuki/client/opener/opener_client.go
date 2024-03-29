// Code generated by go-swagger; DO NOT EDIT.

package opener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new opener API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for opener API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	OpenerBrandResourceGetGet(params *OpenerBrandResourceGetGetParams, opts ...ClientOption) (*OpenerBrandResourceGetGetOK, error)

	OpenerBrandsResourceGetGet(params *OpenerBrandsResourceGetGetParams, opts ...ClientOption) (*OpenerBrandsResourceGetGetOK, error)

	OpenerIntercomResourceGetGet(params *OpenerIntercomResourceGetGetParams, opts ...ClientOption) (*OpenerIntercomResourceGetGetOK, error)

	OpenerIntercomsResourceGetGet(params *OpenerIntercomsResourceGetGetParams, opts ...ClientOption) (*OpenerIntercomsResourceGetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
OpenerBrandResourceGetGet gets an intercom brand
*/
func (a *Client) OpenerBrandResourceGetGet(params *OpenerBrandResourceGetGetParams, opts ...ClientOption) (*OpenerBrandResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenerBrandResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpenerBrandResource_get_get",
		Method:             "GET",
		PathPattern:        "/opener/brand/{brandId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenerBrandResourceGetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OpenerBrandResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpenerBrandResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OpenerBrandsResourceGetGet gets all intercom brands
*/
func (a *Client) OpenerBrandsResourceGetGet(params *OpenerBrandsResourceGetGetParams, opts ...ClientOption) (*OpenerBrandsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenerBrandsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpenerBrandsResource_get_get",
		Method:             "GET",
		PathPattern:        "/opener/brand",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenerBrandsResourceGetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OpenerBrandsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpenerBrandsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OpenerIntercomResourceGetGet gets an intercom model
*/
func (a *Client) OpenerIntercomResourceGetGet(params *OpenerIntercomResourceGetGetParams, opts ...ClientOption) (*OpenerIntercomResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenerIntercomResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpenerIntercomResource_get_get",
		Method:             "GET",
		PathPattern:        "/opener/intercom/{intercomId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenerIntercomResourceGetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OpenerIntercomResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpenerIntercomResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OpenerIntercomsResourceGetGet gets a list of intercom models
*/
func (a *Client) OpenerIntercomsResourceGetGet(params *OpenerIntercomsResourceGetGetParams, opts ...ClientOption) (*OpenerIntercomsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenerIntercomsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpenerIntercomsResource_get_get",
		Method:             "GET",
		PathPattern:        "/opener/intercom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenerIntercomsResourceGetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OpenerIntercomsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpenerIntercomsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
