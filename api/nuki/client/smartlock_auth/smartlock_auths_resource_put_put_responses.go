// Code generated by go-swagger; DO NOT EDIT.

package smartlock_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockAuthsResourcePutPutReader is a Reader for the SmartlockAuthsResourcePutPut structure.
type SmartlockAuthsResourcePutPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockAuthsResourcePutPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockAuthsResourcePutPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockAuthsResourcePutPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewSmartlockAuthsResourcePutPutPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSmartlockAuthsResourcePutPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 426:
		result := NewSmartlockAuthsResourcePutPutUpgradeRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /smartlock/{smartlockId}/auth] SmartlockAuthsResource_put_put", response, response.Code())
	}
}

// NewSmartlockAuthsResourcePutPutNoContent creates a SmartlockAuthsResourcePutPutNoContent with default headers values
func NewSmartlockAuthsResourcePutPutNoContent() *SmartlockAuthsResourcePutPutNoContent {
	return &SmartlockAuthsResourcePutPutNoContent{}
}

/*
SmartlockAuthsResourcePutPutNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockAuthsResourcePutPutNoContent struct {
}

// IsSuccess returns true when this smartlock auths resource put put no content response has a 2xx status code
func (o *SmartlockAuthsResourcePutPutNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock auths resource put put no content response has a 3xx status code
func (o *SmartlockAuthsResourcePutPutNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auths resource put put no content response has a 4xx status code
func (o *SmartlockAuthsResourcePutPutNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock auths resource put put no content response has a 5xx status code
func (o *SmartlockAuthsResourcePutPutNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auths resource put put no content response a status code equal to that given
func (o *SmartlockAuthsResourcePutPutNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock auths resource put put no content response
func (o *SmartlockAuthsResourcePutPutNoContent) Code() int {
	return 204
}

func (o *SmartlockAuthsResourcePutPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutNoContent ", 204)
}

func (o *SmartlockAuthsResourcePutPutNoContent) String() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutNoContent ", 204)
}

func (o *SmartlockAuthsResourcePutPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthsResourcePutPutBadRequest creates a SmartlockAuthsResourcePutPutBadRequest with default headers values
func NewSmartlockAuthsResourcePutPutBadRequest() *SmartlockAuthsResourcePutPutBadRequest {
	return &SmartlockAuthsResourcePutPutBadRequest{}
}

/*
SmartlockAuthsResourcePutPutBadRequest describes a response with status code 400, with default header values.

Bad parameter
*/
type SmartlockAuthsResourcePutPutBadRequest struct {
}

