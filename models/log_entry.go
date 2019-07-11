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

// LogEntry LogEntry
// swagger:model LogEntry
type LogEntry struct {

	// addr
	Addr string `json:"addr,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// event
	// Required: true
	Event Event `json:"event"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// msg
	Msg string `json:"msg,omitempty"`

	// target id
	// Format: uuid
	TargetID strfmt.UUID `json:"target_id,omitempty"`

	// ts
	// Required: true
	// Format: date-time
	Ts *strfmt.DateTime `json:"ts"`

	// user id
	// Format: uuid
	UserID strfmt.UUID `json:"user_id,omitempty"`
}

// Validate validates this log entry
func (m *LogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTs(formats); err != nil {
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

func (m *LogEntry) validateEvent(formats strfmt.Registry) error {

	if err := m.Event.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("event")
		}
		return err
	}

	return nil
}

func (m *LogEntry) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogEntry) validateTargetID(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetID) { // not required
		return nil
	}

	if err := validate.FormatOf("target_id", "body", "uuid", m.TargetID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogEntry) validateTs(formats strfmt.Registry) error {

	if err := validate.Required("ts", "body", m.Ts); err != nil {
		return err
	}

	if err := validate.FormatOf("ts", "body", "date-time", m.Ts.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogEntry) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntry) UnmarshalBinary(b []byte) error {
	var res LogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}