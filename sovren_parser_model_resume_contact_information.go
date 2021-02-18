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

// SovrenParserModelResumeContactInformation sovren parser model resume contact information
//
// swagger:model Sovren.Parser.Model.Resume.ContactInformation
type SovrenParserModelResumeContactInformation struct {

	// candidate name
	CandidateName *SovrenParserModelResumePersonName `json:"CandidateName,omitempty"`

	// email addresses
	EmailAddresses []string `json:"EmailAddresses"`

	// location
	Location *SovrenParserModelSovrenLocation `json:"Location,omitempty"`

	// telephones
	Telephones []*SovrenParserModelResumeTelephone `json:"Telephones"`

	// web addresses
	WebAddresses []*SovrenParserModelResumeWebAddress `json:"WebAddresses"`
}

// Validate validates this sovren parser model resume contact information
func (m *SovrenParserModelResumeContactInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCandidateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelephones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeContactInformation) validateCandidateName(formats strfmt.Registry) error {
	if swag.IsZero(m.CandidateName) { // not required
		return nil
	}

	if m.CandidateName != nil {
		if err := m.CandidateName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CandidateName")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) validateTelephones(formats strfmt.Registry) error {
	if swag.IsZero(m.Telephones) { // not required
		return nil
	}

	for i := 0; i < len(m.Telephones); i++ {
		if swag.IsZero(m.Telephones[i]) { // not required
			continue
		}

		if m.Telephones[i] != nil {
			if err := m.Telephones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Telephones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) validateWebAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.WebAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.WebAddresses); i++ {
		if swag.IsZero(m.WebAddresses[i]) { // not required
			continue
		}

		if m.WebAddresses[i] != nil {
			if err := m.WebAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WebAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sovren parser model resume contact information based on the context it is used
func (m *SovrenParserModelResumeContactInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCandidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTelephones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SovrenParserModelResumeContactInformation) contextValidateCandidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.CandidateName != nil {
		if err := m.CandidateName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CandidateName")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) contextValidateTelephones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Telephones); i++ {

		if m.Telephones[i] != nil {
			if err := m.Telephones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Telephones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SovrenParserModelResumeContactInformation) contextValidateWebAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebAddresses); i++ {

		if m.WebAddresses[i] != nil {
			if err := m.WebAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WebAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SovrenParserModelResumeContactInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeContactInformation) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeContactInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}