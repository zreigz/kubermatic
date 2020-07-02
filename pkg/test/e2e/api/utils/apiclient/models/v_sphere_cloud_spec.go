// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VSphereCloudSpec VSphereCloudSpec specifies access data to VSphere cloud.
//
// swagger:model VSphereCloudSpec
type VSphereCloudSpec struct {

	// folder
	Folder string `json:"folder,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// VM net name
	VMNetName string `json:"vmNetName,omitempty"`

	// credentials reference
	CredentialsReference GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// infra management user
	InfraManagementUser *VSphereCredentials `json:"infraManagementUser,omitempty"`
}

// Validate validates this v sphere cloud spec
func (m *VSphereCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraManagementUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VSphereCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if err := m.CredentialsReference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentialsReference")
		}
		return err
	}

	return nil
}

func (m *VSphereCloudSpec) validateInfraManagementUser(formats strfmt.Registry) error {

	if swag.IsZero(m.InfraManagementUser) { // not required
		return nil
	}

	if m.InfraManagementUser != nil {
		if err := m.InfraManagementUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infraManagementUser")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VSphereCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VSphereCloudSpec) UnmarshalBinary(b []byte) error {
	var res VSphereCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}