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
	"github.com/go-openapi/validate"
)

// SovrenMatchResult sovren match result
//
// swagger:model Sovren.MatchResult
type SovrenMatchResult struct {

	// enriched r c s score data
	EnrichedRCSScoreData *SovrenCategoryScoresCategoryScoresContainer `json:"EnrichedRCSScoreData,omitempty"`

	// enriched score data
	EnrichedScoreData *SovrenCategoryScoresCategoryScoresContainer `json:"EnrichedScoreData,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// index Id
	IndexID string `json:"IndexId,omitempty"`

	// reverse compatibility score
	ReverseCompatibilityScore int32 `json:"ReverseCompatibilityScore,omitempty"`

	// sov score
	// Read Only: true
	SovScore int32 `json:"SovScore,omitempty"`

	// unweighted category scores
	UnweightedCategoryScores []*SovrenCategoryScore `json:"UnweightedCategoryScores"`

	// weighted score
	WeightedScore int32 `json:"WeightedScore,omitempty"`
}

// Validate validates this sovren match result
func (m *SovrenMatchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrichedRCSScoreData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrichedScoreData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnweightedCategoryScores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenMatchResult) validateEnrichedRCSScoreData(formats strfmt.Registry) error {
	if swag.IsZero(m.EnrichedRCSScoreData) { // not required
		return nil
	}

	if m.EnrichedRCSScoreData != nil {
		if err := m.EnrichedRCSScoreData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EnrichedRCSScoreData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenMatchResult) validateEnrichedScoreData(formats strfmt.Registry) error {
	if swag.IsZero(m.EnrichedScoreData) { // not required
		return nil
	}

	if m.EnrichedScoreData != nil {
		if err := m.EnrichedScoreData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EnrichedScoreData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenMatchResult) validateUnweightedCategoryScores(formats strfmt.Registry) error {
	if swag.IsZero(m.UnweightedCategoryScores) { // not required
		return nil
	}

	for i := 0; i < len(m.UnweightedCategoryScores); i++ {
		if swag.IsZero(m.UnweightedCategoryScores[i]) { // not required
			continue
		}

		if m.UnweightedCategoryScores[i] != nil {
			if err := m.UnweightedCategoryScores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnweightedCategoryScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren match result based on the context it is used
func (m *SovrenMatchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnrichedRCSScoreData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnrichedScoreData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSovScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnweightedCategoryScores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenMatchResult) contextValidateEnrichedRCSScoreData(ctx context.Context, formats strfmt.Registry) error {

	if m.EnrichedRCSScoreData != nil {
		if err := m.EnrichedRCSScoreData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EnrichedRCSScoreData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenMatchResult) contextValidateEnrichedScoreData(ctx context.Context, formats strfmt.Registry) error {

	if m.EnrichedScoreData != nil {
		if err := m.EnrichedScoreData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EnrichedScoreData")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenMatchResult) contextValidateSovScore(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "SovScore", "body", int32(m.SovScore)); err != nil {
		return err
	}

	return nil
}

func (m *SovrenMatchResult) contextValidateUnweightedCategoryScores(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnweightedCategoryScores); i++ {

		if m.UnweightedCategoryScores[i] != nil {
			if err := m.UnweightedCategoryScores[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnweightedCategoryScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenMatchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenMatchResult) UnmarshalBinary(b []byte) error {
	var res SovrenMatchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
