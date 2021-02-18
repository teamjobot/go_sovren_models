// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainV10V10IndexDocumentCountResponse sovren saas domain v10 v10 index document count response
//
// swagger:model Sovren.Saas.Domain.V10.V10IndexDocumentCountResponse
type SovrenSaasDomainV10V10IndexDocumentCountResponse struct {

	// count
	Count int64 `json:"Count,omitempty"`
}

// Validate validates this sovren saas domain v10 v10 index document count response
func (m *SovrenSaasDomainV10V10IndexDocumentCountResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren saas domain v10 v10 index document count response based on context it is used
func (m *SovrenSaasDomainV10V10IndexDocumentCountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10V10IndexDocumentCountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10V10IndexDocumentCountResponse) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10V10IndexDocumentCountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}