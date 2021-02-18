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

// SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel sovren saas domain v10 add multiple jobs to index request model
//
// swagger:model Sovren.Saas.Domain.V10.AddMultipleJobsToIndexRequestModel
type SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel struct {

	// jobs
	Jobs []*SovrenIndexJobRequest `json:"Jobs"`
}

// Validate validates this sovren saas domain v10 add multiple jobs to index request model
func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) validateJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren saas domain v10 add multiple jobs to index request model based on the context it is used
func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) contextValidateJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Jobs); i++ {

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainV10AddMultipleJobsToIndexRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}