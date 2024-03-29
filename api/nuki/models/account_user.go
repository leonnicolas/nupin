// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountUser account user
//
// swagger:model AccountUser
type AccountUser struct {

	// The account id
	// Required: true
	AccountID *int32 `json:"accountId"`

	// The account user id
	// Required: true
	AccountUserID *int32 `json:"accountUserId"`

	// The creation date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The email address
	// Required: true
	Email *string `json:"email"`

	// The language code
	// Example: de
	Language string `json:"language,omitempty"`

	// The name
	// Required: true
	Name *string `json:"name"`

	// The operation id - if set it's locked for another operation
	// Read Only: true
	OperationID string `json:"operationId,omitempty"`

	// The optional type: 0 .. user, 1 .. company
	// Maximum: 1
	// Minimum: 0
	Type *int32 `json:"type,omitempty"`

	// The update date
	// Required: true
	// Format: date-time
	UpdateDate *strfmt.DateTime `json:"updateDate"`
}

// Validate validates this account user
func (m *AccountUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountUser) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateAccountUserID(formats strfmt.Registry) error {

	if err := validate.Required("accountUserId", "body", m.AccountUserID); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.MinimumInt("type", "body", int64(*m.Type), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("type", "body", int64(*m.Type), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountUser) validateUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("updateDate", "body", m.UpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account user based on the context it is used
func (m *AccountUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountUser) contextValidateOperationID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "operationId", "body", string(m.OperationID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUser) UnmarshalBinary(b []byte) error {
	var res AccountUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
