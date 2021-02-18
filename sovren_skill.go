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

// SovrenSkill sovren skill
//
// swagger:model Sovren.Skill
type SovrenSkill struct {

	// experience level
	// Enum: [Unknown Low Mid High]
	ExperienceLevel string `json:"ExperienceLevel,omitempty"`

	// is current
	IsCurrent bool `json:"IsCurrent,omitempty"`

	// months of experience range
	MonthsOfExperienceRange *SovrenIntegerRange `json:"MonthsOfExperienceRange,omitempty"`

	// skill name
	SkillName string `json:"SkillName,omitempty"`
}

// Validate validates this sovren skill
func (m *SovrenSkill) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperienceLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthsOfExperienceRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sovrenSkillTypeExperienceLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Low","Mid","High"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sovrenSkillTypeExperienceLevelPropEnum = append(sovrenSkillTypeExperienceLevelPropEnum, v)
	}
}

const (

	// SovrenSkillExperienceLevelUnknown captures enum value "Unknown"
	SovrenSkillExperienceLevelUnknown string = "Unknown"

	// SovrenSkillExperienceLevelLow captures enum value "Low"
	SovrenSkillExperienceLevelLow string = "Low"

	// SovrenSkillExperienceLevelMid captures enum value "Mid"
	SovrenSkillExperienceLevelMid string = "Mid"

	// SovrenSkillExperienceLevelHigh captures enum value "High"
	SovrenSkillExperienceLevelHigh string = "High"
)

// prop value enum
func (m *SovrenSkill) validateExperienceLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sovrenSkillTypeExperienceLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SovrenSkill) validateExperienceLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.ExperienceLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateExperienceLevelEnum("ExperienceLevel", "body", m.ExperienceLevel); err != nil {
		return err
	}

	return nil
}

func (m *SovrenSkill) validateMonthsOfExperienceRange(formats strfmt.Registry) error {
	if swag.IsZero(m.MonthsOfExperienceRange) { // not required
		return nil
	}

	if m.MonthsOfExperienceRange != nil {
		if err := m.MonthsOfExperienceRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MonthsOfExperienceRange")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren skill based on the context it is used
func (m *SovrenSkill) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonthsOfExperienceRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSkill) contextValidateMonthsOfExperienceRange(ctx context.Context, formats strfmt.Registry) error {

	if m.MonthsOfExperienceRange != nil {
		if err := m.MonthsOfExperienceRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MonthsOfExperienceRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSkill) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSkill) UnmarshalBinary(b []byte) error {
	var res SovrenSkill
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
