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

// SovrenSaasDomainResponseInfoModel sovren saas domain response info model
//
// swagger:model Sovren.Saas.Domain.ResponseInfoModel
type SovrenSaasDomainResponseInfoModel struct {

	// code
	// Enum: [Success SomeErrors MissingParameter InvalidParameter InsufficientData AuthenticationError Unauthorized DataNotFound CoordinatesNotFound UnhandledException ConversionException PossibleTruncationFromTimeout WarningsFoundDuringParsing ConstraintError DuplicateAsset UnsupportedContentTypeHeader InstallationError TooManyRequests MaintenanceMode Conflict]
	Code string `json:"Code,omitempty"`

	// message
	Message string `json:"Message,omitempty"`
}

// Validate validates this sovren saas domain response info model
func (m *SovrenSaasDomainResponseInfoModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sovrenSaasDomainResponseInfoModelTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Success","SomeErrors","MissingParameter","InvalidParameter","InsufficientData","AuthenticationError","Unauthorized","DataNotFound","CoordinatesNotFound","UnhandledException","ConversionException","PossibleTruncationFromTimeout","WarningsFoundDuringParsing","ConstraintError","DuplicateAsset","UnsupportedContentTypeHeader","InstallationError","TooManyRequests","MaintenanceMode","Conflict"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sovrenSaasDomainResponseInfoModelTypeCodePropEnum = append(sovrenSaasDomainResponseInfoModelTypeCodePropEnum, v)
	}
}

const (

	// SovrenSaasDomainResponseInfoModelCodeSuccess captures enum value "Success"
	SovrenSaasDomainResponseInfoModelCodeSuccess string = "Success"

	// SovrenSaasDomainResponseInfoModelCodeSomeErrors captures enum value "SomeErrors"
	SovrenSaasDomainResponseInfoModelCodeSomeErrors string = "SomeErrors"

	// SovrenSaasDomainResponseInfoModelCodeMissingParameter captures enum value "MissingParameter"
	SovrenSaasDomainResponseInfoModelCodeMissingParameter string = "MissingParameter"

	// SovrenSaasDomainResponseInfoModelCodeInvalidParameter captures enum value "InvalidParameter"
	SovrenSaasDomainResponseInfoModelCodeInvalidParameter string = "InvalidParameter"

	// SovrenSaasDomainResponseInfoModelCodeInsufficientData captures enum value "InsufficientData"
	SovrenSaasDomainResponseInfoModelCodeInsufficientData string = "InsufficientData"

	// SovrenSaasDomainResponseInfoModelCodeAuthenticationError captures enum value "AuthenticationError"
	SovrenSaasDomainResponseInfoModelCodeAuthenticationError string = "AuthenticationError"

	// SovrenSaasDomainResponseInfoModelCodeUnauthorized captures enum value "Unauthorized"
	SovrenSaasDomainResponseInfoModelCodeUnauthorized string = "Unauthorized"

	// SovrenSaasDomainResponseInfoModelCodeDataNotFound captures enum value "DataNotFound"
	SovrenSaasDomainResponseInfoModelCodeDataNotFound string = "DataNotFound"

	// SovrenSaasDomainResponseInfoModelCodeCoordinatesNotFound captures enum value "CoordinatesNotFound"
	SovrenSaasDomainResponseInfoModelCodeCoordinatesNotFound string = "CoordinatesNotFound"

	// SovrenSaasDomainResponseInfoModelCodeUnhandledException captures enum value "UnhandledException"
	SovrenSaasDomainResponseInfoModelCodeUnhandledException string = "UnhandledException"

	// SovrenSaasDomainResponseInfoModelCodeConversionException captures enum value "ConversionException"
	SovrenSaasDomainResponseInfoModelCodeConversionException string = "ConversionException"

	// SovrenSaasDomainResponseInfoModelCodePossibleTruncationFromTimeout captures enum value "PossibleTruncationFromTimeout"
	SovrenSaasDomainResponseInfoModelCodePossibleTruncationFromTimeout string = "PossibleTruncationFromTimeout"

	// SovrenSaasDomainResponseInfoModelCodeWarningsFoundDuringParsing captures enum value "WarningsFoundDuringParsing"
	SovrenSaasDomainResponseInfoModelCodeWarningsFoundDuringParsing string = "WarningsFoundDuringParsing"

	// SovrenSaasDomainResponseInfoModelCodeConstraintError captures enum value "ConstraintError"
	SovrenSaasDomainResponseInfoModelCodeConstraintError string = "ConstraintError"

	// SovrenSaasDomainResponseInfoModelCodeDuplicateAsset captures enum value "DuplicateAsset"
	SovrenSaasDomainResponseInfoModelCodeDuplicateAsset string = "DuplicateAsset"

	// SovrenSaasDomainResponseInfoModelCodeUnsupportedContentTypeHeader captures enum value "UnsupportedContentTypeHeader"
	SovrenSaasDomainResponseInfoModelCodeUnsupportedContentTypeHeader string = "UnsupportedContentTypeHeader"

	// SovrenSaasDomainResponseInfoModelCodeInstallationError captures enum value "InstallationError"
	SovrenSaasDomainResponseInfoModelCodeInstallationError string = "InstallationError"

	// SovrenSaasDomainResponseInfoModelCodeTooManyRequests captures enum value "TooManyRequests"
	SovrenSaasDomainResponseInfoModelCodeTooManyRequests string = "TooManyRequests"

	// SovrenSaasDomainResponseInfoModelCodeMaintenanceMode captures enum value "MaintenanceMode"
	SovrenSaasDomainResponseInfoModelCodeMaintenanceMode string = "MaintenanceMode"

	// SovrenSaasDomainResponseInfoModelCodeConflict captures enum value "Conflict"
	SovrenSaasDomainResponseInfoModelCodeConflict string = "Conflict"
)

// prop value enum
func (m *SovrenSaasDomainResponseInfoModel) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sovrenSaasDomainResponseInfoModelTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SovrenSaasDomainResponseInfoModel) validateCode(formats strfmt.Registry) error {
	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("Code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sovren saas domain response info model based on context it is used
func (m *SovrenSaasDomainResponseInfoModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SovrenSaasDomainResponseInfoModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SovrenSaasDomainResponseInfoModel) UnmarshalBinary(b []byte) error {
	var res SovrenSaasDomainResponseInfoModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}