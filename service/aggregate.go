package service

import (
	"context"
	"fmt"
	"github.com/loafoe/iamsale/gen/aggregate"
	"github.com/loafoe/iamsale/storage"
	"github.com/philips-software/go-hsdp-api/iam"
	"goa.design/goa/v3/security"
	"strings"
)

type Aggregate struct {
	AuthConfig
	IAMConfig
	client *iam.Client
	db     storage.Store
}

func NewAggregate(authConfig AuthConfig, iamConfig IAMConfig, db storage.Store) (*Aggregate, error) {
	s := &Aggregate{
		AuthConfig: authConfig,
		IAMConfig:  iamConfig,
		db:         db,
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
	accounts, err := a.db.FindAll()
	if err != nil {
		return res, err
	}
	for _, a := range accounts {
		t := aggregate.Account(*a)
		res = append(res, &t)
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
