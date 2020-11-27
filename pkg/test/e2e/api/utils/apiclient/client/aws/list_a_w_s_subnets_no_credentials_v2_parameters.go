// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAWSSubnetsNoCredentialsV2Params creates a new ListAWSSubnetsNoCredentialsV2Params object
// with the default values initialized.
func NewListAWSSubnetsNoCredentialsV2Params() *ListAWSSubnetsNoCredentialsV2Params {
	var ()
	return &ListAWSSubnetsNoCredentialsV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAWSSubnetsNoCredentialsV2ParamsWithTimeout creates a new ListAWSSubnetsNoCredentialsV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAWSSubnetsNoCredentialsV2ParamsWithTimeout(timeout time.Duration) *ListAWSSubnetsNoCredentialsV2Params {
	var ()
	return &ListAWSSubnetsNoCredentialsV2Params{

		timeout: timeout,
	}
}

// NewListAWSSubnetsNoCredentialsV2ParamsWithContext creates a new ListAWSSubnetsNoCredentialsV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListAWSSubnetsNoCredentialsV2ParamsWithContext(ctx context.Context) *ListAWSSubnetsNoCredentialsV2Params {
	var ()
	return &ListAWSSubnetsNoCredentialsV2Params{

		Context: ctx,
	}
}

// NewListAWSSubnetsNoCredentialsV2ParamsWithHTTPClient creates a new ListAWSSubnetsNoCredentialsV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAWSSubnetsNoCredentialsV2ParamsWithHTTPClient(client *http.Client) *ListAWSSubnetsNoCredentialsV2Params {
	var ()
	return &ListAWSSubnetsNoCredentialsV2Params{
		HTTPClient: client,
	}
}

/*ListAWSSubnetsNoCredentialsV2Params contains all the parameters to send to the API endpoint
for the list a w s subnets no credentials v2 operation typically these are written to a http.Request
*/
type ListAWSSubnetsNoCredentialsV2Params struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) WithTimeout(timeout time.Duration) *ListAWSSubnetsNoCredentialsV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) WithContext(ctx context.Context) *ListAWSSubnetsNoCredentialsV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) WithHTTPClient(client *http.Client) *ListAWSSubnetsNoCredentialsV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) WithClusterID(clusterID string) *ListAWSSubnetsNoCredentialsV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) WithProjectID(projectID string) *ListAWSSubnetsNoCredentialsV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list a w s subnets no credentials v2 params
func (o *ListAWSSubnetsNoCredentialsV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAWSSubnetsNoCredentialsV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
