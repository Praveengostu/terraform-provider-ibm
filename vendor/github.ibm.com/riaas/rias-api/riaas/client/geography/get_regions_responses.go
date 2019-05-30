// Code generated by go-swagger; DO NOT EDIT.

package geography

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/*GetRegionsOK handles this case with default header values.

dummy
*/
type GetRegionsOK struct {
	Payload *models.GetRegionsOKBody
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsDefault creates a GetRegionsDefault with default headers values
func NewGetRegionsDefault(code int) *GetRegionsDefault {
	return &GetRegionsDefault{
		_statusCode: code,
	}
}

/*GetRegionsDefault handles this case with default header values.

unexpectederror
*/
type GetRegionsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get regions default response
func (o *GetRegionsDefault) Code() int {
	return o._statusCode
}

func (o *GetRegionsDefault) Error() string {
	return fmt.Sprintf("[GET /regions][%d] GetRegions default  %+v", o._statusCode, o.Payload)
}

func (o *GetRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
