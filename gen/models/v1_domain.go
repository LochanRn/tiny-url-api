// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Domain v1 domain
//
// swagger:model v1Domain
type V1Domain struct {

	// count
	Count int64 `json:"count,omitempty"`

	// domains
	Domains string `json:"domains,omitempty"`
}

// Validate validates this v1 domain
func (m *V1Domain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 domain based on context it is used
func (m *V1Domain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Domain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Domain) UnmarshalBinary(b []byte) error {
	var res V1Domain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
