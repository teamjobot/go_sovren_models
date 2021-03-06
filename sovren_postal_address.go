// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenPostalAddress sovren postal address
//
// swagger:model Sovren.PostalAddress
type SovrenPostalAddress struct {

	// address line
	AddressLine string `json:"AddressLine,omitempty"`

	// country code
	CountryCode string `json:"CountryCode,omitempty"`

	// municipality
	Municipality string `json:"Municipality,omitempty"`

	// postal code
	PostalCode string `json:"PostalCode,omitempty"`

	// region
	Region string `json:"Region,omitempty"`
}

// Validate validates this sovren postal address
func (m *SovrenPostalAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren postal address based on context it is used
func (m *SovrenPostalAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenPostalAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenPostalAddress) UnmarshalBinary(b []byte) error {
	var res SovrenPostalAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
