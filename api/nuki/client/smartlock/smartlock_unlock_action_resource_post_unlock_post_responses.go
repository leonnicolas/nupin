// Code generated by go-swagger; DO NOT EDIT.

package smartlock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockUnlockActionResourcePostUnlockPostReader is a Reader for the SmartlockUnlockActionResourcePostUnlockPost structure.
type SmartlockUnlockActionResourcePostUnlockPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockUnlockActionResourcePostUnlockPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockUnlockActionResourcePostUnlockPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockUnlockActionResourcePostUnlockPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockUnlockActionResourcePostUnlockPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewSmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}/action/unlock] SmartlockUnlockActionResource_postUnlock_post", response, response.Code())
	}
}

// NewSmartlockUnlockActionResourcePostUnlockPostNoContent creates a SmartlockUnlockActionResourcePostUnlockPostNoContent with default headers values
func NewSmartlockUnlockActionResourcePostUnlockPostNoContent() *SmartlockUnlockActionResourcePostUnlockPostNoContent {
	return &SmartlockUnlockActionResourcePostUnlockPostNoContent{}
}

/*
SmartlockUnlockActionResourcePostUnlockPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockUnlockActionResourcePostUnlockPostNoContent struct {
}

// IsSuccess returns true when this smartlock unlock action resource post unlock post no content response has a 2xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock unlock action resource post unlock post no content response has a 3xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock unlock action resource post unlock post no content response has a 4xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock unlock action resource post unlock post no content response has a 5xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock unlock action resource post unlock post no content response a status code equal to that given
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock unlock action resource post unlock post no content response
func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) Code() int {
	return 204
}

func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostNoContent ", 204)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostNoContent ", 204)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockUnlockActionResourcePostUnlockPostBadRequest creates a SmartlockUnlockActionResourcePostUnlockPostBadRequest with default headers values
func NewSmartlockUnlockActionResourcePostUnlockPostBadRequest() *SmartlockUnlockActionResourcePostUnlockPostBadRequest {
	return &SmartlockUnlockActionResourcePostUnlockPostBadRequest{}
}

/*
SmartlockUnlockActionResourcePostUnlockPostBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type SmartlockUnlockActionResourcePostUnlockPostBadRequest struct {
}

// IsSuccess returns true when this smartlock unlock action resource post unlock post bad request response has a 2xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock unlock action resource post unlock post bad request response has a 3xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock unlock action resource post unlock post bad request response has a 4xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock unlock action resource post unlock post bad request response has a 5xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock unlock action resource post unlock post bad request response a status code equal to that given
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock unlock action resource post unlock post bad request response
func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostBadRequest ", 400)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostBadRequest ", 400)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockUnlockActionResourcePostUnlockPostUnauthorized creates a SmartlockUnlockActionResourcePostUnlockPostUnauthorized with default headers values
func NewSmartlockUnlockActionResourcePostUnlockPostUnauthorized() *SmartlockUnlockActionResourcePostUnlockPostUnauthorized {
	return &SmartlockUnlockActionResourcePostUnlockPostUnauthorized{}
}

/*
SmartlockUnlockActionResourcePostUnlockPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockUnlockActionResourcePostUnlockPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock unlock action resource post unlock post unauthorized response has a 2xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock unlock action resource post unlock post unauthorized response has a 3xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock unlock action resource post unlock post unauthorized response has a 4xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock unlock action resource post unlock post unauthorized response has a 5xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock unlock action resource post unlock post unauthorized response a status code equal to that given
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock unlock action resource post unlock post unauthorized response
func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostUnauthorized ", 401)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostUnauthorized ", 401)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed creates a SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed with default headers values
func NewSmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed() *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed {
	return &SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed{}
}

/*
SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed describes a response with status code 405, with default header values.

Not allowed
*/
type SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed struct {
}

// IsSuccess returns true when this smartlock unlock action resource post unlock post method not allowed response has a 2xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock unlock action resource post unlock post method not allowed response has a 3xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock unlock action resource post unlock post method not allowed response has a 4xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock unlock action resource post unlock post method not allowed response has a 5xx status code
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock unlock action resource post unlock post method not allowed response a status code equal to that given
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the smartlock unlock action resource post unlock post method not allowed response
func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) Code() int {
	return 405
}

func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostMethodNotAllowed ", 405)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/action/unlock][%d] smartlockUnlockActionResourcePostUnlockPostMethodNotAllowed ", 405)
}

func (o *SmartlockUnlockActionResourcePostUnlockPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
