// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EmailUpdatePost updates email

Update email with given token
*/
func (a *Client) EmailUpdatePost(params *EmailUpdatePostParams) (*EmailUpdatePostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailUpdatePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EmailUpdatePost",
		Method:             "POST",
		PathPattern:        "/email_update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailUpdatePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailUpdatePostNoContent), nil

}

/*
PasswordResetPost passwords reset

Reset password
*/
func (a *Client) PasswordResetPost(params *PasswordResetPostParams) (*PasswordResetPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPasswordResetPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PasswordResetPost",
		Method:             "POST",
		PathPattern:        "/password_reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PasswordResetPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PasswordResetPostNoContent), nil

}

/*
RequestEmailUpdatePost requests email update

Request email update for specified id
*/
func (a *Client) RequestEmailUpdatePost(params *RequestEmailUpdatePostParams) (*RequestEmailUpdatePostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestEmailUpdatePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestEmailUpdatePost",
		Method:             "POST",
		PathPattern:        "/request_email_update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestEmailUpdatePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestEmailUpdatePostNoContent), nil

}

/*
RequestPasswordResetGet requests reset

Request password reset
*/
func (a *Client) RequestPasswordResetGet(params *RequestPasswordResetGetParams) (*RequestPasswordResetGetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestPasswordResetGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestPasswordResetGet",
		Method:             "GET",
		PathPattern:        "/request_password_reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestPasswordResetGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestPasswordResetGetNoContent), nil

}

/*
RequestPasswordResetPost requests reset post

Request password reset
*/
func (a *Client) RequestPasswordResetPost(params *RequestPasswordResetPostParams) (*RequestPasswordResetPostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestPasswordResetPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RequestPasswordResetPost",
		Method:             "POST",
		PathPattern:        "/request_password_reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestPasswordResetPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RequestPasswordResetPostNoContent), nil

}

/*
UsersByIDDelete deletes user

Delete user
*/
func (a *Client) UsersByIDDelete(params *UsersByIDDeleteParams) (*UsersByIDDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UsersByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersByIDDeleteNoContent), nil

}

/*
UsersByIDGet gets user

Get user info
*/
func (a *Client) UsersByIDGet(params *UsersByIDGetParams) (*UsersByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UsersByIdGet",
		Method:             "GET",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersByIDGetOK), nil

}

/*
UsersByIDPatch patches users

Modify user's properties using JSON Patch encoding
*/
func (a *Client) UsersByIDPatch(params *UsersByIDPatchParams) (*UsersByIDPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersByIDPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UsersByIdPatch",
		Method:             "PATCH",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersByIDPatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersByIDPatchOK), nil

}

/*
UsersGet gets users

List users with filtering
*/
func (a *Client) UsersGet(params *UsersGetParams) (*UsersGetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UsersGet",
		Method:             "GET",
		PathPattern:        "/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersGetNoContent), nil

}

/*
UsersPost creates user

Create new user
*/
func (a *Client) UsersPost(params *UsersPostParams) (*UsersPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UsersPost",
		Method:             "POST",
		PathPattern:        "/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsersPostCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
