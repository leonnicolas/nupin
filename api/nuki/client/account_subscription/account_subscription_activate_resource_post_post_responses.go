// Code generated by go-swagger; DO NOT EDIT.

package account_subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountSubscriptionActivateResourcePostPostReader is a Reader for the AccountSubscriptionActivateResourcePostPost structure.
type AccountSubscriptionActivateResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountSubscriptionActivateResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAccountSubscriptionActivateResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountSubscriptionActivateResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /account/subscription/{id}/activate] AccountSubscriptionActivateResource_post_post", response, response.Code())
	}
}

// NewAccountSubscriptionActivateResourcePostPostNoContent creates a AccountSubscriptionActivateResourcePostPostNoContent with default headers values
func NewAccountSubscriptionActivateResourcePostPostNoContent() *AccountSubscriptionActivateResourcePostPostNoContent {
	return &AccountSubscriptionActivateResourcePostPostNoContent{}
}

/*
AccountSubscriptionActivateResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type AccountSubscriptionActivateResourcePostPostNoContent struct {
}

// IsSuccess returns true when this account subscription activate resource post post no content response has a 2xx status code
func (o *AccountSubscriptionActivateResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account subscription activate resource post post no content response has a 3xx status code
func (o *AccountSubscriptionActivateResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account subscription activate resource post post no content response has a 4xx status code
func (o *AccountSubscriptionActivateResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this account subscription activate resource post post no content response has a 5xx status code
func (o *AccountSubscriptionActivateResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this account subscription activate resource post post no content response a status code equal to that given
func (o *AccountSubscriptionActivateResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the account subscription activate resource post post no content response
func (o *AccountSubscriptionActivateResourcePostPostNoContent) Code() int {
	return 204
}

func (o *AccountSubscriptionActivateResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /account/subscription/{id}/activate][%d] accountSubscriptionActivateResourcePostPostNoContent ", 204)
}

func (o *AccountSubscriptionActivateResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /account/subscription/{id}/activate][%d] accountSubscriptionActivateResourcePostPostNoContent ", 204)
}

func (o *AccountSubscriptionActivateResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountSubscriptionActivateResourcePostPostUnauthorized creates a AccountSubscriptionActivateResourcePostPostUnauthorized with default headers values
func NewAccountSubscriptionActivateResourcePostPostUnauthorized() *AccountSubscriptionActivateResourcePostPostUnauthorized {
	return &AccountSubscriptionActivateResourcePostPostUnauthorized{}
}

/*
AccountSubscriptionActivateResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type AccountSubscriptionActivateResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this account subscription activate resource post post unauthorized response has a 2xx status code
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account subscription activate resource post post unauthorized response has a 3xx status code
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account subscription activate resource post post unauthorized response has a 4xx status code
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account subscription activate resource post post unauthorized response has a 5xx status code
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account subscription activate resource post post unauthorized response a status code equal to that given
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account subscription activate resource post post unauthorized response
func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account/subscription/{id}/activate][%d] accountSubscriptionActivateResourcePostPostUnauthorized ", 401)
}

func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /account/subscription/{id}/activate][%d] accountSubscriptionActivateResourcePostPostUnauthorized ", 401)
}

func (o *AccountSubscriptionActivateResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
