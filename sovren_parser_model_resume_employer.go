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

// SovrenParserModelResumeEmployer sovren parser model resume employer
//
// swagger:model Sovren.Parser.Model.Resume.Employer
type SovrenParserModelResumeEmployer struct {

	// location
	Location *SovrenParserModelSovrenLocation `json:"Location,omitempty"`

	// name
	Name *SovrenParserModelResumeCompanyName `json:"Name,omitempty"`

	// other found name
	OtherFoundName *SovrenParserModelNameValue `json:"OtherFoundName,omitempty"`
}

// Validate validates this sovren parser model resume employer
func (m *SovrenParserModelResumeEmployer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherFoundName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeEmployer) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEmployer) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEmployer) validateOtherFoundName(formats strfmt.Registry) error {
	if swag.IsZero(m.OtherFoundName) { // not required
		return nil
	}

	if m.OtherFoundName != nil {
		if err := m.OtherFoundName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OtherFoundName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren parser model resume employer based on the context it is used
func (m *SovrenParserModelResumeEmployer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOtherFoundName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeEmployer) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEmployer) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEmployer) contextValidateOtherFoundName(ctx context.Context, formats strfmt.Registry) error {

	if m.OtherFoundName != nil {
		if err := m.OtherFoundName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OtherFoundName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeEmployer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeEmployer) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeEmployer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
