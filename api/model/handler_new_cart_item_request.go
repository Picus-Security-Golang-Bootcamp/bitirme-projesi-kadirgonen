// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HandlerNewCartItemRequest handler new cart item request
//
// swagger:model handler.newCartItemRequest
type HandlerNewCartItemRequest struct {

	// number
	Number int64 `json:"number,omitempty"`

	// product id
	ProductID int64 `json:"product_id,omitempty"`
}

// Validate validates this handler new cart item request
func (m *HandlerNewCartItemRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this handler new cart item request based on context it is used
func (m *HandlerNewCartItemRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HandlerNewCartItemRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlerNewCartItemRequest) UnmarshalBinary(b []byte) error {
	var res HandlerNewCartItemRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}