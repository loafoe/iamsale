// Code generated by goa v3.13.2, DO NOT EDIT.
//
// HTTP request path constructors for the account service.
//
// Command:
// $ goa gen github.com/loafoe/sailpoint/design

package client

import (
	"fmt"
)

// CreateAccountPath returns the URL path to the account service create HTTP endpoint.
func CreateAccountPath() string {
	return "/account"
}

// DeleteAccountPath returns the URL path to the account service delete HTTP endpoint.
func DeleteAccountPath(accountID string) string {
	return fmt.Sprintf("/account/%v", accountID)
}
