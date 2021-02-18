// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 sovren range 1 system string mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
//
// swagger:model Sovren.Range`1[[System.String, mscorlib, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089]]
type SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 struct {

	// maximum
	Maximum string `json:"Maximum,omitempty"`

	// minimum
	Minimum string `json:"Minimum,omitempty"`
}

// Validate validates this sovren range 1 system string mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
func (m *SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren range 1 system string mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089 based on context it is used
func (m *SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) UnmarshalBinary(b []byte) error {
	var res SovrenRange1SystemStringMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}