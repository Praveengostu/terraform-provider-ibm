// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Image image
// swagger:model Image
type Image struct {

	// Creation Date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// Description
	Description string `json:"description,omitempty"`

	// Image ID
	// Required: true
	ImageID *string `json:"imageID"`

	// Last Update Date
	// Required: true
	// Format: date-time
	LastUpdateDate *strfmt.DateTime `json:"lastUpdateDate"`

	// Image Name
	// Required: true
	Name *string `json:"name"`

	// List of Servers that have deployed the image
	Servers []string `json:"servers"`

	// Image Size
	// Required: true
	Size *float64 `json:"size"`

	// specifications
	Specifications *ImageSpecifications `json:"specifications,omitempty"`

	// Image State
	State string `json:"state,omitempty"`

	// taskref
	Taskref *TaskReference `json:"taskref,omitempty"`

	// Image Volumes
	Volumes []*ImageVolume `json:"volumes"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Image) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdateDate", "body", "date-time", m.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateSpecifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Specifications) { // not required
		return nil
	}

	if m.Specifications != nil {
		if err := m.Specifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specifications")
			}
			return err
		}
	}

	return nil
}

func (m *Image) validateTaskref(formats strfmt.Registry) error {

	if swag.IsZero(m.Taskref) { // not required
		return nil
	}

	if m.Taskref != nil {
		if err := m.Taskref.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskref")
			}
			return err
		}
	}

	return nil
}

func (m *Image) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
