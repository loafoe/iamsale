// Code generated by goa v3.13.2, DO NOT EDIT.
//
// account HTTP server types
//
// Command:
// $ goa gen github.com/loafoe/iamsale/design

package server

import (
	account "github.com/loafoe/iamsale/gen/account"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "account" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Login of user
	Login *string `form:"login,omitempty" json:"login,omitempty" xml:"login,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// UpdateRequestBody is the type of the "account" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Status of user
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// CreateResponseBody is the type of the "account" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID of account
	ID *string `gorm:"primaryKey" json:"id,omitempty"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Login of user
	Login string `form:"login" json:"login" xml:"login"`
	// Email of user
	Email string `gorm:"index"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// GetResponseBody is the type of the "account" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// ID of account
	ID *string `gorm:"primaryKey" json:"id,omitempty"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Login of user
	Login string `form:"login" json:"login" xml:"login"`
	// Email of user
	Email string `gorm:"index"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// UpdateResponseBody is the type of the "account" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID of account
	ID *string `gorm:"primaryKey" json:"id,omitempty"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Login of user
	Login string `form:"login" json:"login" xml:"login"`
	// Email of user
	Email string `gorm:"index"`
	// Status of account
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "account" service.
func NewCreateResponseBody(res *account.Account) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:     res.ID,
		Name:   res.Name,
		Login:  res.Login,
		Email:  res.Email,
		Status: res.Status,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "account" service.
func NewGetResponseBody(res *account.Account) *GetResponseBody {
	body := &GetResponseBody{
		ID:     res.ID,
		Name:   res.Name,
		Login:  res.Login,
		Email:  res.Email,
		Status: res.Status,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "account" service.
func NewUpdateResponseBody(res *account.Account) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:     res.ID,
		Name:   res.Name,
		Login:  res.Login,
		Email:  res.Email,
		Status: res.Status,
	}
	return body
}

// NewCreatePayload builds a account service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody) *account.CreatePayload {
	v := &account.CreateAccount{
		Name:  *body.Name,
		Login: *body.Login,
		Email: *body.Email,
	}
	res := &account.CreatePayload{
		Account: v,
	}

	return res
}

// NewGetPayload builds a account service get endpoint payload.
func NewGetPayload(accountID string) *account.GetPayload {
	v := &account.GetPayload{}
	v.AccountID = accountID

	return v
}

// NewUpdatePayload builds a account service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, accountID string) *account.UpdatePayload {
	v := &account.UpdateAccount{
		Status: *body.Status,
	}
	res := &account.UpdatePayload{
		Account: v,
	}
	res.AccountID = accountID

	return res
}

// NewDeletePayload builds a account service delete endpoint payload.
func NewDeletePayload(accountID string) *account.DeletePayload {
	v := &account.DeletePayload{}
	v.AccountID = accountID

	return v
}

// NewGroupAddPayload builds a account service groupAdd endpoint payload.
func NewGroupAddPayload(accountID string, groupID string) *account.GroupAddPayload {
	v := &account.GroupAddPayload{}
	v.AccountID = accountID
	v.GroupID = groupID

	return v
}

// NewGroupRemovePayload builds a account service groupRemove endpoint payload.
func NewGroupRemovePayload(accountID string, groupID string) *account.GroupRemovePayload {
	v := &account.GroupRemovePayload{}
	v.AccountID = accountID
	v.GroupID = groupID

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Login == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("login", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "active" || *body.Status == "disabled") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"active", "disabled"}))
		}
	}
	return
}
