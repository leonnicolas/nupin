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
	"github.com/go-openapi/validate"
)

// Representation representation
//
// swagger:model Representation
type Representation struct {

	// available
	Available bool `json:"available,omitempty"`

	// available size
	AvailableSize int64 `json:"availableSize,omitempty"`

	// channel
	Channel *ReadableByteChannel `json:"channel,omitempty"`

	// character set
	CharacterSet *CharacterSet `json:"characterSet,omitempty"`

	// digest
	Digest *Digest `json:"digest,omitempty"`

	// disposition
	Disposition *Disposition `json:"disposition,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// encodings
	Encodings []*Encoding `json:"encodings"`

	// expiration date
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// languages
	Languages []*Language `json:"languages"`

	// location ref
	LocationRef *Reference `json:"locationRef,omitempty"`

	// media type
	MediaType *MediaType `json:"mediaType,omitempty"`

	// modification date
	// Format: date-time
	ModificationDate strfmt.DateTime `json:"modificationDate,omitempty"`

	// range
	Range *Range `json:"range,omitempty"`

	// reader
	Reader Reader `json:"reader,omitempty"`

	// registration
	Registration *SelectionRegistration `json:"registration,omitempty"`

	// selectable
	Selectable bool `json:"selectable,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// stream
	Stream InputStream `json:"stream,omitempty"`

	// tag
	Tag *Tag `json:"tag,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// transient
	Transient bool `json:"transient,omitempty"`
}

// Validate validates this representation
func (m *Representation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCharacterSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncodings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Representation) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateCharacterSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CharacterSet) { // not required
		return nil
	}

	if m.CharacterSet != nil {
		if err := m.CharacterSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("characterSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("characterSet")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateDigest(formats strfmt.Registry) error {
	if swag.IsZero(m.Digest) { // not required
		return nil
	}

	if m.Digest != nil {
		if err := m.Digest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digest")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateDisposition(formats strfmt.Registry) error {
	if swag.IsZero(m.Disposition) { // not required
		return nil
	}

	if m.Disposition != nil {
		if err := m.Disposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disposition")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateEncodings(formats strfmt.Registry) error {
	if swag.IsZero(m.Encodings) { // not required
		return nil
	}

	for i := 0; i < len(m.Encodings); i++ {
		if swag.IsZero(m.Encodings[i]) { // not required
			continue
		}

		if m.Encodings[i] != nil {
			if err := m.Encodings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encodings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encodings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Representation) validateExpirationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Representation) validateLanguages(formats strfmt.Registry) error {
	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Representation) validateLocationRef(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationRef) { // not required
		return nil
	}

	if m.LocationRef != nil {
		if err := m.LocationRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationRef")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	if m.MediaType != nil {
		if err := m.MediaType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateModificationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Representation) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateRegistration(formats strfmt.Registry) error {
	if swag.IsZero(m.Registration) { // not required
		return nil
	}

	if m.Registration != nil {
		if err := m.Registration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) validateTag(formats strfmt.Registry) error {
	if swag.IsZero(m.Tag) { // not required
		return nil
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this representation based on the context it is used
func (m *Representation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCharacterSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisposition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncodings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Representation) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.Channel != nil {

		if swag.IsZero(m.Channel) { // not required
			return nil
		}

		if err := m.Channel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateCharacterSet(ctx context.Context, formats strfmt.Registry) error {

	if m.CharacterSet != nil {

		if swag.IsZero(m.CharacterSet) { // not required
			return nil
		}

		if err := m.CharacterSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("characterSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("characterSet")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateDigest(ctx context.Context, formats strfmt.Registry) error {

	if m.Digest != nil {

		if swag.IsZero(m.Digest) { // not required
			return nil
		}

		if err := m.Digest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digest")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateDisposition(ctx context.Context, formats strfmt.Registry) error {

	if m.Disposition != nil {

		if swag.IsZero(m.Disposition) { // not required
			return nil
		}

		if err := m.Disposition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disposition")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateEncodings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Encodings); i++ {

		if m.Encodings[i] != nil {

			if swag.IsZero(m.Encodings[i]) { // not required
				return nil
			}

			if err := m.Encodings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encodings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encodings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Representation) contextValidateLanguages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Languages); i++ {

		if m.Languages[i] != nil {

			if swag.IsZero(m.Languages[i]) { // not required
				return nil
			}

			if err := m.Languages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Representation) contextValidateLocationRef(ctx context.Context, formats strfmt.Registry) error {

	if m.LocationRef != nil {

		if swag.IsZero(m.LocationRef) { // not required
			return nil
		}

		if err := m.LocationRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationRef")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateMediaType(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaType != nil {

		if swag.IsZero(m.MediaType) { // not required
			return nil
		}

		if err := m.MediaType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.Range != nil {

		if swag.IsZero(m.Range) { // not required
			return nil
		}

		if err := m.Range.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateRegistration(ctx context.Context, formats strfmt.Registry) error {

	if m.Registration != nil {

		if swag.IsZero(m.Registration) { // not required
			return nil
		}

		if err := m.Registration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration")
			}
			return err
		}
	}

	return nil
}

func (m *Representation) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if m.Tag != nil {

		if swag.IsZero(m.Tag) { // not required
			return nil
		}

		if err := m.Tag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Representation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Representation) UnmarshalBinary(b []byte) error {
	var res Representation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}