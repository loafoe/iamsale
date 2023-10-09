package service

import (
	"context"
	"fmt"
	"github.com/loafoe/iamsale/gen/aggregate"
	"github.com/philips-software/go-hsdp-api/iam"
	"goa.design/goa/v3/security"
	"strings"
)

type Aggregate struct {
	AuthConfig
	IAMConfig
	client *iam.Client
}

func NewAggregate(authConfig AuthConfig, iamConfig IAMConfig) (*Aggregate, error) {
	s := &Aggregate{
		AuthConfig: authConfig,
		IAMConfig:  iamConfig,
	}
	var err error
	s.client, err = iam.NewClient(nil, &iam.Config{
		Region:      s.Region,
		Environment: s.Environment,
	})
	if err != nil {
		return nil, fmt.Errorf("iam client create: %w", err)
	}
	err = s.client.ServiceLogin(iam.Service{
		ServiceID:  s.ServiceID,
		PrivateKey: s.ServicePrivateKey,
	})
	if err != nil {
		return nil, fmt.Errorf("service login: %w", err)
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
	users, _, err := a.client.Users.GetAllUsers(&iam.GetUserOptions{
		OrganizationID: &a.OrgID,
	})
	if err != nil {
		return res, fmt.Errorf("get users: %w", err)
	}
	for _, u := range users {
		id := u
		foundUser, _, err := a.client.Users.GetUserByID(id)
		if err != nil {
			return res, fmt.Errorf("get user %s: %w", id, err)
		}
		res = append(res, &aggregate.Account{
			Name:  fmt.Sprintf("%s %s", foundUser.Name.Given, foundUser.Name.Family),
			Login: foundUser.LoginID,
			Email: foundUser.EmailAddress,
			ID:    &id,
		})
	}
	return
}

func (a *Aggregate) Groups(ctx context.Context, payload *aggregate.GroupsPayload) (res []*aggregate.Group, err error) {
	groups, _, err := a.client.Groups.GetGroups(&iam.GetGroupOptions{
		OrganizationID: &a.OrgID,
	})
	if err != nil {
		return res, fmt.Errorf("get groups: %w", err)
	}
	for _, g := range *groups {
		groupName := g.GroupName
		groupID := g.ID
		res = append(res, &aggregate.Group{
			Name: &groupName,
			ID:   &groupID,
		})
	}
	return
}
