// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/leonnicolas/nupin/api/nuki/models"
)

// NotificationsResourceGetGetReader is a Reader for the NotificationsResourceGetGet structure.
type NotificationsResourceGetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsResourceGetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsResourceGetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewNotificationsResourceGetGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /notification] NotificationsResource_get_get", response, response.Code())
	}
}

// NewNotificationsResourceGetGetOK creates a NotificationsResourceGetGetOK with default headers values
func NewNotificationsResourceGetGetOK() *NotificationsResourceGetGetOK {
	return &NotificationsResourceGetGetOK{}
}

/*
NotificationsResourceGetGetOK describes a response with status code 200, with default header values.

successful operation
*/
type NotificationsResourceGetGetOK struct {
	Payload []*models.Notification
}

// IsSuccess returns true when this notifications resource get get o k response has a 2xx status code
func (o *NotificationsResourceGetGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications resource get get o k response has a 3xx status code
func (o *NotificationsResourceGetGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications resource get get o k response has a 4xx status code
func (o *NotificationsResourceGetGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications resource get get o k response has a 5xx status code
func (o *NotificationsResourceGetGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications resource get get o k response a status code equal to that given
func (o *NotificationsResourceGetGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notifications resource get get o k response
func (o *NotificationsResourceGetGetOK) Code() int {
	return 200
}

func (o *NotificationsResourceGetGetOK) Error() string {
	return fmt.Sprintf("[GET /notification][%d] notificationsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *NotificationsResourceGetGetOK) String() string {
	return fmt.Sprintf("[GET /notification][%d] notificationsResourceGetGetOK  %+v", 200, o.Payload)
}

func (o *NotificationsResourceGetGetOK) GetPayload() []*models.Notification {
	return o.Payload
}

func (o *NotificationsResourceGetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsResourceGetGetUnauthorized creates a NotificationsResourceGetGetUnauthorized with default headers values
func NewNotificationsResourceGetGetUnauthorized() *NotificationsResourceGetGetUnauthorized {
	return &NotificationsResourceGetGetUnauthorized{}
}

/*
NotificationsResourceGetGetUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type NotificationsResourceGetGetUnauthorized struct {
}

// IsSuccess returns true when this notifications resource get get unauthorized response has a 2xx status code
func (o *NotificationsResourceGetGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications resource get get unauthorized response has a 3xx status code
func (o *NotificationsResourceGetGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications resource get get unauthorized response has a 4xx status code
func (o *NotificationsResourceGetGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications resource get get unauthorized response has a 5xx status code
func (o *NotificationsResourceGetGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications resource get get unauthorized response a status code equal to that given
func (o *NotificationsResourceGetGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notifications resource get get unauthorized response
func (o *NotificationsResourceGetGetUnauthorized) Code() int {
	return 401
}

func (o *NotificationsResourceGetGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /notification][%d] notificationsResourceGetGetUnauthorized ", 401)
}

func (o *NotificationsResourceGetGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /notification][%d] notificationsResourceGetGetUnauthorized ", 401)
}

func (o *NotificationsResourceGetGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
