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

// SmartlockAuthUpdate smartlock auth update
//
// swagger:model SmartlockAuthUpdate
type SmartlockAuthUpdate struct {

	// The id of the linked account user
	AccountUserID int32 `json:"accountUserId,omitempty"`

	// The allowed from date
	// Format: date-time
	AllowedFromDate strfmt.DateTime `json:"allowedFromDate,omitempty"`

	// The allowed from time (in minutes from midnight)
	AllowedFromTime int32 `json:"allowedFromTime,omitempty"`

	// The allowed until date
	// Format: date-time
	AllowedUntilDate strfmt.DateTime `json:"allowedUntilDate,omitempty"`

	// The allowed until time (in minutes from midnight)
	AllowedUntilTime int32 `json:"allowedUntilTime,omitempty"`

	// The allowed weekdays bitmask: 64 .. monday, 32 .. tuesday, 16 .. wednesday, 8 .. thursday, 4 .. friday, 2 .. saturday, 1 .. sunday
	// Maximum: 127
	// Minimum: 0
	AllowedWeekDays *int32 `json:"allowedWeekDays,omitempty"`

	// The code of the keypad authorization (only for keypad)
	Code int32 `json:"code,omitempty"`

	// True if the auth is enabled
	Enabled bool `json:"enabled,omitempty"`

	// The name of the authorization (max 32 chars)
	// Required: true
	Name *string `json:"name"`

	// True if the auth has remote access
	RemoteAllowed bool `json:"remoteAllowed,omitempty"`
}

// Validate validates this smartlock auth update
func (m *SmartlockAuthUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedFromDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedUntilDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedWeekDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartlockAuthUpdate) validateAllowedFromDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedFromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("allowedFromDate", "body", "date-time", m.AllowedFromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockAuthUpdate) validateAllowedUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("allowedUntilDate", "body", "date-time", m.AllowedUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockAuthUpdate) validateAllowedWeekDays(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedWeekDays) { // not required
		return nil
	}

	if err := validate.MinimumInt("allowedWeekDays", "body", int64(*m.AllowedWeekDays), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allowedWeekDays", "body", int64(*m.AllowedWeekDays), 127, false); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockAuthUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this smartlock auth update based on context it is used
func (m *SmartlockAuthUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmartlockAuthUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartlockAuthUpdate) UnmarshalBinary(b []byte) error {
	var res SmartlockAuthUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
