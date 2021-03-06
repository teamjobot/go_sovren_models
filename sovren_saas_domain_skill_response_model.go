// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainSkillResponseModel sovren saas domain skill response model
//
// swagger:model Sovren.Saas.Domain.SkillResponseModel
type SovrenSaasDomainSkillResponseModel struct {

	// key
	Key string `json:"Key,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// timestamp
	Timestamp string `json:"Timestamp,omitempty"`
}

// Validate validates this sovren saas domain skill response model
func (m *SovrenSaasDomainSkillResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren saas domain skill response model based on context it is used
func (m *SovrenSaasDomainSkillResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainSkillResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainSkillResponseModel) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainSkillResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
