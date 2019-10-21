// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// NewPcloudPvminstancesVolumesDeleteParams creates a new PcloudPvminstancesVolumesDeleteParams object
// with the default values initialized.
func NewPcloudPvminstancesVolumesDeleteParams() *PcloudPvminstancesVolumesDeleteParams {
	var ()
	return &PcloudPvminstancesVolumesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesVolumesDeleteParamsWithTimeout creates a new PcloudPvminstancesVolumesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesVolumesDeleteParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesDeleteParams {
	var ()
	return &PcloudPvminstancesVolumesDeleteParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesVolumesDeleteParamsWithContext creates a new PcloudPvminstancesVolumesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesVolumesDeleteParamsWithContext(ctx context.Context) *PcloudPvminstancesVolumesDeleteParams {
	var ()
	return &PcloudPvminstancesVolumesDeleteParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesVolumesDeleteParamsWithHTTPClient creates a new PcloudPvminstancesVolumesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesVolumesDeleteParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesDeleteParams {
	var ()
	return &PcloudPvminstancesVolumesDeleteParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesVolumesDeleteParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances volumes delete operation typically these are written to a http.Request
*/
type PcloudPvminstancesVolumesDeleteParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string
	/*VolumeID
	  Volume ID

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithContext(ctx context.Context) *PcloudPvminstancesVolumesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesVolumesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesVolumesDeleteParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WithVolumeID adds the volumeID to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) WithVolumeID(volumeID string) *PcloudPvminstancesVolumesDeleteParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the pcloud pvminstances volumes delete params
func (o *PcloudPvminstancesVolumesDeleteParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesVolumesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
