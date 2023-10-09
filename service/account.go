package service

import (
	"context"
	"fmt"
	"github.com/loafoe/iamsale/gen/account"
	"github.com/philips-software/go-hsdp-api/iam"
	"goa.design/goa/v3/security"
	"strings"
)

type Account struct {
	AuthConfig
	IAMConfig
	client *iam.Client
}

func NewAccount(authConfig AuthConfig, iamConfig IAMConfig) (*Account, error) {
	a := &Account{
		AuthConfig: authConfig,
		IAMConfig:  iamConfig,
	}
	var err error
	a.client, err = iam.NewClient(nil, &iam.Config{
		Region:      a.Region,
		Environment: a.Environment,
	})
	if err != nil {
		return nil, fmt.Errorf("iam client create: %w", err)
	}
	err = a.client.ServiceLogin(iam.Service{
		ServiceID:  a.ServiceID,
		PrivateKey: a.ServicePrivateKey,
	})
	if err != nil {
		return nil, fmt.Errorf("service login: %w", err)
	}
	return a, nil
}

func (a *Account) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	if !strings.EqualFold(a.Username, user) || !strings.EqualFold(a.Password, pass) {
		return ctx, fmt.Errorf("invalid username or password")
	}
	return ctx, nil
}

func (a *Account) Create(ctx context.Context, payload *account.CreatePayload) (res *account.Account, err error) {
	//TODO implement me
	id := "ec3f34c7-5142-46c3-adff-b4d3c47ec8b7"
	return &account.Account{
		ID:    &id,
		Name:  payload.Account.Name,
		Login: payload.Account.Login,
		Email: payload.Account.Email,
	}, nil
}

func (a *Account) Get(ctx context.Context, payload *account.GetPayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "amos",
		Email: "amos@hostedzonehere.com",
		Name:  "Amos Burton",
	}, nil
}

func (a *Account) Update(ctx context.Context, payload *account.UpdatePayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "amos",
		Email: "amos@hostedzonehere.com",
		Name:  "Amos Burton",
	}, nil
}

func (a *Account) Delete(ctx context.Context, payload *account.DeletePayload) (err error) {
	return nil
}

func (a *Account) GroupAdd(ctx context.Context, payload *account.GroupAddPayload) (err error) {
	//TODO implement me
	return nil
}

func (a *Account) GroupRemove(ctx context.Context, payload *account.GroupRemovePayload) (err error) {
	//TODO implement me
	return nil
}
