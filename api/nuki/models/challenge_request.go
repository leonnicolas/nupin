// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChallengeRequest challenge request
//
// swagger:model ChallengeRequest
type ChallengeRequest struct {

	// digest algorithm
	DigestAlgorithm string `json:"digestAlgorithm,omitempty"`

	// domain refs
	DomainRefs []*Reference `json:"domainRefs"`

	// opaque
	Opaque string `json:"opaque,omitempty"`

	// parameters
	Parameters []*Parameter `json:"parameters"`

	// quality options
	QualityOptions []string `json:"qualityOptions"`

	// raw value
	RawValue string `json:"rawValue,omitempty"`

	// realm
	Realm string `json:"realm,omitempty"`

	// scheme
	Scheme *ChallengeScheme `json:"scheme,omitempty"`

	// server nonce
	ServerNonce string `json:"serverNonce,omitempty"`

	// stale
	Stale bool `json:"stale,omitempty"`
}

// Validate validates this challenge request
func (m *ChallengeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChallengeRequest) validateDomainRefs(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.DomainRefs); i++ {
		if swag.IsZero(m.DomainRefs[i]) { // not required
			continue
		}

		if m.DomainRefs[i] != nil {
			if err := m.DomainRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domainRefs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domainRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChallengeRequest) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChallengeRequest) validateScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheme) { // not required
		return nil
	}

	if m.Scheme != nil {
		if err := m.Scheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheme")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this challenge request based on the context it is used
func (m *ChallengeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomainRefs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChallengeRequest) contextValidateDomainRefs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DomainRefs); i++ {

		if m.DomainRefs[i] != nil {

			if swag.IsZero(m.DomainRefs[i]) { // not required
				return nil
			}

			if err := m.DomainRefs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domainRefs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domainRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChallengeRequest) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChallengeRequest) contextValidateScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.Scheme != nil {

		if swag.IsZero(m.Scheme) { // not required
			return nil
		}

		if err := m.Scheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheme")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChallengeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChallengeRequest) UnmarshalBinary(b []byte) error {
	var res ChallengeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
