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

// SovrenCategoryScoresTaxonomiesCategoryScore sovren category scores taxonomies category score
//
// swagger:model Sovren.CategoryScores.TaxonomiesCategoryScore
type SovrenCategoryScoresTaxonomiesCategoryScore struct {

	// actual taxonomies
	ActualTaxonomies *SovrenCategoryScoresTaxonomyEvidence `json:"ActualTaxonomies,omitempty"`

	// desired taxonomies
	DesiredTaxonomies *SovrenCategoryScoresTaxonomyEvidence `json:"DesiredTaxonomies,omitempty"`

	// evidence
	Evidence []*SovrenCategoryScoresEvidence `json:"Evidence"`

	// unweighted score
	UnweightedScore float64 `json:"UnweightedScore,omitempty"`
}

// Validate validates this sovren category scores taxonomies category score
func (m *SovrenCategoryScoresTaxonomiesCategoryScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualTaxonomies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredTaxonomies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvidence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) validateActualTaxonomies(formats strfmt.Registry) error {
	if swag.IsZero(m.ActualTaxonomies) { // not required
		return nil
	}

	if m.ActualTaxonomies != nil {
		if err := m.ActualTaxonomies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActualTaxonomies")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) validateDesiredTaxonomies(formats strfmt.Registry) error {
	if swag.IsZero(m.DesiredTaxonomies) { // not required
		return nil
	}

	if m.DesiredTaxonomies != nil {
		if err := m.DesiredTaxonomies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DesiredTaxonomies")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) validateEvidence(formats strfmt.Registry) error {
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

// ContextValidate validate this sovren category scores taxonomies category score based on the context it is used
func (m *SovrenCategoryScoresTaxonomiesCategoryScore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActualTaxonomies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDesiredTaxonomies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvidence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) contextValidateActualTaxonomies(ctx context.Context, formats strfmt.Registry) error {

	if m.ActualTaxonomies != nil {
		if err := m.ActualTaxonomies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActualTaxonomies")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) contextValidateDesiredTaxonomies(ctx context.Context, formats strfmt.Registry) error {

	if m.DesiredTaxonomies != nil {
		if err := m.DesiredTaxonomies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DesiredTaxonomies")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenCategoryScoresTaxonomiesCategoryScore) contextValidateEvidence(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SovrenCategoryScoresTaxonomiesCategoryScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenCategoryScoresTaxonomiesCategoryScore) UnmarshalBinary(b []byte) error {
	var res SovrenCategoryScoresTaxonomiesCategoryScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
