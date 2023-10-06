// Code generated by goa v3.13.2, DO NOT EDIT.
//
// HTTP request path constructors for the aggregate service.
//
// Command:
// $ goa gen github.com/loafoe/sailpoint/design

package server

// AccountsAggregatePath returns the URL path to the aggregate service accounts HTTP endpoint.
func AccountsAggregatePath() string {
	return "/aggregate/accounts"
}

// GroupsAggregatePath returns the URL path to the aggregate service groups HTTP endpoint.
func GroupsAggregatePath() string {
	return "/aggregate/groups"
}
