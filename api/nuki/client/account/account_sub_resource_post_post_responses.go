// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountSubResourcePostPostReader is a Reader for the AccountSubResourcePostPost structure.
type AccountSubResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountSubResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountSubResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountSubResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountSubResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAccountSubResourcePostPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /account/sub/{accountId}] AccountSubResource_post_post", response, response.Code())
	}
}

// NewAccountSubResourcePostPostNoContent creates a AccountSubResourcePostPostNoContent with default headers values
func NewAccountSubResourcePostPostNoContent() *AccountSubResourcePostPostNoContent {
	return &AccountSubResourcePostPostNoContent{}
}

/*
AccountSubResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type AccountSubResourcePostPostNoContent struct {
}

// IsSuccess returns true when this account sub resource post post no content response has a 2xx status code
func (o *AccountSubResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account sub resource post post no content response has a 3xx status code
func (o *AccountSubResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account sub resource post post no content response has a 4xx status code
func (o *AccountSubResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this account sub resource post post no content response has a 5xx status code
func (o *AccountSubResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this account sub resource post post no content response a status code equal to that given
func (o *AccountSubResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the account sub resource post post no content response
func (o *AccountSubResourcePostPostNoContent) Code() int {
	return 204
}

func (o *AccountSubResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostNoContent ", 204)
}

func (o *AccountSubResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostNoContent ", 204)
}

func (o *AccountSubResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSubResourcePostPostBadRequest creates a AccountSubResourcePostPostBadRequest with default headers values
func NewAccountSubResourcePostPostBadRequest() *AccountSubResourcePostPostBadRequest {
	return &AccountSubResourcePostPostBadRequest{}
}

/*
AccountSubResourcePostPostBadRequest describes a response with status code 400, with default header values.

Invalid parameter supplied
*/
type AccountSubResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this account sub resource post post bad request response has a 2xx status code
func (o *AccountSubResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account sub resource post post bad request response has a 3xx status code
func (o *AccountSubResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account sub resource post post bad request response has a 4xx status code
func (o *AccountSubResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account sub resource post post bad request response has a 5xx status code
func (o *AccountSubResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account sub resource post post bad request response a status code equal to that given
func (o *AccountSubResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account sub resource post post bad request response
func (o *AccountSubResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *AccountSubResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostBadRequest ", 400)
}

func (o *AccountSubResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostBadRequest ", 400)
}

func (o *AccountSubResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSubResourcePostPostUnauthorized creates a AccountSubResourcePostPostUnauthorized with default headers values
func NewAccountSubResourcePostPostUnauthorized() *AccountSubResourcePostPostUnauthorized {
	return &AccountSubResourcePostPostUnauthorized{}
}

/*
AccountSubResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountSubResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this account sub resource post post unauthorized response has a 2xx status code
func (o *AccountSubResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account sub resource post post unauthorized response has a 3xx status code
func (o *AccountSubResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account sub resource post post unauthorized response has a 4xx status code
func (o *AccountSubResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account sub resource post post unauthorized response has a 5xx status code
func (o *AccountSubResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account sub resource post post unauthorized response a status code equal to that given
func (o *AccountSubResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account sub resource post post unauthorized response
func (o *AccountSubResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *AccountSubResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostUnauthorized ", 401)
}

func (o *AccountSubResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostUnauthorized ", 401)
}

func (o *AccountSubResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSubResourcePostPostConflict creates a AccountSubResourcePostPostConflict with default headers values
func NewAccountSubResourcePostPostConflict() *AccountSubResourcePostPostConflict {
	return &AccountSubResourcePostPostConflict{}
}

/*
AccountSubResourcePostPostConflict describes a response with status code 409, with default header values.

E-Mail address already exists
*/
type AccountSubResourcePostPostConflict struct {
}

// IsSuccess returns true when this account sub resource post post conflict response has a 2xx status code
func (o *AccountSubResourcePostPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account sub resource post post conflict response has a 3xx status code
func (o *AccountSubResourcePostPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account sub resource post post conflict response has a 4xx status code
func (o *AccountSubResourcePostPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this account sub resource post post conflict response has a 5xx status code
func (o *AccountSubResourcePostPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this account sub resource post post conflict response a status code equal to that given
func (o *AccountSubResourcePostPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the account sub resource post post conflict response
func (o *AccountSubResourcePostPostConflict) Code() int {
	return 409
}

func (o *AccountSubResourcePostPostConflict) Error() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostConflict ", 409)
}

func (o *AccountSubResourcePostPostConflict) String() string {
	return fmt.Sprintf("[POST /account/sub/{accountId}][%d] accountSubResourcePostPostConflict ", 409)
}

func (o *AccountSubResourcePostPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
