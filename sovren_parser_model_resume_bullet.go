// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenParserModelResumeBullet sovren parser model resume bullet
//
// swagger:model Sovren.Parser.Model.Resume.Bullet
type SovrenParserModelResumeBullet struct {

	// text
	Text string `json:"Text,omitempty"`

	// type
	Type string `json:"Type,omitempty"`
}

// Validate validates this sovren parser model resume bullet
func (m *SovrenParserModelResumeBullet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren parser model resume bullet based on context it is used
func (m *SovrenParserModelResumeBullet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeBullet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeBullet) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeBullet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
