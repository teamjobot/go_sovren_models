// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenSaasDomainCreateSkillRequest sovren saas domain create skill request
//
// swagger:model Sovren.Saas.Domain.CreateSkillRequest
type SovrenSaasDomainCreateSkillRequest struct {

	// content
	Content []string `json:"Content"`

	// culture
	Culture string `json:"Culture,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this sovren saas domain create skill request
func (m *SovrenSaasDomainCreateSkillRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren saas domain create skill request based on context it is used
func (m *SovrenSaasDomainCreateSkillRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainCreateSkillRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainCreateSkillRequest) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainCreateSkillRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}