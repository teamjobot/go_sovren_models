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

// SovrenParserModelResumeEducationDetails sovren parser model resume education details
//
// swagger:model Sovren.Parser.Model.Resume.EducationDetails
type SovrenParserModelResumeEducationDetails struct {

	// degree
	Degree *SovrenParserModelResumeDegree `json:"Degree,omitempty"`

	// g p a
	GPA *SovrenParserModelResumeGradepointAverage `json:"GPA,omitempty"`

	// graduated
	Graduated *SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"Graduated,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last education date
	LastEducationDate *SovrenParserModelSovrenDateWithParts `json:"LastEducationDate,omitempty"`

	// location
	Location *SovrenParserModelSovrenLocation `json:"Location,omitempty"`

	// majors
	Majors []string `json:"Majors"`

	// minors
	Minors []string `json:"Minors"`

	// school name
	SchoolName *SovrenParserModelNameValue `json:"SchoolName,omitempty"`

	// school type
	SchoolType string `json:"SchoolType,omitempty"`

	// text
	Text string `json:"Text,omitempty"`
}

// Validate validates this sovren parser model resume education details
func (m *SovrenParserModelResumeEducationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDegree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGPA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraduated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEducationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchoolName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeEducationDetails) validateDegree(formats strfmt.Registry) error {
	if swag.IsZero(m.Degree) { // not required
		return nil
	}

	if m.Degree != nil {
		if err := m.Degree.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Degree")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) validateGPA(formats strfmt.Registry) error {
	if swag.IsZero(m.GPA) { // not required
		return nil
	}

	if m.GPA != nil {
		if err := m.GPA.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GPA")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) validateGraduated(formats strfmt.Registry) error {
	if swag.IsZero(m.Graduated) { // not required
		return nil
	}

	if m.Graduated != nil {
		if err := m.Graduated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Graduated")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) validateLastEducationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastEducationDate) { // not required
		return nil
	}

	if m.LastEducationDate != nil {
		if err := m.LastEducationDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastEducationDate")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) validateLocation(formats strfmt.Registry) error {
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

func (m *SovrenParserModelResumeEducationDetails) validateSchoolName(formats strfmt.Registry) error {
	if swag.IsZero(m.SchoolName) { // not required
		return nil
	}

	if m.SchoolName != nil {
		if err := m.SchoolName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchoolName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren parser model resume education details based on the context it is used
func (m *SovrenParserModelResumeEducationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDegree(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGPA(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGraduated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastEducationDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchoolName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeEducationDetails) contextValidateDegree(ctx context.Context, formats strfmt.Registry) error {

	if m.Degree != nil {
		if err := m.Degree.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Degree")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) contextValidateGPA(ctx context.Context, formats strfmt.Registry) error {

	if m.GPA != nil {
		if err := m.GPA.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GPA")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) contextValidateGraduated(ctx context.Context, formats strfmt.Registry) error {

	if m.Graduated != nil {
		if err := m.Graduated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Graduated")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) contextValidateLastEducationDate(ctx context.Context, formats strfmt.Registry) error {

	if m.LastEducationDate != nil {
		if err := m.LastEducationDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastEducationDate")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeEducationDetails) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenParserModelResumeEducationDetails) contextValidateSchoolName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchoolName != nil {
		if err := m.SchoolName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchoolName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeEducationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeEducationDetails) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeEducationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