// IsSuccess returns true when this smartlock auths resource put put bad request response has a 2xx status code
func (o *SmartlockAuthsResourcePutPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auths resource put put bad request response has a 3xx status code
func (o *SmartlockAuthsResourcePutPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auths resource put put bad request response has a 4xx status code
func (o *SmartlockAuthsResourcePutPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auths resource put put bad request response has a 5xx status code
func (o *SmartlockAuthsResourcePutPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auths resource put put bad request response a status code equal to that given
func (o *SmartlockAuthsResourcePutPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock auths resource put put bad request response
func (o *SmartlockAuthsResourcePutPutBadRequest) Code() int {
	return 400
}

func (o *SmartlockAuthsResourcePutPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutBadRequest ", 400)
}

func (o *SmartlockAuthsResourcePutPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutBadRequest ", 400)
}

func (o *SmartlockAuthsResourcePutPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthsResourcePutPutPaymentRequired creates a SmartlockAuthsResourcePutPutPaymentRequired with default headers values
func NewSmartlockAuthsResourcePutPutPaymentRequired() *SmartlockAuthsResourcePutPutPaymentRequired {
	return &SmartlockAuthsResourcePutPutPaymentRequired{}
}

/*
SmartlockAuthsResourcePutPutPaymentRequired describes a response with status code 402, with default header values.

Account not payed
*/
type SmartlockAuthsResourcePutPutPaymentRequired struct {
}

// IsSuccess returns true when this smartlock auths resource put put payment required response has a 2xx status code
func (o *SmartlockAuthsResourcePutPutPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auths resource put put payment required response has a 3xx status code
func (o *SmartlockAuthsResourcePutPutPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auths resource put put payment required response has a 4xx status code
func (o *SmartlockAuthsResourcePutPutPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auths resource put put payment required response has a 5xx status code
func (o *SmartlockAuthsResourcePutPutPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auths resource put put payment required response a status code equal to that given
func (o *SmartlockAuthsResourcePutPutPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the smartlock auths resource put put payment required response
func (o *SmartlockAuthsResourcePutPutPaymentRequired) Code() int {
	return 402
}

func (o *SmartlockAuthsResourcePutPutPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutPaymentRequired ", 402)
}

func (o *SmartlockAuthsResourcePutPutPaymentRequired) String() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutPaymentRequired ", 402)
}

func (o *SmartlockAuthsResourcePutPutPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthsResourcePutPutConflict creates a SmartlockAuthsResourcePutPutConflict with default headers values
func NewSmartlockAuthsResourcePutPutConflict() *SmartlockAuthsResourcePutPutConflict {
	return &SmartlockAuthsResourcePutPutConflict{}
}

/*
SmartlockAuthsResourcePutPutConflict describes a response with status code 409, with default header values.

Parameter conflicts
*/
type SmartlockAuthsResourcePutPutConflict struct {
}

// IsSuccess returns true when this smartlock auths resource put put conflict response has a 2xx status code
func (o *SmartlockAuthsResourcePutPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auths resource put put conflict response has a 3xx status code
func (o *SmartlockAuthsResourcePutPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auths resource put put conflict response has a 4xx status code
func (o *SmartlockAuthsResourcePutPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auths resource put put conflict response has a 5xx status code
func (o *SmartlockAuthsResourcePutPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auths resource put put conflict response a status code equal to that given
func (o *SmartlockAuthsResourcePutPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the smartlock auths resource put put conflict response
func (o *SmartlockAuthsResourcePutPutConflict) Code() int {
	return 409
}

func (o *SmartlockAuthsResourcePutPutConflict) Error() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutConflict ", 409)
}

func (o *SmartlockAuthsResourcePutPutConflict) String() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutConflict ", 409)
}

func (o *SmartlockAuthsResourcePutPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockAuthsResourcePutPutUpgradeRequired creates a SmartlockAuthsResourcePutPutUpgradeRequired with default headers values
func NewSmartlockAuthsResourcePutPutUpgradeRequired() *SmartlockAuthsResourcePutPutUpgradeRequired {
	return &SmartlockAuthsResourcePutPutUpgradeRequired{}
}

/*
SmartlockAuthsResourcePutPutUpgradeRequired describes a response with status code 426, with default header values.

Account upgrade required
*/
type SmartlockAuthsResourcePutPutUpgradeRequired struct {
}

// IsSuccess returns true when this smartlock auths resource put put upgrade required response has a 2xx status code
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock auths resource put put upgrade required response has a 3xx status code
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock auths resource put put upgrade required response has a 4xx status code
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock auths resource put put upgrade required response has a 5xx status code
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock auths resource put put upgrade required response a status code equal to that given
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) IsCode(code int) bool {
	return code == 426
}

// Code gets the status code for the smartlock auths resource put put upgrade required response
func (o *SmartlockAuthsResourcePutPutUpgradeRequired) Code() int {
	return 426
}

func (o *SmartlockAuthsResourcePutPutUpgradeRequired) Error() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutUpgradeRequired ", 426)
}

func (o *SmartlockAuthsResourcePutPutUpgradeRequired) String() string {
	return fmt.Sprintf("[PUT /smartlock/{smartlockId}/auth][%d] smartlockAuthsResourcePutPutUpgradeRequired ", 426)
}

func (o *SmartlockAuthsResourcePutPutUpgradeRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
