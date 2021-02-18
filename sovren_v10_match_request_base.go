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

// SovrenV10MatchRequestBase sovren v10 match request base
//
// swagger:model Sovren.V10MatchRequestBase
type SovrenV10MatchRequestBase struct {

	// category weights
	// Read Only: true
	CategoryWeights []*SovrenCategoryWeight `json:"CategoryWeights"`

	// filter criteria
	FilterCriteria *SovrenV10FilterCriteria `json:"FilterCriteria,omitempty"`

	// index ids to search into
	IndexIdsToSearchInto []string `json:"IndexIdsToSearchInto"`

	// preferred category weights
	PreferredCategoryWeights *SovrenCategoryWeights `json:"PreferredCategoryWeights,omitempty"`

	// settings
	Settings *SovrenMatchSettings `json:"Settings,omitempty"`

	// take
	Take int32 `json:"Take,omitempty"`
}

// Validate validates this sovren v10 match request base
func (m *SovrenV10MatchRequestBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryWeights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredCategoryWeights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenV10MatchRequestBase) validateCategoryWeights(formats strfmt.Registry) error {
	if swag.IsZero(m.CategoryWeights) { // not required
		return nil
	}

	for i := 0; i < len(m.CategoryWeights); i++ {
		if swag.IsZero(m.CategoryWeights[i]) { // not required
			continue
		}

		if m.CategoryWeights[i] != nil {
			if err := m.CategoryWeights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CategoryWeights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenV10MatchRequestBase) validateFilterCriteria(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterCriteria) { // not required
		return nil
	}

	if m.FilterCriteria != nil {
		if err := m.FilterCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FilterCriteria")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10MatchRequestBase) validatePreferredCategoryWeights(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredCategoryWeights) { // not required
		return nil
	}

	if m.PreferredCategoryWeights != nil {
		if err := m.PreferredCategoryWeights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PreferredCategoryWeights")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10MatchRequestBase) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren v10 match request base based on the context it is used
func (m *SovrenV10MatchRequestBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategoryWeights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredCategoryWeights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenV10MatchRequestBase) contextValidateCategoryWeights(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "CategoryWeights", "body", []*SovrenCategoryWeight(m.CategoryWeights)); err != nil {
		return err
	}

	for i := 0; i < len(m.CategoryWeights); i++ {

		if m.CategoryWeights[i] != nil {
			if err := m.CategoryWeights[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CategoryWeights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenV10MatchRequestBase) contextValidateFilterCriteria(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterCriteria != nil {
		if err := m.FilterCriteria.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FilterCriteria")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10MatchRequestBase) contextValidatePreferredCategoryWeights(ctx context.Context, formats strfmt.Registry) error {

	if m.PreferredCategoryWeights != nil {
		if err := m.PreferredCategoryWeights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PreferredCategoryWeights")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10MatchRequestBase) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenV10MatchRequestBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenV10MatchRequestBase) UnmarshalBinary(b []byte) error {
	var res SovrenV10MatchRequestBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
