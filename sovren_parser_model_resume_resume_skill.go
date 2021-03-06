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

// SovrenParserModelResumeResumeSkill sovren parser model resume resume skill
//
// swagger:model Sovren.Parser.Model.Resume.ResumeSkill
type SovrenParserModelResumeResumeSkill struct {

	// children last used
	ChildrenLastUsed *SovrenCommonV10ParserModelsSovrenPrimitive1SystemDateTimeMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"ChildrenLastUsed,omitempty"`

	// children months experience
	ChildrenMonthsExperience *SovrenCommonV10ParserModelsSovrenPrimitive1SystemInt32MscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"ChildrenMonthsExperience,omitempty"`

	// exists in text
	ExistsInText bool `json:"ExistsInText,omitempty"`

	// found in
	FoundIn []*SovrenParserModelSectionIdentifier `json:"FoundIn"`

	// Id
	ID string `json:"Id,omitempty"`

	// last used
	LastUsed *SovrenCommonV10ParserModelsSovrenPrimitive1SystemDateTimeMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"LastUsed,omitempty"`

	// months experience
	MonthsExperience *SovrenCommonV10ParserModelsSovrenPrimitive1SystemInt32MscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 `json:"MonthsExperience,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// variations
	Variations []*SovrenParserModelResumeResumeSkillVariation `json:"Variations"`
}

// Validate validates this sovren parser model resume resume skill
func (m *SovrenParserModelResumeResumeSkill) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildrenLastUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildrenMonthsExperience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFoundIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthsExperience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateChildrenLastUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildrenLastUsed) { // not required
		return nil
	}

	if m.ChildrenLastUsed != nil {
		if err := m.ChildrenLastUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChildrenLastUsed")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateChildrenMonthsExperience(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildrenMonthsExperience) { // not required
		return nil
	}

	if m.ChildrenMonthsExperience != nil {
		if err := m.ChildrenMonthsExperience.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChildrenMonthsExperience")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateFoundIn(formats strfmt.Registry) error {
	if swag.IsZero(m.FoundIn) { // not required
		return nil
	}

	for i := 0; i < len(m.FoundIn); i++ {
		if swag.IsZero(m.FoundIn[i]) { // not required
			continue
		}

		if m.FoundIn[i] != nil {
			if err := m.FoundIn[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FoundIn" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateLastUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUsed) { // not required
		return nil
	}

	if m.LastUsed != nil {
		if err := m.LastUsed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastUsed")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateMonthsExperience(formats strfmt.Registry) error {
	if swag.IsZero(m.MonthsExperience) { // not required
		return nil
	}

	if m.MonthsExperience != nil {
		if err := m.MonthsExperience.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MonthsExperience")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) validateVariations(formats strfmt.Registry) error {
	if swag.IsZero(m.Variations) { // not required
		return nil
	}

	for i := 0; i < len(m.Variations); i++ {
		if swag.IsZero(m.Variations[i]) { // not required
			continue
		}

		if m.Variations[i] != nil {
			if err := m.Variations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model resume resume skill based on the context it is used
func (m *SovrenParserModelResumeResumeSkill) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildrenLastUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChildrenMonthsExperience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFoundIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonthsExperience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateChildrenLastUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.ChildrenLastUsed != nil {
		if err := m.ChildrenLastUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChildrenLastUsed")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateChildrenMonthsExperience(ctx context.Context, formats strfmt.Registry) error {

	if m.ChildrenMonthsExperience != nil {
		if err := m.ChildrenMonthsExperience.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChildrenMonthsExperience")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateFoundIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FoundIn); i++ {

		if m.FoundIn[i] != nil {
			if err := m.FoundIn[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FoundIn" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateLastUsed(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUsed != nil {
		if err := m.LastUsed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastUsed")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateMonthsExperience(ctx context.Context, formats strfmt.Registry) error {

	if m.MonthsExperience != nil {
		if err := m.MonthsExperience.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MonthsExperience")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeResumeSkill) contextValidateVariations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variations); i++ {

		if m.Variations[i] != nil {
			if err := m.Variations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeSkill) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeResumeSkill) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeResumeSkill
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
