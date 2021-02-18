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
	"github.com/go-openapi/validate"
)

// SovrenParserModelResumeResumeMetadata sovren parser model resume resume metadata
//
// swagger:model Sovren.Parser.Model.Resume.ResumeMetadata
type SovrenParserModelResumeResumeMetadata struct {

	// document culture
	DocumentCulture string `json:"DocumentCulture,omitempty"`

	// document language
	DocumentLanguage string `json:"DocumentLanguage,omitempty"`

	// document last modified
	// Format: date
	DocumentLastModified strfmt.Date `json:"DocumentLastModified,omitempty"`

	// found sections
	FoundSections []*SovrenParserModelResumeResumeSection `json:"FoundSections"`

	// parser settings
	ParserSettings string `json:"ParserSettings,omitempty"`

	// plain text
	PlainText string `json:"PlainText,omitempty"`

	// reserved data
	ReservedData *SovrenParserModelResumeReservedData `json:"ReservedData,omitempty"`

	// resume quality
	ResumeQuality []*SovrenParserModelResumeResumeQualityAssessment `json:"ResumeQuality"`

	// sovren signature
	SovrenSignature []string `json:"SovrenSignature"`
}

// Validate validates this sovren parser model resume resume metadata
func (m *SovrenParserModelResumeResumeMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFoundSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservedData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResumeQuality(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) validateDocumentLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DocumentLastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("DocumentLastModified", "body", "date", m.DocumentLastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) validateFoundSections(formats strfmt.Registry) error {
	if swag.IsZero(m.FoundSections) { // not required
		return nil
	}

	for i := 0; i < len(m.FoundSections); i++ {
		if swag.IsZero(m.FoundSections[i]) { // not required
			continue
		}

		if m.FoundSections[i] != nil {
			if err := m.FoundSections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FoundSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) validateReservedData(formats strfmt.Registry) error {
	if swag.IsZero(m.ReservedData) { // not required
		return nil
	}

	if m.ReservedData != nil {
		if err := m.ReservedData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReservedData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) validateResumeQuality(formats strfmt.Registry) error {
	if swag.IsZero(m.ResumeQuality) { // not required
		return nil
	}

	for i := 0; i < len(m.ResumeQuality); i++ {
		if swag.IsZero(m.ResumeQuality[i]) { // not required
			continue
		}

		if m.ResumeQuality[i] != nil {
			if err := m.ResumeQuality[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResumeQuality" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model resume resume metadata based on the context it is used
func (m *SovrenParserModelResumeResumeMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFoundSections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReservedData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResumeQuality(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) contextValidateFoundSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FoundSections); i++ {

		if m.FoundSections[i] != nil {
			if err := m.FoundSections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FoundSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) contextValidateReservedData(ctx context.Context, formats strfmt.Registry) error {

	if m.ReservedData != nil {
		if err := m.ReservedData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReservedData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeMetadata) contextValidateResumeQuality(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResumeQuality); i++ {

		if m.ResumeQuality[i] != nil {
			if err := m.ResumeQuality[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResumeQuality" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeMetadata) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeResumeMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}