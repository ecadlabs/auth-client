// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ecadlabs/auth-client/models"
)

// LoginGetReader is a Reader for the LoginGet structure.
type LoginGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewLoginGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginGetOK creates a LoginGetOK with default headers values
func NewLoginGetOK() *LoginGetOK {
	return &LoginGetOK{}
}

/*LoginGetOK handles this case with default header values.

Token response
*/
type LoginGetOK struct {
	Payload *models.Token
}

func (o *LoginGetOK) Error() string {
	return fmt.Sprintf("[GET /login][%d] loginGetOK  %+v", 200, o.Payload)
}

func (o *LoginGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginGetDefault creates a LoginGetDefault with default headers values
func NewLoginGetDefault(code int) *LoginGetDefault {
	return &LoginGetDefault{
		_statusCode: code,
	}
}

/*LoginGetDefault handles this case with default header values.

Error response
*/
type LoginGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the login get default response
func (o *LoginGetDefault) Code() int {
	return o._statusCode
}

func (o *LoginGetDefault) Error() string {
	return fmt.Sprintf("[GET /login][%d] LoginGet default  %+v", o._statusCode, o.Payload)
}

func (o *LoginGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
