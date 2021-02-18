// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelResumeAssociation sovren parser model resume association
//
// swagger:model Sovren.Parser.Model.Resume.Association
type SovrenParserModelResumeAssociation struct {

	// found in context
	FoundInContext string `json:"FoundInContext,omitempty"`

	// organization
	Organization string `json:"Organization,omitempty"`

	// role
	Role string `json:"Role,omitempty"`
}

// Validate validates this sovren parser model resume association
func (m *SovrenParserModelResumeAssociation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model resume association based on context it is used
func (m *SovrenParserModelResumeAssociation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeAssociation) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
