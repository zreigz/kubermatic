// Code generated by go-swagger; DO NOT EDIT.

package metric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metric API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metric API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListMachineDeploymentMetrics(params *ListMachineDeploymentMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*ListMachineDeploymentMetricsOK, error)

	ListNodeDeploymentMetrics(params *ListNodeDeploymentMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodeDeploymentMetricsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListMachineDeploymentMetrics lists metrics that belong to the given machine deployment
*/
func (a *Client) ListMachineDeploymentMetrics(params *ListMachineDeploymentMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*ListMachineDeploymentMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMachineDeploymentMetricsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listMachineDeploymentMetrics",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListMachineDeploymentMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMachineDeploymentMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListMachineDeploymentMetricsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListNodeDeploymentMetrics lists metrics that belong to the given node deployment
*/
func (a *Client) ListNodeDeploymentMetrics(params *ListNodeDeploymentMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodeDeploymentMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodeDeploymentMetricsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodeDeploymentMetrics",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodeDeploymentMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListNodeDeploymentMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNodeDeploymentMetricsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
