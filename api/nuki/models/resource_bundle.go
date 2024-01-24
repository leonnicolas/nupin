// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceBundle resource bundle
//
// swagger:model ResourceBundle
type ResourceBundle struct {

	// base bundle name
	BaseBundleName string `json:"baseBundleName,omitempty"`

	// keys
	Keys EnumerationString `json:"keys,omitempty"`

	// locale
	Locale *Locale `json:"locale,omitempty"`
}

// Validate validates this resource bundle
func (m *ResourceBundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceBundle) validateLocale(formats strfmt.Registry) error {
	if swag.IsZero(m.Locale) { // not required
		return nil
	}

	if m.Locale != nil {
		if err := m.Locale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locale")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resource bundle based on the context it is used
func (m *ResourceBundle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceBundle) contextValidateLocale(ctx context.Context, formats strfmt.Registry) error {

	if m.Locale != nil {

		if swag.IsZero(m.Locale) { // not required
			return nil
		}

		if err := m.Locale.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locale")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceBundle) UnmarshalBinary(b []byte) error {
	var res ResourceBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
