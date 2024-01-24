// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// AccountSettingResourceGetGetReader is a Reader for the AccountSettingResourceGetGet structure.
type AccountSettingResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountSettingResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountSettingResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountSettingResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountSettingResourceGetGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountSettingResourceGetGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /account/setting] AccountSettingResource_get_get", response, response.Code())
	}
}

// NewAccountSettingResourceGetGetOK creates a AccountSettingResourceGetGetOK with default headers values
func NewAccountSettingResourceGetGetOK() *AccountSettingResourceGetGetOK {
	return &AccountSettingResourceGetGetOK{}
}

/*
AccountSettingResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type AccountSettingResourceGetGetOK struct {
	Payload *models.AccountSetting
}

// IsSuccess returns true when this account setting resource get get o k response has a 2xx status code
func (o *AccountSettingResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account setting resource get get o k response has a 3xx status code
func (o *AccountSettingResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account setting resource get get o k response has a 4xx status code
func (o *AccountSettingResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account setting resource get get o k response has a 5xx status code
func (o *AccountSettingResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account setting resource get get o k response a status code equal to that given
func (o *AccountSettingResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account setting resource get get o k response
func (o *AccountSettingResourceGetGetOK) Code() int {
	return 200
}

func (o *AccountSettingResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountSettingResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *AccountSettingResourceGetGetOK) GetPayload() *models.AccountSetting {
	return o.Payload
}

func (o *AccountSettingResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountSettingResourceGetGetUnauthorized creates a AccountSettingResourceGetGetUnauthorized with default headers values
func NewAccountSettingResourceGetGetUnauthorized() *AccountSettingResourceGetGetUnauthorized {
	return &AccountSettingResourceGetGetUnauthorized{}
}

/*
AccountSettingResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountSettingResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this account setting resource get get unauthorized response has a 2xx status code
func (o *AccountSettingResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account setting resource get get unauthorized response has a 3xx status code
func (o *AccountSettingResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account setting resource get get unauthorized response has a 4xx status code
func (o *AccountSettingResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account setting resource get get unauthorized response has a 5xx status code
func (o *AccountSettingResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account setting resource get get unauthorized response a status code equal to that given
func (o *AccountSettingResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account setting resource get get unauthorized response
func (o *AccountSettingResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *AccountSettingResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetUnauthorized ", 401)
}

func (o *AccountSettingResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetUnauthorized ", 401)
}

func (o *AccountSettingResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSettingResourceGetGetForbidden creates a AccountSettingResourceGetGetForbidden with default headers values
func NewAccountSettingResourceGetGetForbidden() *AccountSettingResourceGetGetForbidden {
	return &AccountSettingResourceGetGetForbidden{}
}

/*
AccountSettingResourceGetGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccountSettingResourceGetGetForbidden struct {
}

// IsSuccess returns true when this account setting resource get get forbidden response has a 2xx status code
func (o *AccountSettingResourceGetGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account setting resource get get forbidden response has a 3xx status code
func (o *AccountSettingResourceGetGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account setting resource get get forbidden response has a 4xx status code
func (o *AccountSettingResourceGetGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this account setting resource get get forbidden response has a 5xx status code
func (o *AccountSettingResourceGetGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this account setting resource get get forbidden response a status code equal to that given
func (o *AccountSettingResourceGetGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the account setting resource get get forbidden response
func (o *AccountSettingResourceGetGetForbidden) Code() int {
	return 403
}

func (o *AccountSettingResourceGetGetForbidden) Error() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetForbidden ", 403)
}

func (o *AccountSettingResourceGetGetForbidden) String() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetForbidden ", 403)
}

func (o *AccountSettingResourceGetGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSettingResourceGetGetNotFound creates a AccountSettingResourceGetGetNotFound with default headers values
func NewAccountSettingResourceGetGetNotFound() *AccountSettingResourceGetGetNotFound {
	return &AccountSettingResourceGetGetNotFound{}
}

/*
AccountSettingResourceGetGetNotFound describes a response with status code 404, with default header values.

Not found
*/
type AccountSettingResourceGetGetNotFound struct {
}

// IsSuccess returns true when this account setting resource get get not found response has a 2xx status code
func (o *AccountSettingResourceGetGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account setting resource get get not found response has a 3xx status code
func (o *AccountSettingResourceGetGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account setting resource get get not found response has a 4xx status code
func (o *AccountSettingResourceGetGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account setting resource get get not found response has a 5xx status code
func (o *AccountSettingResourceGetGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account setting resource get get not found response a status code equal to that given
func (o *AccountSettingResourceGetGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account setting resource get get not found response
func (o *AccountSettingResourceGetGetNotFound) Code() int {
	return 404
}

func (o *AccountSettingResourceGetGetNotFound) Error() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetNotFound ", 404)
}

func (o *AccountSettingResourceGetGetNotFound) String() string {
	return fmt.Sprintf("[GET /account/setting][%d] accountSettingResourceGetGetNotFound ", 404)
}

func (o *AccountSettingResourceGetGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
