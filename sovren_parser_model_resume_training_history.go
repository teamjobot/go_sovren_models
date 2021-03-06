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

// SovrenParserModelResumeTrainingHistory sovren parser model resume training history
//
// swagger:model Sovren.Parser.Model.Resume.TrainingHistory
type SovrenParserModelResumeTrainingHistory struct {

	// text
	Text string `json:"Text,omitempty"`

	// trainings
	Trainings []*SovrenParserModelResumeTrainingDetail `json:"Trainings"`
}

// Validate validates this sovren parser model resume training history
func (m *SovrenParserModelResumeTrainingHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrainings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeTrainingHistory) validateTrainings(formats strfmt.Registry) error {
	if swag.IsZero(m.Trainings) { // not required
		return nil
	}

	for i := 0; i < len(m.Trainings); i++ {
		if swag.IsZero(m.Trainings[i]) { // not required
			continue
		}

		if m.Trainings[i] != nil {
			if err := m.Trainings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Trainings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model resume training history based on the context it is used
func (m *SovrenParserModelResumeTrainingHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrainings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeTrainingHistory) contextValidateTrainings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Trainings); i++ {

		if m.Trainings[i] != nil {
			if err := m.Trainings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Trainings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeTrainingHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeTrainingHistory) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeTrainingHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
