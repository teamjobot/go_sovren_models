// Code generated by go-swagger; DO NOT EDIT.

package go_sovren_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SovrenV10SearchRequest sovren v10 search request
//
// swagger:model Sovren.V10SearchRequest
type SovrenV10SearchRequest struct {

	// filter criteria
	FilterCriteria *SovrenV10FilterCriteria `json:"FilterCriteria,omitempty"`

	// index ids to search into
	IndexIdsToSearchInto []string `json:"IndexIdsToSearchInto"`

	// pagination settings
	PaginationSettings *SovrenPaginationSettings `json:"PaginationSettings,omitempty"`

	// settings
	Settings *SovrenMatchSettings `json:"Settings,omitempty"`
}

// Validate validates this sovren v10 search request
func (m *SovrenV10SearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaginationSettings(formats); err != nil {
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

func (m *SovrenV10SearchRequest) validateFilterCriteria(formats strfmt.Registry) error {
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

func (m *SovrenV10SearchRequest) validatePaginationSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.PaginationSettings) { // not required
		return nil
	}

	if m.PaginationSettings != nil {
		if err := m.PaginationSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PaginationSettings")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10SearchRequest) validateSettings(formats strfmt.Registry) error {
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

// ContextValidate validate this sovren v10 search request based on the context it is used
func (m *SovrenV10SearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaginationSettings(ctx, formats); err != nil {
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

func (m *SovrenV10SearchRequest) contextValidateFilterCriteria(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenV10SearchRequest) contextValidatePaginationSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.PaginationSettings != nil {
		if err := m.PaginationSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PaginationSettings")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenV10SearchRequest) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SovrenV10SearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenV10SearchRequest) UnmarshalBinary(b []byte) error {
	var res SovrenV10SearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
