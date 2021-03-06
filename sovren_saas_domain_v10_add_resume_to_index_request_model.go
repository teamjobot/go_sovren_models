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

// SovrenSaasDomainV10AddResumeToIndexRequestModel sovren saas domain v10 add resume to index request model
//
// swagger:model Sovren.Saas.Domain.V10.AddResumeToIndexRequestModel
type SovrenSaasDomainV10AddResumeToIndexRequestModel struct {

	// resume data
	ResumeData *SovrenParserModelResumeResumeData `json:"ResumeData,omitempty"`

	// user defined tags
	UserDefinedTags []string `json:"UserDefinedTags"`
}

// Validate validates this sovren saas domain v10 add resume to index request model
func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResumeData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) validateResumeData(formats strfmt.Registry) error {
	if swag.IsZero(m.ResumeData) { // not required
		return nil
	}

	if m.ResumeData != nil {
		if err := m.ResumeData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResumeData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 add resume to index request model based on the context it is used
func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResumeData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) contextValidateResumeData(ctx context.Context, formats strfmt.Registry) error {

	if m.ResumeData != nil {
		if err := m.ResumeData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResumeData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10AddResumeToIndexRequestModel) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10AddResumeToIndexRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
