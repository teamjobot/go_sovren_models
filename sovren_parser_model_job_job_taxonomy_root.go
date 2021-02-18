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

// SovrenParserModelJobJobTaxonomyRoot sovren parser model job job taxonomy root
//
// swagger:model Sovren.Parser.Model.Job.JobTaxonomyRoot
type SovrenParserModelJobJobTaxonomyRoot struct {

	// root
	Root string `json:"Root,omitempty"`

	// taxonomies
	Taxonomies []*SovrenParserModelJobJobTaxonomy `json:"Taxonomies"`
}

// Validate validates this sovren parser model job job taxonomy root
func (m *SovrenParserModelJobJobTaxonomyRoot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxonomies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelJobJobTaxonomyRoot) validateTaxonomies(formats strfmt.Registry) error {
	if swag.IsZero(m.Taxonomies) { // not required
		return nil
	}

	for i := 0; i < len(m.Taxonomies); i++ {
		if swag.IsZero(m.Taxonomies[i]) { // not required
			continue
		}

		if m.Taxonomies[i] != nil {
			if err := m.Taxonomies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Taxonomies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model job job taxonomy root based on the context it is used
func (m *SovrenParserModelJobJobTaxonomyRoot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxonomies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelJobJobTaxonomyRoot) contextValidateTaxonomies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Taxonomies); i++ {

		if m.Taxonomies[i] != nil {
			if err := m.Taxonomies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Taxonomies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelJobJobTaxonomyRoot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelJobJobTaxonomyRoot) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelJobJobTaxonomyRoot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}