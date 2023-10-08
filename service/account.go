package service

import (
	"context"
	"github.com/loafoe/iamsale/gen/account"
	"goa.design/goa/v3/security"
)

type Account struct {
}

func (a Account) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	return ctx, nil
}

func (a Account) Create(ctx context.Context, payload *account.CreatePayload) (res *account.Account, err error) {
	//TODO implement me
	id := "ec3f34c7-5142-46c3-adff-b4d3c47ec8b7"
	return &account.Account{
		ID:    &id,
		Name:  payload.Account.Name,
		Login: payload.Account.Login,
		Email: payload.Account.Email,
	}, nil
}

func (a Account) Get(ctx context.Context, payload *account.GetPayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "amos",
		Email: "amos@hostedzonehere.com",
		Name:  "Amos Burton",
	}, nil
}

func (a Account) Update(ctx context.Context, payload *account.UpdatePayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "amos",
		Email: "amos@hostedzonehere.com",
		Name:  "Amos Burton",
	}, nil
}

func (a Account) Delete(ctx context.Context, payload *account.DeletePayload) (err error) {
	return nil
}

func (a Account) GroupAdd(ctx context.Context, payload *account.GroupAddPayload) (err error) {
	//TODO implement me
	return nil
}

func (a Account) GroupRemove(ctx context.Context, payload *account.GroupRemovePayload) (err error) {
	//TODO implement me
	return nil
}
