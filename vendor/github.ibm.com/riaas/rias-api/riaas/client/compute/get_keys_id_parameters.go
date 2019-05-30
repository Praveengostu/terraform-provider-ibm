// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKeysIDParams creates a new GetKeysIDParams object
// with the default values initialized.
func NewGetKeysIDParams() *GetKeysIDParams {
	var ()
	return &GetKeysIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysIDParamsWithTimeout creates a new GetKeysIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysIDParamsWithTimeout(timeout time.Duration) *GetKeysIDParams {
	var ()
	return &GetKeysIDParams{

		timeout: timeout,
	}
}

// NewGetKeysIDParamsWithContext creates a new GetKeysIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysIDParamsWithContext(ctx context.Context) *GetKeysIDParams {
	var ()
	return &GetKeysIDParams{

		Context: ctx,
	}
}

// NewGetKeysIDParamsWithHTTPClient creates a new GetKeysIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysIDParamsWithHTTPClient(client *http.Client) *GetKeysIDParams {
	var ()
	return &GetKeysIDParams{
		HTTPClient: client,
	}
}

/*GetKeysIDParams contains all the parameters to send to the API endpoint
for the get keys ID operation typically these are written to a http.Request
*/
type GetKeysIDParams struct {

	/*ID
	  The key idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys ID params
func (o *GetKeysIDParams) WithTimeout(timeout time.Duration) *GetKeysIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys ID params
func (o *GetKeysIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys ID params
func (o *GetKeysIDParams) WithContext(ctx context.Context) *GetKeysIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys ID params
func (o *GetKeysIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys ID params
func (o *GetKeysIDParams) WithHTTPClient(client *http.Client) *GetKeysIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys ID params
func (o *GetKeysIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get keys ID params
func (o *GetKeysIDParams) WithID(id string) *GetKeysIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get keys ID params
func (o *GetKeysIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get keys ID params
func (o *GetKeysIDParams) WithVersion(version string) *GetKeysIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get keys ID params
func (o *GetKeysIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
