// Code generated by go-swagger; DO NOT EDIT.

package v_p_cs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// CreateVpcRouteReader is a Reader for the CreateVpcRoute structure.
type CreateVpcRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVpcRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateVpcRouteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateVpcRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateVpcRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateVpcRouteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVpcRouteCreated creates a CreateVpcRouteCreated with default headers values
func NewCreateVpcRouteCreated() *CreateVpcRouteCreated {
	return &CreateVpcRouteCreated{}
}

/*CreateVpcRouteCreated handles this case with default header values.

The route was created successfully.
*/
type CreateVpcRouteCreated struct {
	Payload *models.Route
}

func (o *CreateVpcRouteCreated) Error() string {
	return fmt.Sprintf("[POST /vpcs/{vpc_id}/routes][%d] createVpcRouteCreated  %+v", 201, o.Payload)
}

func (o *CreateVpcRouteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVpcRouteBadRequest creates a CreateVpcRouteBadRequest with default headers values
func NewCreateVpcRouteBadRequest() *CreateVpcRouteBadRequest {
	return &CreateVpcRouteBadRequest{}
}

/*CreateVpcRouteBadRequest handles this case with default header values.

An invalid route template was provided.
*/
type CreateVpcRouteBadRequest struct {
	Payload *models.Riaaserror
}

func (o *CreateVpcRouteBadRequest) Error() string {
	return fmt.Sprintf("[POST /vpcs/{vpc_id}/routes][%d] createVpcRouteBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVpcRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVpcRouteNotFound creates a CreateVpcRouteNotFound with default headers values
func NewCreateVpcRouteNotFound() *CreateVpcRouteNotFound {
	return &CreateVpcRouteNotFound{}
}

/*CreateVpcRouteNotFound handles this case with default header values.

The specified VPC could not be found.
*/
type CreateVpcRouteNotFound struct {
	Payload *models.Riaaserror
}

func (o *CreateVpcRouteNotFound) Error() string {
	return fmt.Sprintf("[POST /vpcs/{vpc_id}/routes][%d] createVpcRouteNotFound  %+v", 404, o.Payload)
}

func (o *CreateVpcRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVpcRouteConflict creates a CreateVpcRouteConflict with default headers values
func NewCreateVpcRouteConflict() *CreateVpcRouteConflict {
	return &CreateVpcRouteConflict{}
}

/*CreateVpcRouteConflict handles this case with default header values.

The route template conflicts with another route in the VPC.
*/
type CreateVpcRouteConflict struct {
	Payload *models.Riaaserror
}

func (o *CreateVpcRouteConflict) Error() string {
	return fmt.Sprintf("[POST /vpcs/{vpc_id}/routes][%d] createVpcRouteConflict  %+v", 409, o.Payload)
}

func (o *CreateVpcRouteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
