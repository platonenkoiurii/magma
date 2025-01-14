// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlatformInfo platform info
//
// swagger:model platform_info
type PlatformInfo struct {

	// config info
	ConfigInfo *ConfigInfo `json:"config_info,omitempty"`

	// kernel version
	// Example: 4.9.0-6-amd64
	KernelVersion string `json:"kernel_version,omitempty"`

	// kernel versions installed
	// Example: ["4.9.0-6-amd64","4.9.0-7-amd64"]
	KernelVersionsInstalled []string `json:"kernel_versions_installed,omitempty"`

	// packages
	Packages []*Package `json:"packages,omitempty"`

	// vpn ip
	// Example: 10.0.0.1
	VpnIP string `json:"vpn_ip,omitempty" magma_alt_name:"VpnIp"`
}

// Validate validates this platform info
func (m *PlatformInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformInfo) validateConfigInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigInfo) { // not required
		return nil
	}

	if m.ConfigInfo != nil {
		if err := m.ConfigInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config_info")
			}
			return err
		}
	}

	return nil
}

func (m *PlatformInfo) validatePackages(formats strfmt.Registry) error {
	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this platform info based on the context it is used
func (m *PlatformInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformInfo) contextValidateConfigInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigInfo != nil {
		if err := m.ConfigInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config_info")
			}
			return err
		}
	}

	return nil
}

func (m *PlatformInfo) contextValidatePackages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Packages); i++ {

		if m.Packages[i] != nil {
			if err := m.Packages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformInfo) UnmarshalBinary(b []byte) error {
	var res PlatformInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
