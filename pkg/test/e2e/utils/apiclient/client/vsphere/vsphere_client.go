// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vsphere API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vsphere API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListVSphereDatastores(params *ListVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereDatastoresOK, error)

	ListVSphereFolders(params *ListVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersOK, error)

	ListVSphereFoldersNoCredentials(params *ListVSphereFoldersNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersNoCredentialsOK, error)

	ListVSphereFoldersNoCredentialsV2(params *ListVSphereFoldersNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersNoCredentialsV2OK, error)

	ListVSphereNetworks(params *ListVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksOK, error)

	ListVSphereNetworksNoCredentials(params *ListVSphereNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksNoCredentialsOK, error)

	ListVSphereNetworksNoCredentialsV2(params *ListVSphereNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksNoCredentialsV2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListVSphereDatastores Lists datastores from vsphere datacenter
*/
func (a *Client) ListVSphereDatastores(params *ListVSphereDatastoresParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereDatastoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereDatastoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereDatastores",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/vsphere/datastores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereDatastoresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereDatastoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereDatastoresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereFolders Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFolders(params *ListVSphereFoldersParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereFolders",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereFoldersNoCredentials Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFoldersNoCredentials(params *ListVSphereFoldersNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersNoCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereFoldersNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereFoldersNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereFoldersNoCredentialsV2 Lists folders from vsphere datacenter
*/
func (a *Client) ListVSphereFoldersNoCredentialsV2(params *ListVSphereFoldersNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereFoldersNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereFoldersNoCredentialsV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereFoldersNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereFoldersNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereFoldersNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereFoldersNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereNetworks Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworks(params *ListVSphereNetworksParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereNetworks",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereNetworksNoCredentials Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworksNoCredentials(params *ListVSphereNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksNoCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereNetworksNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereNetworksNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListVSphereNetworksNoCredentialsV2 Lists networks from vsphere datacenter
*/
func (a *Client) ListVSphereNetworksNoCredentialsV2(params *ListVSphereNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter) (*ListVSphereNetworksNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVSphereNetworksNoCredentialsV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVSphereNetworksNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListVSphereNetworksNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVSphereNetworksNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVSphereNetworksNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
