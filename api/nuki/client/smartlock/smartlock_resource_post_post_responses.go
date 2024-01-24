// Code generated by go-swagger; DO NOT EDIT.

package smartlock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockResourcePostPostReader is a Reader for the SmartlockResourcePostPost structure.
type SmartlockResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSmartlockResourcePostPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}] SmartlockResource_post_post", response, response.Code())
	}
}

// NewSmartlockResourcePostPostNoContent creates a SmartlockResourcePostPostNoContent with default headers values
func NewSmartlockResourcePostPostNoContent() *SmartlockResourcePostPostNoContent {
	return &SmartlockResourcePostPostNoContent{}
}

/*
SmartlockResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockResourcePostPostNoContent struct {
}

// IsSuccess returns true when this smartlock resource post post no content response has a 2xx status code
func (o *SmartlockResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock resource post post no content response has a 3xx status code
func (o *SmartlockResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock resource post post no content response has a 4xx status code
func (o *SmartlockResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock resource post post no content response has a 5xx status code
func (o *SmartlockResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock resource post post no content response a status code equal to that given
func (o *SmartlockResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock resource post post no content response
func (o *SmartlockResourcePostPostNoContent) Code() int {
	return 204
}

func (o *SmartlockResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostNoContent ", 204)
}

func (o *SmartlockResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostNoContent ", 204)
}

func (o *SmartlockResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockResourcePostPostBadRequest creates a SmartlockResourcePostPostBadRequest with default headers values
func NewSmartlockResourcePostPostBadRequest() *SmartlockResourcePostPostBadRequest {
	return &SmartlockResourcePostPostBadRequest{}
}

/*
SmartlockResourcePostPostBadRequest describes a response with status code 400, with default header values.

Invalid parameter given
*/
type SmartlockResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this smartlock resource post post bad request response has a 2xx status code
func (o *SmartlockResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock resource post post bad request response has a 3xx status code
func (o *SmartlockResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock resource post post bad request response has a 4xx status code
func (o *SmartlockResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock resource post post bad request response has a 5xx status code
func (o *SmartlockResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock resource post post bad request response a status code equal to that given
func (o *SmartlockResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock resource post post bad request response
func (o *SmartlockResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostBadRequest ", 400)
}

func (o *SmartlockResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostBadRequest ", 400)
}

func (o *SmartlockResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockResourcePostPostUnauthorized creates a SmartlockResourcePostPostUnauthorized with default headers values
func NewSmartlockResourcePostPostUnauthorized() *SmartlockResourcePostPostUnauthorized {
	return &SmartlockResourcePostPostUnauthorized{}
}

/*
SmartlockResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock resource post post unauthorized response has a 2xx status code
func (o *SmartlockResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock resource post post unauthorized response has a 3xx status code
func (o *SmartlockResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock resource post post unauthorized response has a 4xx status code
func (o *SmartlockResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock resource post post unauthorized response has a 5xx status code
func (o *SmartlockResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock resource post post unauthorized response a status code equal to that given
func (o *SmartlockResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock resource post post unauthorized response
func (o *SmartlockResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockResourcePostPostForbidden creates a SmartlockResourcePostPostForbidden with default headers values
func NewSmartlockResourcePostPostForbidden() *SmartlockResourcePostPostForbidden {
	return &SmartlockResourcePostPostForbidden{}
}

/*
SmartlockResourcePostPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SmartlockResourcePostPostForbidden struct {
}

// IsSuccess returns true when this smartlock resource post post forbidden response has a 2xx status code
func (o *SmartlockResourcePostPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock resource post post forbidden response has a 3xx status code
func (o *SmartlockResourcePostPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock resource post post forbidden response has a 4xx status code
func (o *SmartlockResourcePostPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock resource post post forbidden response has a 5xx status code
func (o *SmartlockResourcePostPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock resource post post forbidden response a status code equal to that given
func (o *SmartlockResourcePostPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the smartlock resource post post forbidden response
func (o *SmartlockResourcePostPostForbidden) Code() int {
	return 403
}

func (o *SmartlockResourcePostPostForbidden) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostForbidden ", 403)
}

func (o *SmartlockResourcePostPostForbidden) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}][%d] smartlockResourcePostPostForbidden ", 403)
}

func (o *SmartlockResourcePostPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}