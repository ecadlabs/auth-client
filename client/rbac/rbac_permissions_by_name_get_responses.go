// Code generated by go-swagger; DO NOT EDIT.

package rbac

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ecadlabs/auth-client/models"
)

// RbacPermissionsByNameGetReader is a Reader for the RbacPermissionsByNameGet structure.
type RbacPermissionsByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RbacPermissionsByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRbacPermissionsByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRbacPermissionsByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRbacPermissionsByNameGetOK creates a RbacPermissionsByNameGetOK with default headers values
func NewRbacPermissionsByNameGetOK() *RbacPermissionsByNameGetOK {
	return &RbacPermissionsByNameGetOK{}
}

/*RbacPermissionsByNameGetOK handles this case with default header values.

Permission info
*/
type RbacPermissionsByNameGetOK struct {
	Payload *models.PermissionDesc
}

func (o *RbacPermissionsByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /rbac/permissions/{name}][%d] rbacPermissionsByNameGetOK  %+v", 200, o.Payload)
}

func (o *RbacPermissionsByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionDesc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRbacPermissionsByNameGetDefault creates a RbacPermissionsByNameGetDefault with default headers values
func NewRbacPermissionsByNameGetDefault(code int) *RbacPermissionsByNameGetDefault {
	return &RbacPermissionsByNameGetDefault{
		_statusCode: code,
	}
}

/*RbacPermissionsByNameGetDefault handles this case with default header values.

Error response
*/
type RbacPermissionsByNameGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the rbac permissions by name get default response
func (o *RbacPermissionsByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *RbacPermissionsByNameGetDefault) Error() string {
	return fmt.Sprintf("[GET /rbac/permissions/{name}][%d] RbacPermissionsByNameGet default  %+v", o._statusCode, o.Payload)
}

func (o *RbacPermissionsByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
