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

// AccountOtpEnable account otp enable
//
// swagger:model AccountOtpEnable
type AccountOtpEnable struct {

	// The one time password (otp)
	// Required: true
	Otp *string `json:"otp"`
}

// Validate validates this account otp enable
func (m *AccountOtpEnable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountOtpEnable) validateOtp(formats strfmt.Registry) error {

	if err := validate.Required("otp", "body", m.Otp); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account otp enable based on context it is used
func (m *AccountOtpEnable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountOtpEnable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountOtpEnable) UnmarshalBinary(b []byte) error {
	var res AccountOtpEnable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
