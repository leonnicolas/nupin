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

// SmartlocksAuthCreate smartlocks auth create
//
// swagger:model SmartlocksAuthCreate
type SmartlocksAuthCreate struct {

	// The id of the linked account user (required if type is NOT 13 .. keypad)
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

	// The name of the authorization (max 32 chars)
	// Required: true
	Name *string `json:"name"`

	// True if the auth has remote access
	// Required: true
	RemoteAllowed *bool `json:"remoteAllowed"`

	// The smart actions enabled flag
	SmartActionsEnabled bool `json:"smartActionsEnabled,omitempty"`

	// The list of smartlock ids
	SmartlockIds []int64 `json:"smartlockIds"`

	// The optional type of the auth 0 .. app (default), 2 .. fob, 13 .. keypad
	Type int32 `json:"type,omitempty"`
}

// Validate validates this smartlocks auth create
func (m *SmartlocksAuthCreate) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRemoteAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartlocksAuthCreate) validateAllowedFromDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedFromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("allowedFromDate", "body", "date-time", m.AllowedFromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmartlocksAuthCreate) validateAllowedUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("allowedUntilDate", "body", "date-time", m.AllowedUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmartlocksAuthCreate) validateAllowedWeekDays(formats strfmt.Registry) error {
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

func (m *SmartlocksAuthCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SmartlocksAuthCreate) validateRemoteAllowed(formats strfmt.Registry) error {

	if err := validate.Required("remoteAllowed", "body", m.RemoteAllowed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this smartlocks auth create based on context it is used
func (m *SmartlocksAuthCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmartlocksAuthCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartlocksAuthCreate) UnmarshalBinary(b []byte) error {
	var res SmartlocksAuthCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}