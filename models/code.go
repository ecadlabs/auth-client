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

// Code Code
// swagger:model Code
type Code string

const (

	// CodeUnknown captures enum value "unknown"
	CodeUnknown Code = "unknown"

	// CodeUserNotFound captures enum value "user_not_found"
	CodeUserNotFound Code = "user_not_found"

	// CodeResourceNotFound captures enum value "resource_not_found"
	CodeResourceNotFound Code = "resource_not_found"

	// CodeEmailInUse captures enum value "email_in_use"
	CodeEmailInUse Code = "email_in_use"

	// CodePatchFormat captures enum value "patch_format"
	CodePatchFormat Code = "patch_format"

	// CodeRoleExists captures enum value "role_exists"
	CodeRoleExists Code = "role_exists"

	// CodeTokenExpired captures enum value "token_expired"
	CodeTokenExpired Code = "token_expired"

	// CodeBadRequest captures enum value "bad_request"
	CodeBadRequest Code = "bad_request"

	// CodeQuerySyntax captures enum value "query_syntax"
	CodeQuerySyntax Code = "query_syntax"

	// CodeForbidden captures enum value "forbidden"
	CodeForbidden Code = "forbidden"

	// CodeEmailFormat captures enum value "email_format"
	CodeEmailFormat Code = "email_format"

	// CodeUnauthorized captures enum value "unauthorized"
	CodeUnauthorized Code = "unauthorized"

	// CodeEmptyEmail captures enum value "empty_email"
	CodeEmptyEmail Code = "empty_email"

	// CodeEmptyPassword captures enum value "empty_password"
	CodeEmptyPassword Code = "empty_password"

	// CodeEmptyToken captures enum value "empty_token"
	CodeEmptyToken Code = "empty_token"

	// CodeInvalidToken captures enum value "invalid_token"
	CodeInvalidToken Code = "invalid_token"

	// CodeNvalidTokenFmt captures enum value "nvalid_token_fmt"
	CodeNvalidTokenFmt Code = "nvalid_token_fmt"

	// CodeInvalidAudience captures enum value "invalid_audience"
	CodeInvalidAudience Code = "invalid_audience"

	// CodeEmailNotVerified captures enum value "email_not_verified"
	CodeEmailNotVerified Code = "email_not_verified"

	// CodeRoleNotFound captures enum value "role_not_found"
	CodeRoleNotFound Code = "role_not_found"

	// CodePermissionNotFound captures enum value "permission_not_found"
	CodePermissionNotFound Code = "permission_not_found"
)

// for schema
var codeEnum []interface{}

func init() {
	var res []Code
	if err := json.Unmarshal([]byte(`["unknown","user_not_found","resource_not_found","email_in_use","patch_format","role_exists","token_expired","bad_request","query_syntax","forbidden","email_format","unauthorized","empty_email","empty_password","empty_token","invalid_token","nvalid_token_fmt","invalid_audience","email_not_verified","role_not_found","permission_not_found"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		codeEnum = append(codeEnum, v)
	}
}

func (m Code) validateCodeEnum(path, location string, value Code) error {
	if err := validate.Enum(path, location, value, codeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this code
func (m Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
