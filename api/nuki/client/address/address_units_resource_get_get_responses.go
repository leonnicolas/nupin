// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// AddressUnitsResourceGetGetReader is a Reader for the AddressUnitsResourceGetGet structure.
type AddressUnitsResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressUnitsResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressUnitsResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddressUnitsResourceGetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddressUnitsResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /address/{addressId}/unit] AddressUnitsResource_get_get", response, response.Code())
	}
}

// NewAddressUnitsResourceGetGetOK creates a AddressUnitsResourceGetGetOK with default headers values
func NewAddressUnitsResourceGetGetOK() *AddressUnitsResourceGetGetOK {
	return &AddressUnitsResourceGetGetOK{}
}

/*
AddressUnitsResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type AddressUnitsResourceGetGetOK struct {
	Payload []*models.AddressUnitResponse
}

// IsSuccess returns true when this address units resource get get o k response has a 2xx status code
func (o *AddressUnitsResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this address units resource get get o k response has a 3xx status code
func (o *AddressUnitsResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address units resource get get o k response has a 4xx status code
func (o *AddressUnitsResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this address units resource get get o k response has a 5xx status code
func (o *AddressUnitsResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this address units resource get get o k response a status code equal to that given
func (o *AddressUnitsResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the address units resource get get o k response
func (o *AddressUnitsResourceGetGetOK) Code() int {
	return 200
}

func (o *AddressUnitsResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AddressUnitsResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AddressUnitsResourceGetGetOK) GetPayload() []*models.AddressUnitResponse {
	return o.Payload
}

func (o *AddressUnitsResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressUnitsResourceGetGetBadRequest creates a AddressUnitsResourceGetGetBadRequest with default headers values
func NewAddressUnitsResourceGetGetBadRequest() *AddressUnitsResourceGetGetBadRequest {
	return &AddressUnitsResourceGetGetBadRequest{}
}

/*
AddressUnitsResourceGetGetBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type AddressUnitsResourceGetGetBadRequest struct {
}

// IsSuccess returns true when this address units resource get get bad request response has a 2xx status code
func (o *AddressUnitsResourceGetGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address units resource get get bad request response has a 3xx status code
func (o *AddressUnitsResourceGetGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address units resource get get bad request response has a 4xx status code
func (o *AddressUnitsResourceGetGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this address units resource get get bad request response has a 5xx status code
func (o *AddressUnitsResourceGetGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this address units resource get get bad request response a status code equal to that given
func (o *AddressUnitsResourceGetGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the address units resource get get bad request response
func (o *AddressUnitsResourceGetGetBadRequest) Code() int {
	return 400
}

func (o *AddressUnitsResourceGetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetBadRequest ", 400)
}

func (o *AddressUnitsResourceGetGetBadRequest) String() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetBadRequest ", 400)
}

func (o *AddressUnitsResourceGetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressUnitsResourceGetGetUnauthorized creates a AddressUnitsResourceGetGetUnauthorized with default headers values
func NewAddressUnitsResourceGetGetUnauthorized() *AddressUnitsResourceGetGetUnauthorized {
	return &AddressUnitsResourceGetGetUnauthorized{}
}

/*
AddressUnitsResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AddressUnitsResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this address units resource get get unauthorized response has a 2xx status code
func (o *AddressUnitsResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this address units resource get get unauthorized response has a 3xx status code
func (o *AddressUnitsResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address units resource get get unauthorized response has a 4xx status code
func (o *AddressUnitsResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this address units resource get get unauthorized response has a 5xx status code
func (o *AddressUnitsResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this address units resource get get unauthorized response a status code equal to that given
func (o *AddressUnitsResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the address units resource get get unauthorized response
func (o *AddressUnitsResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *AddressUnitsResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetUnauthorized ", 401)
}

func (o *AddressUnitsResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /address/{addressId}/unit][%d] addressUnitsResourceGetGetUnauthorized ", 401)
}

func (o *AddressUnitsResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
