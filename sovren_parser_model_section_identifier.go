// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelSectionIdentifier sovren parser model section identifier
//
// swagger:model Sovren.Parser.Model.SectionIdentifier
type SovrenParserModelSectionIdentifier struct {

	// Id
	ID string `json:"Id,omitempty"`

	// section type
	SectionType string `json:"SectionType,omitempty"`
}

// Validate validates this sovren parser model section identifier
func (m *SovrenParserModelSectionIdentifier) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model section identifier based on context it is used
func (m *SovrenParserModelSectionIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelSectionIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelSectionIdentifier) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelSectionIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}