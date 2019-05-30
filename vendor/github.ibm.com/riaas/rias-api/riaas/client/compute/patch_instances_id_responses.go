// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PatchInstancesIDReader is a Reader for the PatchInstancesID structure.
type PatchInstancesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchInstancesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchInstancesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchInstancesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchInstancesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchInstancesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPatchInstancesIDOK creates a PatchInstancesIDOK with default headers values
func NewPatchInstancesIDOK() *PatchInstancesIDOK {
	return &PatchInstancesIDOK{}
}

/*PatchInstancesIDOK handles this case with default header values.

dummy
*/
type PatchInstancesIDOK struct {
	Payload *models.Instance
}

func (o *PatchInstancesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /instances/{id}][%d] patchInstancesIdOK  %+v", 200, o.Payload)
}

func (o *PatchInstancesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchInstancesIDBadRequest creates a PatchInstancesIDBadRequest with default headers values
func NewPatchInstancesIDBadRequest() *PatchInstancesIDBadRequest {
	return &PatchInstancesIDBadRequest{}
}

/*PatchInstancesIDBadRequest handles this case with default header values.

error
*/
type PatchInstancesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchInstancesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /instances/{id}][%d] patchInstancesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchInstancesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchInstancesIDNotFound creates a PatchInstancesIDNotFound with default headers values
func NewPatchInstancesIDNotFound() *PatchInstancesIDNotFound {
	return &PatchInstancesIDNotFound{}
}

/*PatchInstancesIDNotFound handles this case with default header values.

error
*/
type PatchInstancesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchInstancesIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /instances/{id}][%d] patchInstancesIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchInstancesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchInstancesIDDefault creates a PatchInstancesIDDefault with default headers values
func NewPatchInstancesIDDefault(code int) *PatchInstancesIDDefault {
	return &PatchInstancesIDDefault{
		_statusCode: code,
	}
}

/*PatchInstancesIDDefault handles this case with default header values.

unexpectederror
*/
type PatchInstancesIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the patch instances ID default response
func (o *PatchInstancesIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchInstancesIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /instances/{id}][%d] PatchInstancesID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchInstancesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
