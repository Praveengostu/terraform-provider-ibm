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

// NewPostInstancesInstanceIDVolumeAttachmentsParams creates a new PostInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized.
func NewPostInstancesInstanceIDVolumeAttachmentsParams() *PostInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &PostInstancesInstanceIDVolumeAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostInstancesInstanceIDVolumeAttachmentsParamsWithTimeout creates a new PostInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostInstancesInstanceIDVolumeAttachmentsParamsWithTimeout(timeout time.Duration) *PostInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &PostInstancesInstanceIDVolumeAttachmentsParams{

		timeout: timeout,
	}
}

// NewPostInstancesInstanceIDVolumeAttachmentsParamsWithContext creates a new PostInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostInstancesInstanceIDVolumeAttachmentsParamsWithContext(ctx context.Context) *PostInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &PostInstancesInstanceIDVolumeAttachmentsParams{

		Context: ctx,
	}
}

// NewPostInstancesInstanceIDVolumeAttachmentsParamsWithHTTPClient creates a new PostInstancesInstanceIDVolumeAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostInstancesInstanceIDVolumeAttachmentsParamsWithHTTPClient(client *http.Client) *PostInstancesInstanceIDVolumeAttachmentsParams {
	var ()
	return &PostInstancesInstanceIDVolumeAttachmentsParams{
		HTTPClient: client,
	}
}

/*PostInstancesInstanceIDVolumeAttachmentsParams contains all the parameters to send to the API endpoint
for the post instances instance ID volume attachments operation typically these are written to a http.Request
*/
type PostInstancesInstanceIDVolumeAttachmentsParams struct {

	/*Body*/
	Body PostInstancesInstanceIDVolumeAttachmentsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
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

// WithTimeout adds the timeout to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithTimeout(timeout time.Duration) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithContext(ctx context.Context) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithHTTPClient(client *http.Client) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithBody(body PostInstancesInstanceIDVolumeAttachmentsBody) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetBody(body PostInstancesInstanceIDVolumeAttachmentsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithGeneration(generation int64) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithInstanceID adds the instanceID to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithInstanceID(instanceID string) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WithVersion(version string) *PostInstancesInstanceIDVolumeAttachmentsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post instances instance ID volume attachments params
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostInstancesInstanceIDVolumeAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
