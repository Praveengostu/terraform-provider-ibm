// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchInstancesInstanceIDNetworkInterfacesIDParams creates a new PatchInstancesInstanceIDNetworkInterfacesIDParams object
// with the default values initialized.
func NewPatchInstancesInstanceIDNetworkInterfacesIDParams() *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	var ()
	return &PatchInstancesInstanceIDNetworkInterfacesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithTimeout creates a new PatchInstancesInstanceIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithTimeout(timeout time.Duration) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	var ()
	return &PatchInstancesInstanceIDNetworkInterfacesIDParams{

		timeout: timeout,
	}
}

// NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithContext creates a new PatchInstancesInstanceIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithContext(ctx context.Context) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	var ()
	return &PatchInstancesInstanceIDNetworkInterfacesIDParams{

		Context: ctx,
	}
}

// NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithHTTPClient creates a new PatchInstancesInstanceIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithHTTPClient(client *http.Client) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	var ()
	return &PatchInstancesInstanceIDNetworkInterfacesIDParams{
		HTTPClient: client,
	}
}

/*PatchInstancesInstanceIDNetworkInterfacesIDParams contains all the parameters to send to the API endpoint
for the patch instances instance ID network interfaces ID operation typically these are written to a http.Request
*/
type PatchInstancesInstanceIDNetworkInterfacesIDParams struct {

	/*Body*/
	Body PatchInstancesInstanceIDNetworkInterfacesIDBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The network interface identifier

	*/
	ID string
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithTimeout(timeout time.Duration) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithContext(ctx context.Context) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithHTTPClient(client *http.Client) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithBody(body PatchInstancesInstanceIDNetworkInterfacesIDBody) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetBody(body PatchInstancesInstanceIDNetworkInterfacesIDBody) {
	o.Body = body
}

// WithGeneration adds the generation to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithGeneration(generation int64) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithID(id string) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithInstanceID(instanceID string) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WithVersion(version string) *PatchInstancesInstanceIDNetworkInterfacesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch instances instance ID network interfaces ID params
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchInstancesInstanceIDNetworkInterfacesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
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
