// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HandlerNewOrderRequest handler new order request
//
// swagger:model handler.newOrderRequest
type HandlerNewOrderRequest struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// name
	// Required: true
	Name *string `json:"name"`

	// phone
	// Required: true
	Phone *string `json:"phone"`
}

// Validate validates this handler new order request
func (m *HandlerNewOrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandlerNewOrderRequest) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *HandlerNewOrderRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HandlerNewOrderRequest) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this handler new order request based on context it is used
func (m *HandlerNewOrderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HandlerNewOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlerNewOrderRequest) UnmarshalBinary(b []byte) error {
	var res HandlerNewOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
