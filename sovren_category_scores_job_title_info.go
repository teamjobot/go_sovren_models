// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenCategoryScoresJobTitleInfo sovren category scores job title info
//
// swagger:model Sovren.CategoryScores.JobTitleInfo
type SovrenCategoryScoresJobTitleInfo struct {

	// is current
	IsCurrent bool `json:"IsCurrent,omitempty"`

	// raw term
	RawTerm string `json:"RawTerm,omitempty"`

	// variation of
	VariationOf string `json:"VariationOf,omitempty"`
}

// Validate validates this sovren category scores job title info
func (m *SovrenCategoryScoresJobTitleInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sovren category scores job title info based on context it is used
func (m *SovrenCategoryScoresJobTitleInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenCategoryScoresJobTitleInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCategoryScoresJobTitleInfo) UnmarshalBinary(b []byte) error {
	var res SovrenCategoryScoresJobTitleInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
