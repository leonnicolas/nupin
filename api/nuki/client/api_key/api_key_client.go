// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api key API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api key API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	APIKeyAdvancedReactivateResourcePostPost(params *APIKeyAdvancedReactivateResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedReactivateResourcePostPostNoContent, error)

	APIKeyAdvancedResourceDeleteDelete(params *APIKeyAdvancedResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourceDeleteDeleteNoContent, error)

	APIKeyAdvancedResourceGetGet(params *APIKeyAdvancedResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourceGetGetOK, error)

	APIKeyAdvancedResourcePostPost(params *APIKeyAdvancedResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourcePostPostNoContent, error)

	APIKeyAdvancedResourcePutPut(params *APIKeyAdvancedResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourcePutPutNoContent, error)

	APIKeyResourceDeleteDelete(params *APIKeyResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyResourceDeleteDeleteNoContent, error)

	APIKeyResourcePostPost(params *APIKeyResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyResourcePostPostNoContent, error)

	APIKeyTokenResourceDeleteDelete(params *APIKeyTokenResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokenResourceDeleteDeleteNoContent, error)

	APIKeyTokenResourcePostPost(params *APIKeyTokenResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokenResourcePostPostNoContent, error)

	APIKeyTokensResourceGetGet(params *APIKeyTokensResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokensResourceGetGetOK, error)

	APIKeyTokensResourcePutPut(params *APIKeyTokensResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokensResourcePutPutOK, error)

	APIKeysResourceGetGet(params *APIKeysResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeysResourceGetGetOK, error)

	APIKeysResourcePutPut(params *APIKeysResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeysResourcePutPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
APIKeyAdvancedReactivateResourcePostPost reactivates a deactivated advanced webhook integration
*/
func (a *Client) APIKeyAdvancedReactivateResourcePostPost(params *APIKeyAdvancedReactivateResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedReactivateResourcePostPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyAdvancedReactivateResourcePostPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyAdvancedReactivateResource_post_post",
		Method:             "POST",
		PathPattern:        "/api/key/{apiKeyId}/advanced/reactivate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyAdvancedReactivateResourcePostPostReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyAdvancedReactivateResourcePostPostNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyAdvancedReactivateResource_post_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyAdvancedResourceDeleteDelete deletes an advanced api key
*/
func (a *Client) APIKeyAdvancedResourceDeleteDelete(params *APIKeyAdvancedResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourceDeleteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyAdvancedResourceDeleteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyAdvancedResource_delete_delete",
		Method:             "DELETE",
		PathPattern:        "/api/key/{apiKeyId}/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyAdvancedResourceDeleteDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyAdvancedResourceDeleteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyAdvancedResource_delete_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyAdvancedResourceGetGet gets an advanced api key
*/
func (a *Client) APIKeyAdvancedResourceGetGet(params *APIKeyAdvancedResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyAdvancedResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyAdvancedResource_get_get",
		Method:             "GET",
		PathPattern:        "/api/key/{apiKeyId}/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyAdvancedResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyAdvancedResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyAdvancedResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyAdvancedResourcePostPost updates an advanced api key
*/
func (a *Client) APIKeyAdvancedResourcePostPost(params *APIKeyAdvancedResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourcePostPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyAdvancedResourcePostPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyAdvancedResource_post_post",
		Method:             "POST",
		PathPattern:        "/api/key/{apiKeyId}/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyAdvancedResourcePostPostReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyAdvancedResourcePostPostNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyAdvancedResource_post_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyAdvancedResourcePutPut creates an advanced api key
*/
func (a *Client) APIKeyAdvancedResourcePutPut(params *APIKeyAdvancedResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyAdvancedResourcePutPutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyAdvancedResourcePutPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyAdvancedResource_put_put",
		Method:             "PUT",
		PathPattern:        "/api/key/{apiKeyId}/advanced",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyAdvancedResourcePutPutReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyAdvancedResourcePutPutNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyAdvancedResource_put_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyResourceDeleteDelete deletes an api key
*/
func (a *Client) APIKeyResourceDeleteDelete(params *APIKeyResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyResourceDeleteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyResourceDeleteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyResource_delete_delete",
		Method:             "DELETE",
		PathPattern:        "/api/key/{apiKeyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyResourceDeleteDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyResourceDeleteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyResource_delete_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyResourcePostPost updates an api key
*/
func (a *Client) APIKeyResourcePostPost(params *APIKeyResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyResourcePostPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyResourcePostPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyResource_post_post",
		Method:             "POST",
		PathPattern:        "/api/key/{apiKeyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyResourcePostPostReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyResourcePostPostNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyResource_post_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyTokenResourceDeleteDelete deletes an api key token
*/
func (a *Client) APIKeyTokenResourceDeleteDelete(params *APIKeyTokenResourceDeleteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokenResourceDeleteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyTokenResourceDeleteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyTokenResource_delete_delete",
		Method:             "DELETE",
		PathPattern:        "/api/key/{apiKeyId}/token/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyTokenResourceDeleteDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyTokenResourceDeleteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyTokenResource_delete_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyTokenResourcePostPost updates an api key token
*/
func (a *Client) APIKeyTokenResourcePostPost(params *APIKeyTokenResourcePostPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokenResourcePostPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyTokenResourcePostPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyTokenResource_post_post",
		Method:             "POST",
		PathPattern:        "/api/key/{apiKeyId}/token/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyTokenResourcePostPostReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyTokenResourcePostPostNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyTokenResource_post_post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyTokensResourceGetGet gets a list of api key tokens
*/
func (a *Client) APIKeyTokensResourceGetGet(params *APIKeyTokensResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokensResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyTokensResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyTokensResource_get_get",
		Method:             "GET",
		PathPattern:        "/api/key/{apiKeyId}/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyTokensResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyTokensResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyTokensResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeyTokensResourcePutPut creates an api key token
*/
func (a *Client) APIKeyTokensResourcePutPut(params *APIKeyTokensResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeyTokensResourcePutPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyTokensResourcePutPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeyTokensResource_put_put",
		Method:             "PUT",
		PathPattern:        "/api/key/{apiKeyId}/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyTokensResourcePutPutReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeyTokensResourcePutPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeyTokensResource_put_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeysResourceGetGet gets a list of api keys
*/
func (a *Client) APIKeysResourceGetGet(params *APIKeysResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeysResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeysResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeysResource_get_get",
		Method:             "GET",
		PathPattern:        "/api/key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeysResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeysResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeysResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
APIKeysResourcePutPut creates an api key
*/
func (a *Client) APIKeysResourcePutPut(params *APIKeysResourcePutPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*APIKeysResourcePutPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeysResourcePutPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApiKeysResource_put_put",
		Method:             "PUT",
		PathPattern:        "/api/key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeysResourcePutPutReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*APIKeysResourcePutPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApiKeysResource_put_put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}