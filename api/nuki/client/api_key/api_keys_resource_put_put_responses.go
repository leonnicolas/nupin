// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// APIKeysResourcePutPutReader is a Reader for the APIKeysResourcePutPut structure.
type APIKeysResourcePutPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIKeysResourcePutPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIKeysResourcePutPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAPIKeysResourcePutPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAPIKeysResourcePutPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/key] ApiKeysResource_put_put", response, response.Code())
	}
}

// NewAPIKeysResourcePutPutOK creates a APIKeysResourcePutPutOK with default headers values
func NewAPIKeysResourcePutPutOK() *APIKeysResourcePutPutOK {
	return &APIKeysResourcePutPutOK{}
}

/*
APIKeysResourcePutPutOK describes a response with status code 200, with default header values.

Ok
*/
type APIKeysResourcePutPutOK struct {
	Payload *models.APIKey
}

// IsSuccess returns true when this api keys resource put put o k response has a 2xx status code
func (o *APIKeysResourcePutPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api keys resource put put o k response has a 3xx status code
func (o *APIKeysResourcePutPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api keys resource put put o k response has a 4xx status code
func (o *APIKeysResourcePutPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api keys resource put put o k response has a 5xx status code
func (o *APIKeysResourcePutPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api keys resource put put o k response a status code equal to that given
func (o *APIKeysResourcePutPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api keys resource put put o k response
func (o *APIKeysResourcePutPutOK) Code() int {
	return 200
}

func (o *APIKeysResourcePutPutOK) Error() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutOK  %+v", 200, o.Payload)
}

func (o *APIKeysResourcePutPutOK) String() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutOK  %+v", 200, o.Payload)
}

func (o *APIKeysResourcePutPutOK) GetPayload() *models.APIKey {
	return o.Payload
}

func (o *APIKeysResourcePutPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeysResourcePutPutBadRequest creates a APIKeysResourcePutPutBadRequest with default headers values
func NewAPIKeysResourcePutPutBadRequest() *APIKeysResourcePutPutBadRequest {
	return &APIKeysResourcePutPutBadRequest{}
}

/*
APIKeysResourcePutPutBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type APIKeysResourcePutPutBadRequest struct {
}

// IsSuccess returns true when this api keys resource put put bad request response has a 2xx status code
func (o *APIKeysResourcePutPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api keys resource put put bad request response has a 3xx status code
func (o *APIKeysResourcePutPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api keys resource put put bad request response has a 4xx status code
func (o *APIKeysResourcePutPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this api keys resource put put bad request response has a 5xx status code
func (o *APIKeysResourcePutPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this api keys resource put put bad request response a status code equal to that given
func (o *APIKeysResourcePutPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the api keys resource put put bad request response
func (o *APIKeysResourcePutPutBadRequest) Code() int {
	return 400
}

func (o *APIKeysResourcePutPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutBadRequest ", 400)
}

func (o *APIKeysResourcePutPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutBadRequest ", 400)
}

func (o *APIKeysResourcePutPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIKeysResourcePutPutUnauthorized creates a APIKeysResourcePutPutUnauthorized with default headers values
func NewAPIKeysResourcePutPutUnauthorized() *APIKeysResourcePutPutUnauthorized {
	return &APIKeysResourcePutPutUnauthorized{}
}

/*
APIKeysResourcePutPutUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type APIKeysResourcePutPutUnauthorized struct {
}

// IsSuccess returns true when this api keys resource put put unauthorized response has a 2xx status code
func (o *APIKeysResourcePutPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api keys resource put put unauthorized response has a 3xx status code
func (o *APIKeysResourcePutPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api keys resource put put unauthorized response has a 4xx status code
func (o *APIKeysResourcePutPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this api keys resource put put unauthorized response has a 5xx status code
func (o *APIKeysResourcePutPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this api keys resource put put unauthorized response a status code equal to that given
func (o *APIKeysResourcePutPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the api keys resource put put unauthorized response
func (o *APIKeysResourcePutPutUnauthorized) Code() int {
	return 401
}

func (o *APIKeysResourcePutPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutUnauthorized ", 401)
}

func (o *APIKeysResourcePutPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/key][%d] apiKeysResourcePutPutUnauthorized ", 401)
}

func (o *APIKeysResourcePutPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
