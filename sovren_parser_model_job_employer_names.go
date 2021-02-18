// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelJobEmployerNames sovren parser model job employer names
//
// swagger:model Sovren.Parser.Model.Job.EmployerNames
type SovrenParserModelJobEmployerNames struct {

	// employer name
	EmployerName []string `json:"EmployerName"`

	// main employer name
	MainEmployerName string `json:"MainEmployerName,omitempty"`
}

// Validate validates this sovren parser model job employer names
func (m *SovrenParserModelJobEmployerNames) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model job employer names based on context it is used
func (m *SovrenParserModelJobEmployerNames) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelJobEmployerNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelJobEmployerNames) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelJobEmployerNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
