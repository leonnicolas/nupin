// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmartlockLog smartlock log
//
// swagger:model SmartlockLog
type SmartlockLog struct {

	// The id of the linked account user
	AccountUserID int32 `json:"accountUserId,omitempty"`

	// The action: 1 .. unlock, 2 .. lock, 3 .. unlatch, 4 .. lock'n'go, 5 .. lock'n'go with unlatch, 208 .. door warning ajar, 209 door warning status mismatch, 224 .. doorbell recognition (only Opener), 240 .. door opened, 241 .. door closed, 242 .. door sensor jammed, 243 .. firmware update, 250 .. door log enabled, 251 .. door log disabled, 252 .. initialization, 253 .. calibration, 254 .. log enabled, 255 .. log disabled
	// Required: true
	// Maximum: 255
	// Minimum: 1
	Action *int32 `json:"action"`

	// The door sensor warning ajar timeout (in minutes, only for action = 208)
	AjarTimeout int32 `json:"ajarTimeout,omitempty"`

	// The id of the linked smartlock auth
	AuthID string `json:"authId,omitempty"`

	// True if it was an auto unlock
	// Required: true
	AutoUnlock *bool `json:"autoUnlock"`

	// The log date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// The device type: 0 .. smartlock and box, 2 .. opener, 3 .. smartdoor
	// Required: true
	// Enum: [0 2 3]
	DeviceType *int32 `json:"deviceType"`

	// In case of any error, it contains the error message
	Error string `json:"error,omitempty"`

	// The unique id for the smartlock log
	// Required: true
	ID *string `json:"id"`

	// The name
	// Required: true
	Name *string `json:"name"`

	// The opener specific log
	OpenerLog *SmartlockLogOpenerLog `json:"openerLog,omitempty"`

	// The smartlock id
	// Required: true
	SmartlockID *int64 `json:"smartlockId"`

	// The source of action: 1 .. Keypad code, 2 .. Fingerprint, 0 .. Default
	Source int32 `json:"source,omitempty"`

	// The completion state: 0 .. Success, 1 .. Motor blocked, 2 .. Canceled, 3 .. Too recent, 4 .. Busy, 5 .. Low motor voltage, 6 .. Clutch failure, 7 .. Motor power failure, 8 .. Incomplete, 9 .. Rejected, 10 .. Rejected night mode, 254 .. Other error, 255 .. Unknown error
	// Required: true
	// Maximum: 255
	// Minimum: 1
	State *int32 `json:"state"`

	// The trigger: 0 .. system, 1 .. manual, 2 .. button, 3 .. automatic, 4 .. web, 5 .. app, 6 .. auto lock, 7 .. accessory, 255 .. keypad
	// Required: true
	// Maximum: 255
	// Minimum: 0
	Trigger *int32 `json:"trigger"`
}

// Validate validates this smartlock log
func (m *SmartlockLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoUnlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenerLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmartlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartlockLog) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	if err := validate.MinimumInt("action", "body", int64(*m.Action), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("action", "body", int64(*m.Action), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateAutoUnlock(formats strfmt.Registry) error {

	if err := validate.Required("autoUnlock", "body", m.AutoUnlock); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var smartlockLogTypeDeviceTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smartlockLogTypeDeviceTypePropEnum = append(smartlockLogTypeDeviceTypePropEnum, v)
	}
}

// prop value enum
func (m *SmartlockLog) validateDeviceTypeEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, smartlockLogTypeDeviceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmartlockLog) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("deviceType", "body", m.DeviceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeviceTypeEnum("deviceType", "body", *m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateOpenerLog(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenerLog) { // not required
		return nil
	}

	if m.OpenerLog != nil {
		if err := m.OpenerLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openerLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openerLog")
			}
			return err
		}
	}

	return nil
}

func (m *SmartlockLog) validateSmartlockID(formats strfmt.Registry) error {

	if err := validate.Required("smartlockId", "body", m.SmartlockID); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.MinimumInt("state", "body", int64(*m.State), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("state", "body", int64(*m.State), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *SmartlockLog) validateTrigger(formats strfmt.Registry) error {

	if err := validate.Required("trigger", "body", m.Trigger); err != nil {
		return err
	}

	if err := validate.MinimumInt("trigger", "body", int64(*m.Trigger), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("trigger", "body", int64(*m.Trigger), 255, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this smartlock log based on the context it is used
func (m *SmartlockLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOpenerLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartlockLog) contextValidateOpenerLog(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenerLog != nil {

		if swag.IsZero(m.OpenerLog) { // not required
			return nil
		}

		if err := m.OpenerLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openerLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openerLog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmartlockLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartlockLog) UnmarshalBinary(b []byte) error {
	var res SmartlockLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
