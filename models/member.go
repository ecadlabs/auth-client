// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Member Member
// swagger:model Member
type Member struct {

	// added
	// Required: true
	// Format: date-time
	Added *strfmt.DateTime `json:"added"`

	// email
	// Required: true
	Email *string `json:"email"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified
	// Required: true
	// Format: date-time
	Modified *strfmt.DateTime `json:"modified"`

	// roles
	Roles map[string]bool `json:"roles,omitempty"`

	// user id
	// Required: true
	// Format: uuid
	UserID *strfmt.UUID `json:"user_id"`
}

// Validate validates this member
func (m *Member) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Member) validateAdded(formats strfmt.Registry) error {

	if err := validate.Required("added", "body", m.Added); err != nil {
		return err
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateModified(formats strfmt.Registry) error {

	if err := validate.Required("modified", "body", m.Modified); err != nil {
		return err
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Member) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Member) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Member) UnmarshalBinary(b []byte) error {
	var res Member
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
