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

// SovrenIndexDetail sovren index detail
//
// swagger:model Sovren.IndexDetail
type SovrenIndexDetail struct {

	// index type
	// Enum: [Resume Job]
	IndexType string `json:"IndexType,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// owner Id
	OwnerID string `json:"OwnerId,omitempty"`
}

// Validate validates this sovren index detail
func (m *SovrenIndexDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sovrenIndexDetailTypeIndexTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Resume","Job"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sovrenIndexDetailTypeIndexTypePropEnum = append(sovrenIndexDetailTypeIndexTypePropEnum, v)
	}
}

const (

	// SovrenIndexDetailIndexTypeResume captures enum value "Resume"
	SovrenIndexDetailIndexTypeResume string = "Resume"

	// SovrenIndexDetailIndexTypeJob captures enum value "Job"
	SovrenIndexDetailIndexTypeJob string = "Job"
)

// prop value enum
func (m *SovrenIndexDetail) validateIndexTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sovrenIndexDetailTypeIndexTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SovrenIndexDetail) validateIndexType(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndexTypeEnum("IndexType", "body", m.IndexType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sovren index detail based on context it is used
func (m *SovrenIndexDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenIndexDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenIndexDetail) UnmarshalBinary(b []byte) error {
	var res SovrenIndexDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
