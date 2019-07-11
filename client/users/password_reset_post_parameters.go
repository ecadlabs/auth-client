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

	models "github.com/ecadlabs/auth-client/models"
)

// NewPasswordResetPostParams creates a new PasswordResetPostParams object
// with the default values initialized.
func NewPasswordResetPostParams() *PasswordResetPostParams {
	var ()
	return &PasswordResetPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPasswordResetPostParamsWithTimeout creates a new PasswordResetPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPasswordResetPostParamsWithTimeout(timeout time.Duration) *PasswordResetPostParams {
	var ()
	return &PasswordResetPostParams{

		timeout: timeout,
	}
}

// NewPasswordResetPostParamsWithContext creates a new PasswordResetPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPasswordResetPostParamsWithContext(ctx context.Context) *PasswordResetPostParams {
	var ()
	return &PasswordResetPostParams{

		Context: ctx,
	}
}

// NewPasswordResetPostParamsWithHTTPClient creates a new PasswordResetPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPasswordResetPostParamsWithHTTPClient(client *http.Client) *PasswordResetPostParams {
	var ()
	return &PasswordResetPostParams{
		HTTPClient: client,
	}
}

/*PasswordResetPostParams contains all the parameters to send to the API endpoint
for the password reset post operation typically these are written to a http.Request
*/
type PasswordResetPostParams struct {

	/*Body*/
	Body *models.PasswordResetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the password reset post params
func (o *PasswordResetPostParams) WithTimeout(timeout time.Duration) *PasswordResetPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the password reset post params
func (o *PasswordResetPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the password reset post params
func (o *PasswordResetPostParams) WithContext(ctx context.Context) *PasswordResetPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the password reset post params
func (o *PasswordResetPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the password reset post params
func (o *PasswordResetPostParams) WithHTTPClient(client *http.Client) *PasswordResetPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the password reset post params
func (o *PasswordResetPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the password reset post params
func (o *PasswordResetPostParams) WithBody(body *models.PasswordResetRequest) *PasswordResetPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the password reset post params
func (o *PasswordResetPostParams) SetBody(body *models.PasswordResetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PasswordResetPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
