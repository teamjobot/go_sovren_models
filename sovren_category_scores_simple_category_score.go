// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenCategoryScoresSimpleCategoryScore sovren category scores simple category score
//
// swagger:model Sovren.CategoryScores.SimpleCategoryScore
type SovrenCategoryScoresSimpleCategoryScore struct {

	// evidence
	Evidence []*SovrenCategoryScoresEvidence `json:"Evidence"`

	// found
	Found []string `json:"Found"`

	// not found
	NotFound []string `json:"NotFound"`

	// unweighted score
	UnweightedScore float64 `json:"UnweightedScore,omitempty"`
}

// Validate validates this sovren category scores simple category score
func (m *SovrenCategoryScoresSimpleCategoryScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvidence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenCategoryScoresSimpleCategoryScore) validateEvidence(formats strfmt.Registry) error {
	if swag.IsZero(m.Evidence) { // not required
		return nil
	}

	for i := 0; i < len(m.Evidence); i++ {
		if swag.IsZero(m.Evidence[i]) { // not required
			continue
		}

		if m.Evidence[i] != nil {
			if err := m.Evidence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Evidence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren category scores simple category score based on the context it is used
func (m *SovrenCategoryScoresSimpleCategoryScore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvidence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenCategoryScoresSimpleCategoryScore) contextValidateEvidence(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Evidence); i++ {

		if m.Evidence[i] != nil {
			if err := m.Evidence[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Evidence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenCategoryScoresSimpleCategoryScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCategoryScoresSimpleCategoryScore) UnmarshalBinary(b []byte) error {
	var res SovrenCategoryScoresSimpleCategoryScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}