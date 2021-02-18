// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SovrenParserModelJobJobMetadata sovren parser model job job metadata
//
// swagger:model Sovren.Parser.Model.Job.JobMetadata
type SovrenParserModelJobJobMetadata struct {

	// document culture
	DocumentCulture string `json:"DocumentCulture,omitempty"`

	// document language
	DocumentLanguage string `json:"DocumentLanguage,omitempty"`

	// document last modified
	// Format: date
	DocumentLastModified strfmt.Date `json:"DocumentLastModified,omitempty"`

	// parser settings
	ParserSettings string `json:"ParserSettings,omitempty"`

	// plain text
	PlainText string `json:"PlainText,omitempty"`

	// sovren signature
	SovrenSignature []string `json:"SovrenSignature"`
}

// Validate validates this sovren parser model job job metadata
func (m *SovrenParserModelJobJobMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentLastModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelJobJobMetadata) validateDocumentLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DocumentLastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("DocumentLastModified", "body", "date", m.DocumentLastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sovren parser model job job metadata based on context it is used
func (m *SovrenParserModelJobJobMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelJobJobMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelJobJobMetadata) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelJobJobMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
