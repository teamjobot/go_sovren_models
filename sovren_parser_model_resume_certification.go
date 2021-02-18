// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelResumeCertification sovren parser model resume certification
//
// swagger:model Sovren.Parser.Model.Resume.Certification
type SovrenParserModelResumeCertification struct {

	// matched from list
	MatchedFromList bool `json:"MatchedFromList,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this sovren parser model resume certification
func (m *SovrenParserModelResumeCertification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model resume certification based on context it is used
func (m *SovrenParserModelResumeCertification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeCertification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeCertification) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeCertification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
