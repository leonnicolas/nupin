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

// OpenerIntercomModel opener intercom model
//
// swagger:model OpenerIntercomModel
type OpenerIntercomModel struct {

	// The related brand ID
	// Required: true
	BrandID *int32 `json:"brandId"`

	// Settings value for BUS mode switch
	// Required: true
	BusModeSwitch *int32 `json:"busModeSwitch"`

	// Settings value for BUS mode switch short cicuit duration
	// Required: true
	BusModeSwitchShortCircuitDuration *int32 `json:"busModeSwitchShortCircuitDuration"`

	// Connection for audio out
	// Required: true
	ConAudioout *string `json:"conAudioout"`

	// Connection for audio BUS
	// Required: true
	ConBusAudio *string `json:"conBusAudio"`

	// Connection for doorbell minus
	// Required: true
	ConDoorbellMinus *string `json:"conDoorbellMinus"`

	// Connection for doorbell plus
	// Required: true
	ConDoorbellPlus *string `json:"conDoorbellPlus"`

	// Connection for ground analogue
	// Required: true
	ConGndAnalogue *string `json:"conGndAnalogue"`

	// Connection for ground BUS
	// Required: true
	ConGndBus *string `json:"conGndBus"`

	// Connection for open the door
	// Required: true
	ConOpendoor *string `json:"conOpendoor"`

	// The creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The intercom ID
	// Required: true
	IntercomID *int32 `json:"intercomId"`

	// The model name
	// Required: true
	Model *string `json:"model"`

	// The type of the model
	// Required: true
	Type *int32 `json:"type"`

	// The update date
	// Format: date-time
	UpdateDate strfmt.DateTime `json:"updateDate,omitempty"`

	// Verified Nuki intercom: 1 .. verified to work, 2 .. may be compatible, but not verified, 3 .. not compatible
	// Required: true
	Verified *int32 `json:"verified"`
}

// Validate validates this opener intercom model
func (m *OpenerIntercomModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusModeSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusModeSwitchShortCircuitDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConAudioout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConBusAudio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConDoorbellMinus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConDoorbellPlus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConGndAnalogue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConGndBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConOpendoor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntercomID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenerIntercomModel) validateBrandID(formats strfmt.Registry) error {

	if err := validate.Required("brandId", "body", m.BrandID); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateBusModeSwitch(formats strfmt.Registry) error {

	if err := validate.Required("busModeSwitch", "body", m.BusModeSwitch); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateBusModeSwitchShortCircuitDuration(formats strfmt.Registry) error {

	if err := validate.Required("busModeSwitchShortCircuitDuration", "body", m.BusModeSwitchShortCircuitDuration); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConAudioout(formats strfmt.Registry) error {

	if err := validate.Required("conAudioout", "body", m.ConAudioout); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConBusAudio(formats strfmt.Registry) error {

	if err := validate.Required("conBusAudio", "body", m.ConBusAudio); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConDoorbellMinus(formats strfmt.Registry) error {

	if err := validate.Required("conDoorbellMinus", "body", m.ConDoorbellMinus); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConDoorbellPlus(formats strfmt.Registry) error {

	if err := validate.Required("conDoorbellPlus", "body", m.ConDoorbellPlus); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConGndAnalogue(formats strfmt.Registry) error {

	if err := validate.Required("conGndAnalogue", "body", m.ConGndAnalogue); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConGndBus(formats strfmt.Registry) error {

	if err := validate.Required("conGndBus", "body", m.ConGndBus); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateConOpendoor(formats strfmt.Registry) error {

	if err := validate.Required("conOpendoor", "body", m.ConOpendoor); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateIntercomID(formats strfmt.Registry) error {

	if err := validate.Required("intercomId", "body", m.IntercomID); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateUpdateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updateDate", "body", "date-time", m.UpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenerIntercomModel) validateVerified(formats strfmt.Registry) error {

	if err := validate.Required("verified", "body", m.Verified); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this opener intercom model based on context it is used
func (m *OpenerIntercomModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenerIntercomModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenerIntercomModel) UnmarshalBinary(b []byte) error {
	var res OpenerIntercomModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}