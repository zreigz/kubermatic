// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// ListOpenstackTenantsNoCredentialsV2Reader is a Reader for the ListOpenstackTenantsNoCredentialsV2 structure.
type ListOpenstackTenantsNoCredentialsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackTenantsNoCredentialsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackTenantsNoCredentialsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackTenantsNoCredentialsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackTenantsNoCredentialsV2OK creates a ListOpenstackTenantsNoCredentialsV2OK with default headers values
func NewListOpenstackTenantsNoCredentialsV2OK() *ListOpenstackTenantsNoCredentialsV2OK {
	return &ListOpenstackTenantsNoCredentialsV2OK{}
}

/*ListOpenstackTenantsNoCredentialsV2OK handles this case with default header values.

OpenstackTenant
*/
type ListOpenstackTenantsNoCredentialsV2OK struct {
	Payload []*models.OpenstackTenant
}

func (o *ListOpenstackTenantsNoCredentialsV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/tenants][%d] listOpenstackTenantsNoCredentialsV2OK  %+v", 200, o.Payload)
}

func (o *ListOpenstackTenantsNoCredentialsV2OK) GetPayload() []*models.OpenstackTenant {
	return o.Payload
}

func (o *ListOpenstackTenantsNoCredentialsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackTenantsNoCredentialsV2Default creates a ListOpenstackTenantsNoCredentialsV2Default with default headers values
func NewListOpenstackTenantsNoCredentialsV2Default(code int) *ListOpenstackTenantsNoCredentialsV2Default {
	return &ListOpenstackTenantsNoCredentialsV2Default{
		_statusCode: code,
	}
}

/*ListOpenstackTenantsNoCredentialsV2Default handles this case with default header values.

errorResponse
*/
type ListOpenstackTenantsNoCredentialsV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack tenants no credentials v2 default response
func (o *ListOpenstackTenantsNoCredentialsV2Default) Code() int {
	return o._statusCode
}

func (o *ListOpenstackTenantsNoCredentialsV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/tenants][%d] listOpenstackTenantsNoCredentialsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackTenantsNoCredentialsV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackTenantsNoCredentialsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
