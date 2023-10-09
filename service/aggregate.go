package service

import (
	"context"
	"fmt"
	"github.com/loafoe/iamsale/gen/aggregate"
	"goa.design/goa/v3/security"
	"strings"
)

type Aggregate struct {
	AuthConfig
	IAMConfig
}

func NewAggregate(authConfig AuthConfig, iamConfig IAMConfig) (*Aggregate, error) {
	s := &Aggregate{
		AuthConfig: authConfig,
		IAMConfig:  iamConfig,
	}
	return s, nil
}

func (a *Aggregate) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	if !strings.EqualFold(a.Username, user) || !strings.EqualFold(a.Password, pass) {
		return ctx, fmt.Errorf("invalid username or password")
	}
	return ctx, nil
}

func (a *Aggregate) Accounts(ctx context.Context, payload *aggregate.AccountsPayload) (res []*aggregate.Account, err error) {
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

func (a *Aggregate) Groups(ctx context.Context, payload *aggregate.GroupsPayload) (res []*aggregate.Group, err error) {
	//TODO implement me
	name := "GROUPNAME"
	id := "ID"
	res = append(res, &aggregate.Group{
		Name: &name,
		GUID: &id,
	})
	return
}
