// Code generated by goa v3.13.2, DO NOT EDIT.
//
// aggregate service
//
// Command:
// $ goa gen github.com/loafoe/sailpoint/design

package aggregate

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Integration APIs
type Service interface {
	// Account aggregation. Returns all known accounts
	Accounts(context.Context, *AccountsPayload) (res []*Account, err error)
	// Group aggregation. Returns list of all known groups
	Groups(context.Context, *GroupsPayload) (res []*Group, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "aggregate"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"accounts", "groups"}

type Account struct {
	// Name of user
	Name *string
	// Login of user
	Login *string
	// Email of user
	Email *string
}

// AccountsPayload is the payload type of the aggregate service accounts method.
type AccountsPayload struct {
	Username string
	Password string
}

type Group struct {
	// Name of group
	Name *string
	// GUID of group
	GUID *string
}

// GroupsPayload is the payload type of the aggregate service groups method.
type GroupsPayload struct {
	Username string
	Password string
}

// MakePermissionDenied builds a goa.ServiceError from an error.
func MakePermissionDenied(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "PermissionDenied", false, false, false)
}
