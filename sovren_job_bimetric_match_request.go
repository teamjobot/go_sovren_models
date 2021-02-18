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

// SovrenJobBimetricMatchRequest sovren job bimetric match request
//
// swagger:model Sovren.JobBimetricMatchRequest
type SovrenJobBimetricMatchRequest struct {

	// preferred category weights
	PreferredCategoryWeights *SovrenCategoryWeights `json:"PreferredCategoryWeights,omitempty"`

	// settings
	Settings *SovrenMatchSettings `json:"Settings,omitempty"`

	// source job
	SourceJob *SovrenV10ParsedJob `json:"SourceJob,omitempty"`

	// target jobs
	TargetJobs []*SovrenV10ParsedJob `json:"TargetJobs"`

	// target resumes
	TargetResumes []*SovrenV10ParsedResume `json:"TargetResumes"`
}

// Validate validates this sovren job bimetric match request
func (m *SovrenJobBimetricMatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredCategoryWeights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetResumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenJobBimetricMatchRequest) validatePreferredCategoryWeights(formats strfmt.Registry) error {
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

func (m *SovrenJobBimetricMatchRequest) validateSettings(formats strfmt.Registry) error {
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

func (m *SovrenJobBimetricMatchRequest) validateSourceJob(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceJob) { // not required
		return nil
	}

	if m.SourceJob != nil {
		if err := m.SourceJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SourceJob")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenJobBimetricMatchRequest) validateTargetJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetJobs) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetJobs); i++ {
		if swag.IsZero(m.TargetJobs[i]) { // not required
			continue
		}

		if m.TargetJobs[i] != nil {
			if err := m.TargetJobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TargetJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenJobBimetricMatchRequest) validateTargetResumes(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetResumes) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetResumes); i++ {
		if swag.IsZero(m.TargetResumes[i]) { // not required
			continue
		}

		if m.TargetResumes[i] != nil {
			if err := m.TargetResumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TargetResumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren job bimetric match request based on the context it is used
func (m *SovrenJobBimetricMatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreferredCategoryWeights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetResumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenJobBimetricMatchRequest) contextValidatePreferredCategoryWeights(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenJobBimetricMatchRequest) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenJobBimetricMatchRequest) contextValidateSourceJob(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceJob != nil {
		if err := m.SourceJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SourceJob")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenJobBimetricMatchRequest) contextValidateTargetJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetJobs); i++ {

		if m.TargetJobs[i] != nil {
			if err := m.TargetJobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TargetJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenJobBimetricMatchRequest) contextValidateTargetResumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetResumes); i++ {

		if m.TargetResumes[i] != nil {
			if err := m.TargetResumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TargetResumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenJobBimetricMatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenJobBimetricMatchRequest) UnmarshalBinary(b []byte) error {
	var res SovrenJobBimetricMatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}