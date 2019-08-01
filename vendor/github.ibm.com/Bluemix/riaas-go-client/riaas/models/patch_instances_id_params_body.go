// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchInstancesIDParamsBody patch instances Id params body
// swagger:model patchInstancesIdParamsBody
type PatchInstancesIDParamsBody struct {

	// The user-defined name for this network ACL
	// Pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$
	Name string `json:"name,omitempty"`

	// profile
	Profile *PatchInstancesIDParamsBodyProfile `json:"profile,omitempty"`
}

// Validate validates this patch instances Id params body
func (m *PatchInstancesIDParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchInstancesIDParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`); err != nil {
		return err
	}

	return nil
}

func (m *PatchInstancesIDParamsBody) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchInstancesIDParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchInstancesIDParamsBody) UnmarshalBinary(b []byte) error {
	var res PatchInstancesIDParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
