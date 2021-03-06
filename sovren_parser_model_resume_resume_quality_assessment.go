// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelResumeResumeQualityAssessment sovren parser model resume resume quality assessment
//
// swagger:model Sovren.Parser.Model.Resume.ResumeQualityAssessment
type SovrenParserModelResumeResumeQualityAssessment struct {

	// findings
	Findings []*SovrenParserModelResumeFinding `json:"Findings"`

	// level
	Level string `json:"Level,omitempty"`
}

// Validate validates this sovren parser model resume resume quality assessment
func (m *SovrenParserModelResumeResumeQualityAssessment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFindings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeQualityAssessment) validateFindings(formats strfmt.Registry) error {
	if swag.IsZero(m.Findings) { // not required
		return nil
	}

	for i := 0; i < len(m.Findings); i++ {
		if swag.IsZero(m.Findings[i]) { // not required
			continue
		}

		if m.Findings[i] != nil {
			if err := m.Findings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Findings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model resume resume quality assessment based on the context it is used
func (m *SovrenParserModelResumeResumeQualityAssessment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFindings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeQualityAssessment) contextValidateFindings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Findings); i++ {

		if m.Findings[i] != nil {
			if err := m.Findings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Findings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeQualityAssessment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeQualityAssessment) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeResumeQualityAssessment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
