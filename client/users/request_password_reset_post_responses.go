// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ecadlabs/auth-client/models"
)

// RequestPasswordResetPostReader is a Reader for the RequestPasswordResetPost structure.
type RequestPasswordResetPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestPasswordResetPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRequestPasswordResetPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRequestPasswordResetPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRequestPasswordResetPostNoContent creates a RequestPasswordResetPostNoContent with default headers values
func NewRequestPasswordResetPostNoContent() *RequestPasswordResetPostNoContent {
	return &RequestPasswordResetPostNoContent{}
}

/*RequestPasswordResetPostNoContent handles this case with default header values.

Empty response
*/
type RequestPasswordResetPostNoContent struct {
}

func (o *RequestPasswordResetPostNoContent) Error() string {
	return fmt.Sprintf("[POST /request_password_reset][%d] requestPasswordResetPostNoContent ", 204)
}

func (o *RequestPasswordResetPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestPasswordResetPostDefault creates a RequestPasswordResetPostDefault with default headers values
func NewRequestPasswordResetPostDefault(code int) *RequestPasswordResetPostDefault {
	return &RequestPasswordResetPostDefault{
		_statusCode: code,
	}
}

/*RequestPasswordResetPostDefault handles this case with default header values.

Error response
*/
type RequestPasswordResetPostDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the request password reset post default response
func (o *RequestPasswordResetPostDefault) Code() int {
	return o._statusCode
}

func (o *RequestPasswordResetPostDefault) Error() string {
	return fmt.Sprintf("[POST /request_password_reset][%d] RequestPasswordResetPost default  %+v", o._statusCode, o.Payload)
}

func (o *RequestPasswordResetPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}