// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IKEPolicyCollection i k e policy collection
// swagger:model IKEPolicyCollection
type IKEPolicyCollection struct {

	// first
	First *IKEPolicyCollectionFirst `json:"first,omitempty"`

	// Collection of IKE policies
	IkePolicies []*IKEPolicy `json:"ike_policies"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// next
	Next *IKEPolicyCollectionNext `json:"next,omitempty"`
}

// Validate validates this i k e policy collection
func (m *IKEPolicyCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIkePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicyCollection) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *IKEPolicyCollection) validateIkePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.IkePolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.IkePolicies); i++ {
		if swag.IsZero(m.IkePolicies[i]) { // not required
			continue
		}

		if m.IkePolicies[i] != nil {
			if err := m.IkePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ike_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IKEPolicyCollection) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", int64(m.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyCollection) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IKEPolicyCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IKEPolicyCollection) UnmarshalBinary(b []byte) error {
	var res IKEPolicyCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
