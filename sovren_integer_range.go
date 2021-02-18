// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenIntegerRange sovren integer range
//
// swagger:model Sovren.IntegerRange
type SovrenIntegerRange struct {

	// maximum
	Maximum int32 `json:"Maximum,omitempty"`

	// minimum
	Minimum int32 `json:"Minimum,omitempty"`
}

// Validate validates this sovren integer range
func (m *SovrenIntegerRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren integer range based on context it is used
func (m *SovrenIntegerRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenIntegerRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenIntegerRange) UnmarshalBinary(b []byte) error {
	var res SovrenIntegerRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}