// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainV10ConversionMetadata sovren saas domain v10 conversion metadata
//
// swagger:model Sovren.Saas.Domain.V10.ConversionMetadata
type SovrenSaasDomainV10ConversionMetadata struct {

	// detected type
	DetectedType string `json:"DetectedType,omitempty"`

	// elapsed milliseconds
	ElapsedMilliseconds int32 `json:"ElapsedMilliseconds,omitempty"`

	// output validity code
	OutputValidityCode string `json:"OutputValidityCode,omitempty"`

	// suggested file extension
	SuggestedFileExtension string `json:"SuggestedFileExtension,omitempty"`
}

// Validate validates this sovren saas domain v10 conversion metadata
func (m *SovrenSaasDomainV10ConversionMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren saas domain v10 conversion metadata based on context it is used
func (m *SovrenSaasDomainV10ConversionMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10ConversionMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10ConversionMetadata) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10ConversionMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
