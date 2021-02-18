// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenCategoryWeight sovren category weight
//
// swagger:model Sovren.CategoryWeight
type SovrenCategoryWeight struct {

	// category
	Category string `json:"Category,omitempty"`

	// weight
	Weight float64 `json:"Weight,omitempty"`
}

// Validate validates this sovren category weight
func (m *SovrenCategoryWeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren category weight based on context it is used
func (m *SovrenCategoryWeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenCategoryWeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCategoryWeight) UnmarshalBinary(b []byte) error {
	var res SovrenCategoryWeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
