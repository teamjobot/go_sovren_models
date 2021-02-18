// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainV10IndexingOptions sovren saas domain v10 indexing options
//
// swagger:model Sovren.Saas.Domain.V10IndexingOptions
type SovrenSaasDomainV10IndexingOptions struct {

	// document Id
	DocumentID string `json:"DocumentId,omitempty"`

	// index Id
	IndexID string `json:"IndexId,omitempty"`

	// user defined tags
	UserDefinedTags []string `json:"UserDefinedTags"`
}

// Validate validates this sovren saas domain v10 indexing options
func (m *SovrenSaasDomainV10IndexingOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren saas domain v10 indexing options based on context it is used
func (m *SovrenSaasDomainV10IndexingOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10IndexingOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10IndexingOptions) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10IndexingOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
