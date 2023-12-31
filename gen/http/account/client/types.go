// Code generated by goa v3.13.2, DO NOT EDIT.
//
// account HTTP client types
//
// Command:
// $ goa gen github.com/loafoe/iamsale/design

package client

import (
	account "github.com/loafoe/iamsale/gen/account"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "account" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Login of user
	Login string `form:"login" json:"login" xml:"login"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// UpdateRequestBody is the type of the "account" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Status of user
	Status string `form:"status" json:"status" xml:"status"`
}

// CreateResponseBody is the type of the "account" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// Temporary account identifier
	ID *int64 `gorm:"autoIncrement" json:"id,omitempty"`
	// IDP account identifier
	GUID *string `json:"guid,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Login of user
	Login *string `gorm:"uniqueIndex" json:"login"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// GetResponseBody is the type of the "account" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Temporary account identifier
	ID *int64 `gorm:"autoIncrement" json:"id,omitempty"`
	// IDP account identifier
	GUID *string `json:"guid,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Login of user
	Login *string `gorm:"uniqueIndex" json:"login"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// UpdateResponseBody is the type of the "account" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// Temporary account identifier
	ID *int64 `gorm:"autoIncrement" json:"id,omitempty"`
	// IDP account identifier
	GUID *string `json:"guid,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Login of user
	Login *string `gorm:"uniqueIndex" json:"login"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// UpdateNotImplementedResponseBody is the type of the "account" service
// "update" endpoint HTTP response body for the "NotImplemented" error.
type UpdateNotImplementedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "account" service.
func NewCreateRequestBody(p *account.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:  p.Account.Name,
		Login: p.Account.Login,
		Email: p.Account.Email,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "account" service.
func NewUpdateRequestBody(p *account.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Status: p.Account.Status,
	}
	return body
}

// NewCreateAccountCreated builds a "account" service "create" endpoint result
// from a HTTP "Created" response.
func NewCreateAccountCreated(body *CreateResponseBody) *account.Account {
	v := &account.Account{
		ID:     body.ID,
		GUID:   body.GUID,
		Name:   *body.Name,
		Login:  *body.Login,
		Email:  *body.Email,
		Status: body.Status,
	}

	return v
}

// NewGetAccountOK builds a "account" service "get" endpoint result from a HTTP
// "OK" response.
func NewGetAccountOK(body *GetResponseBody) *account.Account {
	v := &account.Account{
		ID:     body.ID,
		GUID:   body.GUID,
		Name:   *body.Name,
		Login:  *body.Login,
		Email:  *body.Email,
		Status: body.Status,
	}

	return v
}

// NewUpdateAccountOK builds a "account" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateAccountOK(body *UpdateResponseBody) *account.Account {
	v := &account.Account{
		ID:     body.ID,
		GUID:   body.GUID,
		Name:   *body.Name,
		Login:  *body.Login,
		Email:  *body.Email,
		Status: body.Status,
	}

	return v
}

// NewUpdateNotImplemented builds a account service update endpoint
// NotImplemented error.
func NewUpdateNotImplemented(body *UpdateNotImplementedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Login == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("login", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "disabled" || *body.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "disabled", "pending"}))
		}
	}
	return
}

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Login == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("login", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "disabled" || *body.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "disabled", "pending"}))
		}
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Login == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("login", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "disabled" || *body.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "disabled", "pending"}))
		}
	}
	return
}

// ValidateUpdateNotImplementedResponseBody runs the validations defined on
// update_NotImplemented_response_body
func ValidateUpdateNotImplementedResponseBody(body *UpdateNotImplementedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
