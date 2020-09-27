// Code generated by go-swagger; DO NOT EDIT.

package constrainttemplates

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

// NewListConstraintTemplatesParams creates a new ListConstraintTemplatesParams object
// with the default values initialized.
func NewListConstraintTemplatesParams() *ListConstraintTemplatesParams {

	return &ListConstraintTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListConstraintTemplatesParamsWithTimeout creates a new ListConstraintTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConstraintTemplatesParamsWithTimeout(timeout time.Duration) *ListConstraintTemplatesParams {

	return &ListConstraintTemplatesParams{

		timeout: timeout,
	}
}

// NewListConstraintTemplatesParamsWithContext creates a new ListConstraintTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConstraintTemplatesParamsWithContext(ctx context.Context) *ListConstraintTemplatesParams {

	return &ListConstraintTemplatesParams{

		Context: ctx,
	}
}

// NewListConstraintTemplatesParamsWithHTTPClient creates a new ListConstraintTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConstraintTemplatesParamsWithHTTPClient(client *http.Client) *ListConstraintTemplatesParams {

	return &ListConstraintTemplatesParams{
		HTTPClient: client,
	}
}

/*ListConstraintTemplatesParams contains all the parameters to send to the API endpoint
for the list constraint templates operation typically these are written to a http.Request
*/
type ListConstraintTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list constraint templates params
func (o *ListConstraintTemplatesParams) WithTimeout(timeout time.Duration) *ListConstraintTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list constraint templates params
func (o *ListConstraintTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list constraint templates params
func (o *ListConstraintTemplatesParams) WithContext(ctx context.Context) *ListConstraintTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list constraint templates params
func (o *ListConstraintTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list constraint templates params
func (o *ListConstraintTemplatesParams) WithHTTPClient(client *http.Client) *ListConstraintTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list constraint templates params
func (o *ListConstraintTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListConstraintTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}