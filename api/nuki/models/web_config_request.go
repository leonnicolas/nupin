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

// WebConfigRequest web config request
//
// swagger:model WebConfigRequest
type WebConfigRequest struct {

	// smartlock Id
	SmartlockID int64 `json:"smartlockId,omitempty"`

	// web config
	WebConfig *SmartlockWebConfig `json:"webConfig,omitempty"`
}

// Validate validates this web config request
func (m *WebConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebConfigRequest) validateWebConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.WebConfig) { // not required
		return nil
	}

	if m.WebConfig != nil {
		if err := m.WebConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web config request based on the context it is used
func (m *WebConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebConfigRequest) contextValidateWebConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.WebConfig != nil {

		if swag.IsZero(m.WebConfig) { // not required
			return nil
		}

		if err := m.WebConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebConfigRequest) UnmarshalBinary(b []byte) error {
	var res WebConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
