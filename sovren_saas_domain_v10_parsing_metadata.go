// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainV10ParsingMetadata sovren saas domain v10 parsing metadata
//
// swagger:model Sovren.Saas.Domain.V10.ParsingMetadata
type SovrenSaasDomainV10ParsingMetadata struct {

	// elapsed milliseconds
	ElapsedMilliseconds int32 `json:"ElapsedMilliseconds,omitempty"`

	// license serial number
	LicenseSerialNumber string `json:"LicenseSerialNumber,omitempty"`

	// timed out
	TimedOut bool `json:"TimedOut,omitempty"`

	// timed out at milliseconds
	TimedOutAtMilliseconds *SovrenCommonV10ParserModelsSovrenPrimitive1SystemInt32MscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"TimedOutAtMilliseconds,omitempty"`
}

// Validate validates this sovren saas domain v10 parsing metadata
func (m *SovrenSaasDomainV10ParsingMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimedOutAtMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ParsingMetadata) validateTimedOutAtMilliseconds(formats strfmt.Registry) error {
	if swag.IsZero(m.TimedOutAtMilliseconds) { // not required
		return nil
	}

	if m.TimedOutAtMilliseconds != nil {
		if err := m.TimedOutAtMilliseconds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimedOutAtMilliseconds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 parsing metadata based on the context it is used
func (m *SovrenSaasDomainV10ParsingMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimedOutAtMilliseconds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ParsingMetadata) contextValidateTimedOutAtMilliseconds(ctx context.Context, formats strfmt.Registry) error {

	if m.TimedOutAtMilliseconds != nil {
		if err := m.TimedOutAtMilliseconds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimedOutAtMilliseconds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10ParsingMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10ParsingMetadata) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10ParsingMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
