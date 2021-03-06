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

// SovrenParserModelResumeCandidateReference sovren parser model resume candidate reference
//
// swagger:model Sovren.Parser.Model.Resume.CandidateReference
type SovrenParserModelResumeCandidateReference struct {

	// company
	Company string `json:"Company,omitempty"`

	// email addresses
	EmailAddresses []string `json:"EmailAddresses"`

	// location
	Location *SovrenParserModelSovrenLocation `json:"Location,omitempty"`

	// reference name
	ReferenceName *SovrenParserModelResumePersonName `json:"ReferenceName,omitempty"`

	// telephones
	Telephones []*SovrenParserModelNameValue `json:"Telephones"`

	// title
	Title string `json:"Title,omitempty"`

	// type
	Type string `json:"Type,omitempty"`

	// web addresses
	WebAddresses []*SovrenParserModelResumeWebAddress `json:"WebAddresses"`
}

// Validate validates this sovren parser model resume candidate reference
func (m *SovrenParserModelResumeCandidateReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceName(formats); err != nil {
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

func (m *SovrenParserModelResumeCandidateReference) validateLocation(formats strfmt.Registry) error {
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

func (m *SovrenParserModelResumeCandidateReference) validateReferenceName(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceName) { // not required
		return nil
	}

	if m.ReferenceName != nil {
		if err := m.ReferenceName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReferenceName")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeCandidateReference) validateTelephones(formats strfmt.Registry) error {
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

func (m *SovrenParserModelResumeCandidateReference) validateWebAddresses(formats strfmt.Registry) error {
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

// ContextValidate validate this sovren parser model resume candidate reference based on the context it is used
func (m *SovrenParserModelResumeCandidateReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferenceName(ctx, formats); err != nil {
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

func (m *SovrenParserModelResumeCandidateReference) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenParserModelResumeCandidateReference) contextValidateReferenceName(ctx context.Context, formats strfmt.Registry) error {

	if m.ReferenceName != nil {
		if err := m.ReferenceName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReferenceName")
			}
			return err
		}
	}

	return nil
}

func (m *SovrenParserModelResumeCandidateReference) contextValidateTelephones(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SovrenParserModelResumeCandidateReference) contextValidateWebAddresses(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SovrenParserModelResumeCandidateReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenParserModelResumeCandidateReference) UnmarshalBinary(b []byte) error {
	var res SovrenParserModelResumeCandidateReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
