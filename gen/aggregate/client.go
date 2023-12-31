// Code generated by goa v3.13.2, DO NOT EDIT.
//
// aggregate client
//
// Command:
// $ goa gen github.com/loafoe/iamsale/design

package aggregate

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "aggregate" service client.
type Client struct {
	AccountsEndpoint goa.Endpoint
	GroupsEndpoint   goa.Endpoint
}

// NewClient initializes a "aggregate" service client given the endpoints.
func NewClient(accounts, groups goa.Endpoint) *Client {
	return &Client{
		AccountsEndpoint: accounts,
		GroupsEndpoint:   groups,
	}
}

// Accounts calls the "accounts" endpoint of the "aggregate" service.
// Accounts may return the following errors:
//   - "PermissionDenied" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Accounts(ctx context.Context, p *AccountsPayload) (res []*Account, err error) {
	var ires any
	ires, err = c.AccountsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Account), nil
}

// Groups calls the "groups" endpoint of the "aggregate" service.
// Groups may return the following errors:
//   - "PermissionDenied" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Groups(ctx context.Context, p *GroupsPayload) (res []*Group, err error) {
	var ires any
	ires, err = c.GroupsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Group), nil
}
