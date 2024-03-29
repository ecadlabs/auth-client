// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Order order
// swagger:model order
type Order string

const (

	// OrderAsc captures enum value "asc"
	OrderAsc Order = "asc"

	// OrderDesc captures enum value "desc"
	OrderDesc Order = "desc"
)

// for schema
var orderEnum []interface{}

func init() {
	var res []Order
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderEnum = append(orderEnum, v)
	}
}

func (m Order) validateOrderEnum(path, location string, value Order) error {
	if err := validate.Enum(path, location, value, orderEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this order
func (m Order) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
