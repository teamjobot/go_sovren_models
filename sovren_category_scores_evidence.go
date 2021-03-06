// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SovrenCategoryScoresEvidence sovren category scores evidence
//
// swagger:model Sovren.CategoryScores.Evidence
type SovrenCategoryScoresEvidence struct {

	// fact
	Fact string `json:"Fact,omitempty"`

	// type
	// Enum: [Negative Mixed Positive]
	Type string `json:"Type,omitempty"`
}

// Validate validates this sovren category scores evidence
func (m *SovrenCategoryScoresEvidence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sovrenCategoryScoresEvidenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Negative","Mixed","Positive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sovrenCategoryScoresEvidenceTypeTypePropEnum = append(sovrenCategoryScoresEvidenceTypeTypePropEnum, v)
	}
}

const (

	// SovrenCategoryScoresEvidenceTypeNegative captures enum value "Negative"
	SovrenCategoryScoresEvidenceTypeNegative string = "Negative"

	// SovrenCategoryScoresEvidenceTypeMixed captures enum value "Mixed"
	SovrenCategoryScoresEvidenceTypeMixed string = "Mixed"

	// SovrenCategoryScoresEvidenceTypePositive captures enum value "Positive"
	SovrenCategoryScoresEvidenceTypePositive string = "Positive"
)

// prop value enum
func (m *SovrenCategoryScoresEvidence) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sovrenCategoryScoresEvidenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SovrenCategoryScoresEvidence) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sovren category scores evidence based on context it is used
func (m *SovrenCategoryScoresEvidence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenCategoryScoresEvidence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCategoryScoresEvidence) UnmarshalBinary(b []byte) error {
	var res SovrenCategoryScoresEvidence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
