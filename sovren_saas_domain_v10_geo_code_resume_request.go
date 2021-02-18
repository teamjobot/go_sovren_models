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

// SovrenSaasDomainV10GeoCodeResumeRequest sovren saas domain v10 geo code resume request
//
// swagger:model Sovren.Saas.Domain.V10.GeoCodeResumeRequest
type SovrenSaasDomainV10GeoCodeResumeRequest struct {

	// geo coordinates
	GeoCoordinates *SovrenGeoCoordinates `json:"GeoCoordinates,omitempty"`

	// postal address
	PostalAddress *SovrenPostalAddress `json:"PostalAddress,omitempty"`

	// provider
	Provider string `json:"Provider,omitempty"`

	// provider key
	ProviderKey string `json:"ProviderKey,omitempty"`

	// resume data
	ResumeData *SovrenParserModelResumeResumeData `json:"ResumeData,omitempty"`
}

// Validate validates this sovren saas domain v10 geo code resume request
func (m *SovrenSaasDomainV10GeoCodeResumeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeoCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResumeData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) validateGeoCoordinates(formats strfmt.Registry) error {
	if swag.IsZero(m.GeoCoordinates) { // not required
		return nil
	}

	if m.GeoCoordinates != nil {
		if err := m.GeoCoordinates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GeoCoordinates")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) validatePostalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PostalAddress) { // not required
		return nil
	}

	if m.PostalAddress != nil {
		if err := m.PostalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) validateResumeData(formats strfmt.Registry) error {
	if swag.IsZero(m.ResumeData) { // not required
		return nil
	}

	if m.ResumeData != nil {
		if err := m.ResumeData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResumeData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 geo code resume request based on the context it is used
func (m *SovrenSaasDomainV10GeoCodeResumeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeoCoordinates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResumeData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) contextValidateGeoCoordinates(ctx context.Context, formats strfmt.Registry) error {

	if m.GeoCoordinates != nil {
		if err := m.GeoCoordinates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GeoCoordinates")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) contextValidatePostalAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PostalAddress != nil {
		if err := m.PostalAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenSaasDomainV10GeoCodeResumeRequest) contextValidateResumeData(ctx context.Context, formats strfmt.Registry) error {

	if m.ResumeData != nil {
		if err := m.ResumeData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResumeData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10GeoCodeResumeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10GeoCodeResumeRequest) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10GeoCodeResumeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
