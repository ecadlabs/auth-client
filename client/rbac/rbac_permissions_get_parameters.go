// Code generated by go-swagger; DO NOT EDIT.

package rbac

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRbacPermissionsGetParams creates a new RbacPermissionsGetParams object
// with the default values initialized.
func NewRbacPermissionsGetParams() *RbacPermissionsGetParams {
	var ()
	return &RbacPermissionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRbacPermissionsGetParamsWithTimeout creates a new RbacPermissionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRbacPermissionsGetParamsWithTimeout(timeout time.Duration) *RbacPermissionsGetParams {
	var ()
	return &RbacPermissionsGetParams{

		timeout: timeout,
	}
}

// NewRbacPermissionsGetParamsWithContext creates a new RbacPermissionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRbacPermissionsGetParamsWithContext(ctx context.Context) *RbacPermissionsGetParams {
	var ()
	return &RbacPermissionsGetParams{

		Context: ctx,
	}
}

// NewRbacPermissionsGetParamsWithHTTPClient creates a new RbacPermissionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRbacPermissionsGetParamsWithHTTPClient(client *http.Client) *RbacPermissionsGetParams {
	var ()
	return &RbacPermissionsGetParams{
		HTTPClient: client,
	}
}

/*RbacPermissionsGetParams contains all the parameters to send to the API endpoint
for the rbac permissions get operation typically these are written to a http.Request
*/
type RbacPermissionsGetParams struct {

	/*Role
	  Filter by role

	*/
	Role []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rbac permissions get params
func (o *RbacPermissionsGetParams) WithTimeout(timeout time.Duration) *RbacPermissionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rbac permissions get params
func (o *RbacPermissionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rbac permissions get params
func (o *RbacPermissionsGetParams) WithContext(ctx context.Context) *RbacPermissionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rbac permissions get params
func (o *RbacPermissionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rbac permissions get params
func (o *RbacPermissionsGetParams) WithHTTPClient(client *http.Client) *RbacPermissionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rbac permissions get params
func (o *RbacPermissionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the rbac permissions get params
func (o *RbacPermissionsGetParams) WithRole(role []string) *RbacPermissionsGetParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the rbac permissions get params
func (o *RbacPermissionsGetParams) SetRole(role []string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *RbacPermissionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesRole := o.Role

	joinedRole := swag.JoinByFormat(valuesRole, "")
	// query array param role
	if err := r.SetQueryParam("role", joinedRole...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
