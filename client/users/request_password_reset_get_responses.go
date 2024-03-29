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

// RequestPasswordResetGetReader is a Reader for the RequestPasswordResetGet structure.
type RequestPasswordResetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestPasswordResetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRequestPasswordResetGetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRequestPasswordResetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRequestPasswordResetGetNoContent creates a RequestPasswordResetGetNoContent with default headers values
func NewRequestPasswordResetGetNoContent() *RequestPasswordResetGetNoContent {
	return &RequestPasswordResetGetNoContent{}
}

/*RequestPasswordResetGetNoContent handles this case with default header values.

Empty response
*/
type RequestPasswordResetGetNoContent struct {
}

func (o *RequestPasswordResetGetNoContent) Error() string {
	return fmt.Sprintf("[GET /request_password_reset][%d] requestPasswordResetGetNoContent ", 204)
}

func (o *RequestPasswordResetGetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestPasswordResetGetDefault creates a RequestPasswordResetGetDefault with default headers values
func NewRequestPasswordResetGetDefault(code int) *RequestPasswordResetGetDefault {
	return &RequestPasswordResetGetDefault{
		_statusCode: code,
	}
}

/*RequestPasswordResetGetDefault handles this case with default header values.

Error response
*/
type RequestPasswordResetGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the request password reset get default response
func (o *RequestPasswordResetGetDefault) Code() int {
	return o._statusCode
}

func (o *RequestPasswordResetGetDefault) Error() string {
	return fmt.Sprintf("[GET /request_password_reset][%d] RequestPasswordResetGet default  %+v", o._statusCode, o.Payload)
}

func (o *RequestPasswordResetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
