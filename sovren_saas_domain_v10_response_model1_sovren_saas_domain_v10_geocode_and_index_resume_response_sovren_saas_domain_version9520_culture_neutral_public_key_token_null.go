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

// SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull sovren saas domain v10 response model 1 sovren saas domain v10 geocode and index resume response sovren saas domain version 9 5 2 0 culture neutral public key token null
//
// swagger:model Sovren.Saas.Domain.V10ResponseModel`1[[Sovren.Saas.Domain.V10.GeocodeAndIndexResumeResponse, Sovren.Saas.Domain, Version=9.5.2.0, Culture=neutral, PublicKeyToken=null]]
type SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull struct {

	// info
	Info *SovrenSaasDomainV10ResponseInfoModel `json:"Info,omitempty"`

	// value
	Value *SovrenSaasDomainV10GeocodeAndIndexResumeResponse `json:"Value,omitempty"`
}

// Validate validates this sovren saas domain v10 response model 1 sovren saas domain v10 geocode and index resume response sovren saas domain version 9 5 2 0 culture neutral public key token null
func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) validateInfo(formats strfmt.Registry) error {
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

func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 response model 1 sovren saas domain v10 geocode and index resume response sovren saas domain version 9 5 2 0 culture neutral public key token null based on the context it is used
func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
