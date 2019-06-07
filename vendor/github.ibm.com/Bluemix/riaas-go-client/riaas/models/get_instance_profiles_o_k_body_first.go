// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetInstanceProfilesOKBodyFirst get instance profiles o k body first
// swagger:model getInstanceProfilesOKBodyFirst
type GetInstanceProfilesOKBodyFirst struct {

	// href
	Href interface{} `json:"href,omitempty"`
}

// Validate validates this get instance profiles o k body first
func (m *GetInstanceProfilesOKBodyFirst) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetInstanceProfilesOKBodyFirst) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInstanceProfilesOKBodyFirst) UnmarshalBinary(b []byte) error {
	var res GetInstanceProfilesOKBodyFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
