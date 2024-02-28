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

// AccountSubscription account subscription
//
// swagger:model AccountSubscription
type AccountSubscription struct {

	// The account id
	// Required: true
	AccountID *int32 `json:"accountId"`

	// The optional gift article after first purchase
	GiftArticle string `json:"giftArticle,omitempty"`

	// The id
	// Required: true
	ID *ObjectID `json:"id"`

	// The next payment date
	NextPaymentDate int64 `json:"nextPaymentDate,omitempty"`

	// The payment type
	// Required: true
	// Enum: [free paypal card account]
	PaymentType *string `json:"paymentType"`

	// The actual period
	// Required: true
	Period *int32 `json:"period"`

	// The period end date
	PeriodEndDate int64 `json:"periodEndDate,omitempty"`

	// The quantity of authorizations
	// Required: true
	Quantity *int32 `json:"quantity"`

	// The start date
	StartDate int64 `json:"startDate,omitempty"`

	// The status
	// Required: true
	// Enum: [active deactivated finished]
	Status *string `json:"status"`

	// The subscription id
	// Required: true
	SubscriptionID *int32 `json:"subscriptionId"`
}

// Validate validates this account subscription
func (m *AccountSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountSubscription) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *AccountSubscription) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

var accountSubscriptionTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["free","paypal","card","account"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountSubscriptionTypePaymentTypePropEnum = append(accountSubscriptionTypePaymentTypePropEnum, v)
	}
}

const (

	// AccountSubscriptionPaymentTypeFree captures enum value "free"
	AccountSubscriptionPaymentTypeFree string = "free"

	// AccountSubscriptionPaymentTypePaypal captures enum value "paypal"
	AccountSubscriptionPaymentTypePaypal string = "paypal"

	// AccountSubscriptionPaymentTypeCard captures enum value "card"
	AccountSubscriptionPaymentTypeCard string = "card"

	// AccountSubscriptionPaymentTypeAccount captures enum value "account"
	AccountSubscriptionPaymentTypeAccount string = "account"
)

// prop value enum
func (m *AccountSubscription) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountSubscriptionTypePaymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountSubscription) validatePaymentType(formats strfmt.Registry) error {

	if err := validate.Required("paymentType", "body", m.PaymentType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentTypeEnum("paymentType", "body", *m.PaymentType); err != nil {
		return err
	}

	return nil
}

func (m *AccountSubscription) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *AccountSubscription) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

var accountSubscriptionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","deactivated","finished"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountSubscriptionTypeStatusPropEnum = append(accountSubscriptionTypeStatusPropEnum, v)
	}
}

const (

	// AccountSubscriptionStatusActive captures enum value "active"
	AccountSubscriptionStatusActive string = "active"

	// AccountSubscriptionStatusDeactivated captures enum value "deactivated"
	AccountSubscriptionStatusDeactivated string = "deactivated"

	// AccountSubscriptionStatusFinished captures enum value "finished"
	AccountSubscriptionStatusFinished string = "finished"
)

// prop value enum
func (m *AccountSubscription) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountSubscriptionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountSubscription) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountSubscription) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account subscription based on the context it is used
func (m *AccountSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountSubscription) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {

		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountSubscription) UnmarshalBinary(b []byte) error {
	var res AccountSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
