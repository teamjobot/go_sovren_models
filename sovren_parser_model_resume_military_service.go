// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelResumeMilitaryService sovren parser model resume military service
//
// swagger:model Sovren.Parser.Model.Resume.MilitaryService
type SovrenParserModelResumeMilitaryService struct {

	// branch
	Branch string `json:"Branch,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// rank
	Rank string `json:"Rank,omitempty"`
}

// Validate validates this sovren parser model resume military service
func (m *SovrenParserModelResumeMilitaryService) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model resume military service based on context it is used
func (m *SovrenParserModelResumeMilitaryService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeMilitaryService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeMilitaryService) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeMilitaryService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}