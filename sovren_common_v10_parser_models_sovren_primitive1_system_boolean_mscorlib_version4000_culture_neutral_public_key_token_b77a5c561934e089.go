// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 sovren common v10 parser models sovren primitive 1 system boolean mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
//
// swagger:model Sovren.Common.V10.ParserModels.SovrenPrimitive`1[[System.Boolean, mscorlib, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089]]
type SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 struct {

	// value
	Value bool `json:"Value,omitempty"`
}

// Validate validates this sovren common v10 parser models sovren primitive 1 system boolean mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089
func (m *SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren common v10 parser models sovren primitive 1 system boolean mscorlib version 4 0 0 0 culture neutral public key token b77a5c561934e089 based on context it is used
func (m *SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089) UnmarshalBinary(b []byte) error {
	var res SovrenCommonV10ParserModelsSovrenPrimitive1SystemBooleanMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
