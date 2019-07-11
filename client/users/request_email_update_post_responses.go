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

// RequestEmailUpdatePostReader is a Reader for the RequestEmailUpdatePost structure.
type RequestEmailUpdatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestEmailUpdatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRequestEmailUpdatePostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRequestEmailUpdatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRequestEmailUpdatePostNoContent creates a RequestEmailUpdatePostNoContent with default headers values
func NewRequestEmailUpdatePostNoContent() *RequestEmailUpdatePostNoContent {
	return &RequestEmailUpdatePostNoContent{}
}

/*RequestEmailUpdatePostNoContent handles this case with default header values.

Empty response
*/
type RequestEmailUpdatePostNoContent struct {
}

func (o *RequestEmailUpdatePostNoContent) Error() string {
	return fmt.Sprintf("[POST /request_email_update][%d] requestEmailUpdatePostNoContent ", 204)
}

func (o *RequestEmailUpdatePostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestEmailUpdatePostDefault creates a RequestEmailUpdatePostDefault with default headers values
func NewRequestEmailUpdatePostDefault(code int) *RequestEmailUpdatePostDefault {
	return &RequestEmailUpdatePostDefault{
		_statusCode: code,
	}
}

/*RequestEmailUpdatePostDefault handles this case with default header values.

Error response
*/
type RequestEmailUpdatePostDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the request email update post default response
func (o *RequestEmailUpdatePostDefault) Code() int {
	return o._statusCode
}

func (o *RequestEmailUpdatePostDefault) Error() string {
	return fmt.Sprintf("[POST /request_email_update][%d] RequestEmailUpdatePost default  %+v", o._statusCode, o.Payload)
}

func (o *RequestEmailUpdatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}