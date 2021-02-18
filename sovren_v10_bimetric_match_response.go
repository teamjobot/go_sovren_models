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

// SovrenV10BimetricMatchResponse sovren v10 bimetric match response
//
// swagger:model Sovren.V10BimetricMatchResponse
type SovrenV10BimetricMatchResponse struct {

	// applied category weights
	AppliedCategoryWeights *SovrenCategoryWeights `json:"AppliedCategoryWeights,omitempty"`

	// matches
	Matches []*SovrenBimetricMatchResult `json:"Matches"`

	// suggested category weights
	SuggestedCategoryWeights *SovrenCategoryWeights `json:"SuggestedCategoryWeights,omitempty"`
}

// Validate validates this sovren v10 bimetric match response
func (m *SovrenV10BimetricMatchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppliedCategoryWeights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuggestedCategoryWeights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenV10BimetricMatchResponse) validateAppliedCategoryWeights(formats strfmt.Registry) error {
	if swag.IsZero(m.AppliedCategoryWeights) { // not required
		return nil
	}

	if m.AppliedCategoryWeights != nil {
		if err := m.AppliedCategoryWeights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AppliedCategoryWeights")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10BimetricMatchResponse) validateMatches(formats strfmt.Registry) error {
	if swag.IsZero(m.Matches) { // not required
		return nil
	}

	for i := 0; i < len(m.Matches); i++ {
		if swag.IsZero(m.Matches[i]) { // not required
			continue
		}

		if m.Matches[i] != nil {
			if err := m.Matches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenV10BimetricMatchResponse) validateSuggestedCategoryWeights(formats strfmt.Registry) error {
	if swag.IsZero(m.SuggestedCategoryWeights) { // not required
		return nil
	}

	if m.SuggestedCategoryWeights != nil {
		if err := m.SuggestedCategoryWeights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SuggestedCategoryWeights")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren v10 bimetric match response based on the context it is used
func (m *SovrenV10BimetricMatchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppliedCategoryWeights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuggestedCategoryWeights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenV10BimetricMatchResponse) contextValidateAppliedCategoryWeights(ctx context.Context, formats strfmt.Registry) error {

	if m.AppliedCategoryWeights != nil {
		if err := m.AppliedCategoryWeights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AppliedCategoryWeights")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10BimetricMatchResponse) contextValidateMatches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Matches); i++ {

		if m.Matches[i] != nil {
			if err := m.Matches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenV10BimetricMatchResponse) contextValidateSuggestedCategoryWeights(ctx context.Context, formats strfmt.Registry) error {

	if m.SuggestedCategoryWeights != nil {
		if err := m.SuggestedCategoryWeights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SuggestedCategoryWeights")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenV10BimetricMatchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenV10BimetricMatchResponse) UnmarshalBinary(b []byte) error {
	var res SovrenV10BimetricMatchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}