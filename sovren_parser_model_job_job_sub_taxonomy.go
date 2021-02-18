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

// SovrenParserModelJobJobSubTaxonomy sovren parser model job job sub taxonomy
//
// swagger:model Sovren.Parser.Model.Job.JobSubTaxonomy
type SovrenParserModelJobJobSubTaxonomy struct {

	// percent of overall
	PercentOfOverall int32 `json:"PercentOfOverall,omitempty"`

	// percent of parent
	PercentOfParent int32 `json:"PercentOfParent,omitempty"`

	// skills
	Skills []*SovrenParserModelJobJobSkill `json:"Skills"`

	// sub taxonomy Id
	SubTaxonomyID string `json:"SubTaxonomyId,omitempty"`

	// sub taxonomy name
	SubTaxonomyName string `json:"SubTaxonomyName,omitempty"`
}

// Validate validates this sovren parser model job job sub taxonomy
func (m *SovrenParserModelJobJobSubTaxonomy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelJobJobSubTaxonomy) validateSkills(formats strfmt.Registry) error {
	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model job job sub taxonomy based on the context it is used
func (m *SovrenParserModelJobJobSubTaxonomy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSkills(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelJobJobSubTaxonomy) contextValidateSkills(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Skills); i++ {

		if m.Skills[i] != nil {
			if err := m.Skills[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelJobJobSubTaxonomy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelJobJobSubTaxonomy) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelJobJobSubTaxonomy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
