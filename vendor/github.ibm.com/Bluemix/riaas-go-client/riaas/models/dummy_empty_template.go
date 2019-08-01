// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DummyEmptyTemplate dummy empty template
// swagger:model DummyEmptyTemplate
type DummyEmptyTemplate struct {

	// Dummy property
	DumpProps string `json:"dumpProps,omitempty"`
}

// Validate validates this dummy empty template
func (m *DummyEmptyTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DummyEmptyTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DummyEmptyTemplate) UnmarshalBinary(b []byte) error {
	var res DummyEmptyTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
