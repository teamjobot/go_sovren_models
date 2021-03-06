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

// SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 sovren saas domain v10 response model 1 system byte mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
//
// swagger:model Sovren.Saas.Domain.V10ResponseModel`1[[System.Byte[], mscorlib, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089]]
type SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 struct {

	// info
	Info *SovrenSaasDomainV10ResponseInfoModel `json:"Info,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"Value,omitempty"`
}

// Validate validates this sovren saas domain v10 response model 1 system byte mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) validateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 response model 1 system byte mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089 based on the context it is used
func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Info != nil {
		if err := m.Info.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10ResponseModel1SystemByteMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
