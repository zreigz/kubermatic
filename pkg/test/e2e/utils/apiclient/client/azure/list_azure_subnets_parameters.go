// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListAzureSubnetsParams creates a new ListAzureSubnetsParams object
// with the default values initialized.
func NewListAzureSubnetsParams() *ListAzureSubnetsParams {
	var ()
	return &ListAzureSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureSubnetsParamsWithTimeout creates a new ListAzureSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAzureSubnetsParamsWithTimeout(timeout time.Duration) *ListAzureSubnetsParams {
	var ()
	return &ListAzureSubnetsParams{

		timeout: timeout,
	}
}

// NewListAzureSubnetsParamsWithContext creates a new ListAzureSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAzureSubnetsParamsWithContext(ctx context.Context) *ListAzureSubnetsParams {
	var ()
	return &ListAzureSubnetsParams{

		Context: ctx,
	}
}

// NewListAzureSubnetsParamsWithHTTPClient creates a new ListAzureSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAzureSubnetsParamsWithHTTPClient(client *http.Client) *ListAzureSubnetsParams {
	var ()
	return &ListAzureSubnetsParams{
		HTTPClient: client,
	}
}

/*ListAzureSubnetsParams contains all the parameters to send to the API endpoint
for the list azure subnets operation typically these are written to a http.Request
*/
type ListAzureSubnetsParams struct {

	/*ClientID*/
	ClientID *string
	/*ClientSecret*/
	ClientSecret *string
	/*Credential*/
	Credential *string
	/*ResourceGroup*/
	ResourceGroup *string
	/*SubscriptionID*/
	SubscriptionID *string
	/*TenantID*/
	TenantID *string
	/*VirtualNetwork*/
	VirtualNetwork *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list azure subnets params
func (o *ListAzureSubnetsParams) WithTimeout(timeout time.Duration) *ListAzureSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure subnets params
func (o *ListAzureSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure subnets params
func (o *ListAzureSubnetsParams) WithContext(ctx context.Context) *ListAzureSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure subnets params
func (o *ListAzureSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure subnets params
func (o *ListAzureSubnetsParams) WithHTTPClient(client *http.Client) *ListAzureSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure subnets params
func (o *ListAzureSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list azure subnets params
func (o *ListAzureSubnetsParams) WithClientID(clientID *string) *ListAzureSubnetsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list azure subnets params
func (o *ListAzureSubnetsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the list azure subnets params
func (o *ListAzureSubnetsParams) WithClientSecret(clientSecret *string) *ListAzureSubnetsParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the list azure subnets params
func (o *ListAzureSubnetsParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the list azure subnets params
func (o *ListAzureSubnetsParams) WithCredential(credential *string) *ListAzureSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list azure subnets params
func (o *ListAzureSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithResourceGroup adds the resourceGroup to the list azure subnets params
func (o *ListAzureSubnetsParams) WithResourceGroup(resourceGroup *string) *ListAzureSubnetsParams {
	o.SetResourceGroup(resourceGroup)
	return o
}

// SetResourceGroup adds the resourceGroup to the list azure subnets params
func (o *ListAzureSubnetsParams) SetResourceGroup(resourceGroup *string) {
	o.ResourceGroup = resourceGroup
}

// WithSubscriptionID adds the subscriptionID to the list azure subnets params
func (o *ListAzureSubnetsParams) WithSubscriptionID(subscriptionID *string) *ListAzureSubnetsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the list azure subnets params
func (o *ListAzureSubnetsParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the list azure subnets params
func (o *ListAzureSubnetsParams) WithTenantID(tenantID *string) *ListAzureSubnetsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list azure subnets params
func (o *ListAzureSubnetsParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithVirtualNetwork adds the virtualNetwork to the list azure subnets params
func (o *ListAzureSubnetsParams) WithVirtualNetwork(virtualNetwork *string) *ListAzureSubnetsParams {
	o.SetVirtualNetwork(virtualNetwork)
	return o
}

// SetVirtualNetwork adds the virtualNetwork to the list azure subnets params
func (o *ListAzureSubnetsParams) SetVirtualNetwork(virtualNetwork *string) {
	o.VirtualNetwork = virtualNetwork
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// header param ClientID
		if err := r.SetHeaderParam("ClientID", *o.ClientID); err != nil {
			return err
		}

	}

	if o.ClientSecret != nil {

		// header param ClientSecret
		if err := r.SetHeaderParam("ClientSecret", *o.ClientSecret); err != nil {
			return err
		}

	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}

	}

	if o.ResourceGroup != nil {

		// header param ResourceGroup
		if err := r.SetHeaderParam("ResourceGroup", *o.ResourceGroup); err != nil {
			return err
		}

	}

	if o.SubscriptionID != nil {

		// header param SubscriptionID
		if err := r.SetHeaderParam("SubscriptionID", *o.SubscriptionID); err != nil {
			return err
		}

	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}

	}

	if o.VirtualNetwork != nil {

		// header param VirtualNetwork
		if err := r.SetHeaderParam("VirtualNetwork", *o.VirtualNetwork); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
