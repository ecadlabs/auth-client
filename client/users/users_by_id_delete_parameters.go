// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUsersByIDDeleteParams creates a new UsersByIDDeleteParams object
// with the default values initialized.
func NewUsersByIDDeleteParams() *UsersByIDDeleteParams {
	var ()
	return &UsersByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersByIDDeleteParamsWithTimeout creates a new UsersByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersByIDDeleteParamsWithTimeout(timeout time.Duration) *UsersByIDDeleteParams {
	var ()
	return &UsersByIDDeleteParams{

		timeout: timeout,
	}
}

// NewUsersByIDDeleteParamsWithContext creates a new UsersByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersByIDDeleteParamsWithContext(ctx context.Context) *UsersByIDDeleteParams {
	var ()
	return &UsersByIDDeleteParams{

		Context: ctx,
	}
}

// NewUsersByIDDeleteParamsWithHTTPClient creates a new UsersByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersByIDDeleteParamsWithHTTPClient(client *http.Client) *UsersByIDDeleteParams {
	var ()
	return &UsersByIDDeleteParams{
		HTTPClient: client,
	}
}

/*UsersByIDDeleteParams contains all the parameters to send to the API endpoint
for the users by Id delete operation typically these are written to a http.Request
*/
type UsersByIDDeleteParams struct {

	/*ID
	  User ID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users by Id delete params
func (o *UsersByIDDeleteParams) WithTimeout(timeout time.Duration) *UsersByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users by Id delete params
func (o *UsersByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users by Id delete params
func (o *UsersByIDDeleteParams) WithContext(ctx context.Context) *UsersByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users by Id delete params
func (o *UsersByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users by Id delete params
func (o *UsersByIDDeleteParams) WithHTTPClient(client *http.Client) *UsersByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users by Id delete params
func (o *UsersByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users by Id delete params
func (o *UsersByIDDeleteParams) WithID(id strfmt.UUID) *UsersByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users by Id delete params
func (o *UsersByIDDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
