// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationResourceDeleteDeleteReader is a Reader for the NotificationResourceDeleteDelete structure.
type NotificationResourceDeleteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationResourceDeleteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNotificationResourceDeleteDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewNotificationResourceDeleteDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationResourceDeleteDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewNotificationResourceDeleteDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /notification/{notificationId}] NotificationResource_delete_delete", response, response.Code())
	}
}

// NewNotificationResourceDeleteDeleteNoContent creates a NotificationResourceDeleteDeleteNoContent with default headers values
func NewNotificationResourceDeleteDeleteNoContent() *NotificationResourceDeleteDeleteNoContent {
	return &NotificationResourceDeleteDeleteNoContent{}
}

/*
NotificationResourceDeleteDeleteNoContent describes a response with status code 204, with default header values.

Ok
*/
type NotificationResourceDeleteDeleteNoContent struct {
}

// IsSuccess returns true when this notification resource delete delete no content response has a 2xx status code
func (o *NotificationResourceDeleteDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification resource delete delete no content response has a 3xx status code
func (o *NotificationResourceDeleteDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification resource delete delete no content response has a 4xx status code
func (o *NotificationResourceDeleteDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification resource delete delete no content response has a 5xx status code
func (o *NotificationResourceDeleteDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this notification resource delete delete no content response a status code equal to that given
func (o *NotificationResourceDeleteDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the notification resource delete delete no content response
func (o *NotificationResourceDeleteDeleteNoContent) Code() int {
	return 204
}

func (o *NotificationResourceDeleteDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteNoContent ", 204)
}

func (o *NotificationResourceDeleteDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteNoContent ", 204)
}

func (o *NotificationResourceDeleteDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationResourceDeleteDeleteUnauthorized creates a NotificationResourceDeleteDeleteUnauthorized with default headers values
func NewNotificationResourceDeleteDeleteUnauthorized() *NotificationResourceDeleteDeleteUnauthorized {
	return &NotificationResourceDeleteDeleteUnauthorized{}
}

/*
NotificationResourceDeleteDeleteUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type NotificationResourceDeleteDeleteUnauthorized struct {
}

// IsSuccess returns true when this notification resource delete delete unauthorized response has a 2xx status code
func (o *NotificationResourceDeleteDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification resource delete delete unauthorized response has a 3xx status code
func (o *NotificationResourceDeleteDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification resource delete delete unauthorized response has a 4xx status code
func (o *NotificationResourceDeleteDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification resource delete delete unauthorized response has a 5xx status code
func (o *NotificationResourceDeleteDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notification resource delete delete unauthorized response a status code equal to that given
func (o *NotificationResourceDeleteDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notification resource delete delete unauthorized response
func (o *NotificationResourceDeleteDeleteUnauthorized) Code() int {
	return 401
}

func (o *NotificationResourceDeleteDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteUnauthorized ", 401)
}

func (o *NotificationResourceDeleteDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteUnauthorized ", 401)
}

func (o *NotificationResourceDeleteDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationResourceDeleteDeleteForbidden creates a NotificationResourceDeleteDeleteForbidden with default headers values
func NewNotificationResourceDeleteDeleteForbidden() *NotificationResourceDeleteDeleteForbidden {
	return &NotificationResourceDeleteDeleteForbidden{}
}

/*
NotificationResourceDeleteDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationResourceDeleteDeleteForbidden struct {
}

// IsSuccess returns true when this notification resource delete delete forbidden response has a 2xx status code
func (o *NotificationResourceDeleteDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification resource delete delete forbidden response has a 3xx status code
func (o *NotificationResourceDeleteDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification resource delete delete forbidden response has a 4xx status code
func (o *NotificationResourceDeleteDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification resource delete delete forbidden response has a 5xx status code
func (o *NotificationResourceDeleteDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notification resource delete delete forbidden response a status code equal to that given
func (o *NotificationResourceDeleteDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notification resource delete delete forbidden response
func (o *NotificationResourceDeleteDeleteForbidden) Code() int {
	return 403
}

func (o *NotificationResourceDeleteDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteForbidden ", 403)
}

func (o *NotificationResourceDeleteDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteForbidden ", 403)
}

func (o *NotificationResourceDeleteDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationResourceDeleteDeleteMethodNotAllowed creates a NotificationResourceDeleteDeleteMethodNotAllowed with default headers values
func NewNotificationResourceDeleteDeleteMethodNotAllowed() *NotificationResourceDeleteDeleteMethodNotAllowed {
	return &NotificationResourceDeleteDeleteMethodNotAllowed{}
}

/*
NotificationResourceDeleteDeleteMethodNotAllowed describes a response with status code 405, with default header values.

Not allowed
*/
type NotificationResourceDeleteDeleteMethodNotAllowed struct {
}

// IsSuccess returns true when this notification resource delete delete method not allowed response has a 2xx status code
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notification resource delete delete method not allowed response has a 3xx status code
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification resource delete delete method not allowed response has a 4xx status code
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this notification resource delete delete method not allowed response has a 5xx status code
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this notification resource delete delete method not allowed response a status code equal to that given
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the notification resource delete delete method not allowed response
func (o *NotificationResourceDeleteDeleteMethodNotAllowed) Code() int {
	return 405
}

func (o *NotificationResourceDeleteDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteMethodNotAllowed ", 405)
}

func (o *NotificationResourceDeleteDeleteMethodNotAllowed) String() string {
	return fmt.Sprintf("[DELETE /notification/{notificationId}][%d] notificationResourceDeleteDeleteMethodNotAllowed ", 405)
}

func (o *NotificationResourceDeleteDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
