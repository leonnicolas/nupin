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

// APIKeysResourceGetGetReader is a Reader for the APIKeysResourceGetGet structure.
type APIKeysResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIKeysResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIKeysResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAPIKeysResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/key] ApiKeysResource_get_get", response, response.Code())
	}
}

// NewAPIKeysResourceGetGetOK creates a APIKeysResourceGetGetOK with default headers values
func NewAPIKeysResourceGetGetOK() *APIKeysResourceGetGetOK {
	return &APIKeysResourceGetGetOK{}
}

/*
APIKeysResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type APIKeysResourceGetGetOK struct {
	Payload []*models.APIKey
}

// IsSuccess returns true when this api keys resource get get o k response has a 2xx status code
func (o *APIKeysResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api keys resource get get o k response has a 3xx status code
func (o *APIKeysResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api keys resource get get o k response has a 4xx status code
func (o *APIKeysResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this api keys resource get get o k response has a 5xx status code
func (o *APIKeysResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this api keys resource get get o k response a status code equal to that given
func (o *APIKeysResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the api keys resource get get o k response
func (o *APIKeysResourceGetGetOK) Code() int {
	return 200
}

func (o *APIKeysResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /api/key][%d] apiKeysResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *APIKeysResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /api/key][%d] apiKeysResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *APIKeysResourceGetGetOK) GetPayload() []*models.APIKey {
	return o.Payload
}

func (o *APIKeysResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeysResourceGetGetUnauthorized creates a APIKeysResourceGetGetUnauthorized with default headers values
func NewAPIKeysResourceGetGetUnauthorized() *APIKeysResourceGetGetUnauthorized {
	return &APIKeysResourceGetGetUnauthorized{}
}

/*
APIKeysResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type APIKeysResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this api keys resource get get unauthorized response has a 2xx status code
func (o *APIKeysResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api keys resource get get unauthorized response has a 3xx status code
func (o *APIKeysResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api keys resource get get unauthorized response has a 4xx status code
func (o *APIKeysResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this api keys resource get get unauthorized response has a 5xx status code
func (o *APIKeysResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this api keys resource get get unauthorized response a status code equal to that given
func (o *APIKeysResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the api keys resource get get unauthorized response
func (o *APIKeysResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *APIKeysResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/key][%d] apiKeysResourceGetGetUnauthorized ", 401)
}

func (o *APIKeysResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/key][%d] apiKeysResourceGetGetUnauthorized ", 401)
}

func (o *APIKeysResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
