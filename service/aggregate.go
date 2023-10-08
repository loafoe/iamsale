package service

import (
	"context"
	"github.com/loafoe/iamsale/gen/aggregate"
	"goa.design/goa/v3/security"
)

type Aggregate struct {
}

func (a Aggregate) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	return ctx, nil
}

func (a Aggregate) Accounts(ctx context.Context, payload *aggregate.AccountsPayload) (res []*aggregate.Account, err error) {
	//TODO implement me
	id := "foo"
	res = append(res, &aggregate.Account{
		ID:    &id,
		Name:  "Amos Burton",
		Login: "amos",
		Email: "amos@hostedzonehere.com",
	})
	return
}

func (a Aggregate) Groups(ctx context.Context, payload *aggregate.GroupsPayload) (res []*aggregate.Group, err error) {
	//TODO implement me
	name := "GROUPNAME"
	id := "ID"
	res = append(res, &aggregate.Group{
		Name: &name,
		GUID: &id,
	})
	return
}
