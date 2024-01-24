// Code generated by go-swagger; DO NOT EDIT.

package smartlock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockActionResourcePostPostReader is a Reader for the SmartlockActionResourcePostPost structure.
type SmartlockActionResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockActionResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockActionResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockActionResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockActionResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewSmartlockActionResourcePostPostPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}/action] SmartlockActionResource_post_post", response, response.Code())
	}
}

// NewSmartlockActionResourcePostPostNoContent creates a SmartlockActionResourcePostPostNoContent with default headers values
func NewSmartlockActionResourcePostPostNoContent() *SmartlockActionResourcePostPostNoContent {
	return &SmartlockActionResourcePostPostNoContent{}
}

/*
SmartlockActionResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockActionResourcePostPostNoContent struct {
}

// IsSuccess returns true when this smartlock action resource post post no content response has a 2xx status code
func (o *SmartlockActionResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock action resource post post no content response has a 3xx status code
func (o *SmartlockActionResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock action resource post post no content response has a 4xx status code
func (o *SmartlockActionResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock action resource post post no content response has a 5xx status code
func (o *SmartlockActionResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock action resource post post no content response a status code equal to that given
func (o *SmartlockActionResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock action resource post post no content response
func (o *SmartlockActionResourcePostPostNoContent) Code() int {
	return 204
}

func (o *SmartlockActionResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostNoContent ", 204)
}

func (o *SmartlockActionResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostNoContent ", 204)
}

func (o *SmartlockActionResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockActionResourcePostPostBadRequest creates a SmartlockActionResourcePostPostBadRequest with default headers values
func NewSmartlockActionResourcePostPostBadRequest() *SmartlockActionResourcePostPostBadRequest {
	return &SmartlockActionResourcePostPostBadRequest{}
}

/*
SmartlockActionResourcePostPostBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type SmartlockActionResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this smartlock action resource post post bad request response has a 2xx status code
func (o *SmartlockActionResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock action resource post post bad request response has a 3xx status code
func (o *SmartlockActionResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock action resource post post bad request response has a 4xx status code
func (o *SmartlockActionResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock action resource post post bad request response has a 5xx status code
func (o *SmartlockActionResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock action resource post post bad request response a status code equal to that given
func (o *SmartlockActionResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock action resource post post bad request response
func (o *SmartlockActionResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockActionResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostBadRequest ", 400)
}

func (o *SmartlockActionResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostBadRequest ", 400)
}

func (o *SmartlockActionResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockActionResourcePostPostUnauthorized creates a SmartlockActionResourcePostPostUnauthorized with default headers values
func NewSmartlockActionResourcePostPostUnauthorized() *SmartlockActionResourcePostPostUnauthorized {
	return &SmartlockActionResourcePostPostUnauthorized{}
}

/*
SmartlockActionResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockActionResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock action resource post post unauthorized response has a 2xx status code
func (o *SmartlockActionResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock action resource post post unauthorized response has a 3xx status code
func (o *SmartlockActionResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock action resource post post unauthorized response has a 4xx status code
func (o *SmartlockActionResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock action resource post post unauthorized response has a 5xx status code
func (o *SmartlockActionResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock action resource post post unauthorized response a status code equal to that given
func (o *SmartlockActionResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock action resource post post unauthorized response
func (o *SmartlockActionResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockActionResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockActionResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockActionResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockActionResourcePostPostPaymentRequired creates a SmartlockActionResourcePostPostPaymentRequired with default headers values
func NewSmartlockActionResourcePostPostPaymentRequired() *SmartlockActionResourcePostPostPaymentRequired {
	return &SmartlockActionResourcePostPostPaymentRequired{}
}

/*
SmartlockActionResourcePostPostPaymentRequired describes a response with status code 402, with default header values.

Account not payed
*/
type SmartlockActionResourcePostPostPaymentRequired struct {
}

// IsSuccess returns true when this smartlock action resource post post payment required response has a 2xx status code
func (o *SmartlockActionResourcePostPostPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock action resource post post payment required response has a 3xx status code
func (o *SmartlockActionResourcePostPostPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock action resource post post payment required response has a 4xx status code
func (o *SmartlockActionResourcePostPostPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock action resource post post payment required response has a 5xx status code
func (o *SmartlockActionResourcePostPostPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock action resource post post payment required response a status code equal to that given
func (o *SmartlockActionResourcePostPostPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the smartlock action resource post post payment required response
func (o *SmartlockActionResourcePostPostPaymentRequired) Code() int {
	return 402
}

func (o *SmartlockActionResourcePostPostPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostPaymentRequired ", 402)
}

func (o *SmartlockActionResourcePostPostPaymentRequired) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action][%d] smartlockActionResourcePostPostPaymentRequired ", 402)
}

func (o *SmartlockActionResourcePostPostPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
