package service

import (
	"context"
	"fmt"
	"github.com/loafoe/iamsale/gen/account"
	"github.com/loafoe/iamsale/storage"
	"github.com/philips-software/go-hsdp-api/iam"
	"goa.design/goa/v3/security"
	"strings"
)

type Account struct {
	AuthConfig
	IAMConfig
	client *iam.Client
	db     storage.Store
}

func NewAccount(authConfig AuthConfig, iamConfig IAMConfig, db storage.Store) (*Account, error) {
	a := &Account{
		AuthConfig: authConfig,
		IAMConfig:  iamConfig,
		db:         db,
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
		PrivateKey: SecretFromEnv(a.ServicePrivateKey),
	})
	if err != nil {
		return nil, fmt.Errorf("service login: %w", err)
	}
	return a, nil
}

func (a *Account) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	if !strings.EqualFold(a.Username, user) || !strings.EqualFold(SecretFromEnv(a.Password), pass) {
		return ctx, fmt.Errorf("invalid username or password")
	}
	return ctx, nil
}

func (a *Account) Create(ctx context.Context, payload *account.CreatePayload) (res *account.Account, err error) {
	// Check if user already exists
	list, _, err := a.client.Users.GetAllUsers(&iam.GetUserOptions{
		LoginID: &payload.Account.Login,
	})
	if err != nil {
		return nil, fmt.Errorf("user exist check: %w", err)
	}

	status := "pending"

	acc := account.Account{
		Name:   payload.Account.Name,
		Login:  payload.Account.Login,
		Email:  payload.Account.Email,
		Status: &status,
	}
	if len(list) > 0 {
		acc.GUID = &list[0]
		status = "active"
	}
	created, err := a.db.Add(acc)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}
	return created, nil
}

func (a *Account) Get(ctx context.Context, payload *account.GetPayload) (res *account.Account, err error) {
	found, err := a.db.FindByGUID(payload.AccountID)
	if err != nil {
		return nil, fmt.Errorf("get account: %w", err)
	}
	return found, nil
}

func (a *Account) Update(ctx context.Context, payload *account.UpdatePayload) (res *account.Account, err error) {
	found, err := a.db.FindByGUID(payload.AccountID)
	if err != nil {
		return nil, fmt.Errorf("get before update account: %w", err)
	}
	found.Status = &payload.Account.Status
	updated, err := a.db.Update(*found)
	if err != nil {
		return nil, fmt.Errorf("update account: %w", err)
	}
	return updated, nil
}

func (a *Account) Delete(ctx context.Context, payload *account.DeletePayload) (err error) {
	err = a.db.Remove(account.Account{
		GUID: &payload.AccountID,
	})
	if err != nil {
		return fmt.Errorf("delete account: %w", err)
	}
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
