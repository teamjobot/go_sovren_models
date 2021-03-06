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

// SovrenParserModelResumeMilitaryDetails sovren parser model resume military details
//
// swagger:model Sovren.Parser.Model.Resume.MilitaryDetails
type SovrenParserModelResumeMilitaryDetails struct {

	// country
	Country string `json:"Country,omitempty"`

	// end date
	EndDate *SovrenParserModelSovrenDateWithParts `json:"EndDate,omitempty"`

	// found in context
	FoundInContext string `json:"FoundInContext,omitempty"`

	// service
	Service *SovrenParserModelResumeMilitaryService `json:"Service,omitempty"`

	// start date
	StartDate *SovrenParserModelSovrenDateWithParts `json:"StartDate,omitempty"`
}

// Validate validates this sovren parser model resume military details
func (m *SovrenParserModelResumeMilitaryDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if m.EndDate != nil {
		if err := m.EndDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndDate")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Service")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if m.StartDate != nil {
		if err := m.StartDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartDate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren parser model resume military details based on the context it is used
func (m *SovrenParserModelResumeMilitaryDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if m.EndDate != nil {
		if err := m.EndDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndDate")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {
		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Service")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeMilitaryDetails) contextValidateStartDate(ctx context.Context, formats strfmt.Registry) error {

	if m.StartDate != nil {
		if err := m.StartDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartDate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeMilitaryDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeMilitaryDetails) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeMilitaryDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
