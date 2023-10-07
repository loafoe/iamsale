package main

import (
	"context"
	"github.com/loafoe/sailpoint/gen/account"
	"github.com/loafoe/sailpoint/gen/http/account/server"
	goahttp "goa.design/goa/v3/http"
	"goa.design/goa/v3/security"
	"net/http"
)

type acct struct {
}

func (a acct) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	return ctx, nil
}

func (a acct) Create(ctx context.Context, payload *account.CreatePayload) (res *account.Account, err error) {
	//TODO implement me
	id := "ec3f34c7-5142-46c3-adff-b4d3c47ec8b7"
	return &account.Account{
		ID:    &id,
		Name:  payload.Account.Name,
		Login: payload.Account.Login,
		Email: payload.Account.Email,
	}, nil
}

func (a acct) Get(ctx context.Context, payload *account.GetPayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "andy",
		Email: "andy@loafoe.com",
		Name:  "Andy Lo",
	}, nil
}

func (a acct) Update(ctx context.Context, payload *account.UpdatePayload) (res *account.Account, err error) {
	return &account.Account{
		ID:    &payload.AccountID,
		Login: "andy",
		Email: "andy@loafoe.com",
		Name:  "Andy Lo",
	}, nil
}

func (a acct) Delete(ctx context.Context, payload *account.DeletePayload) (err error) {
	return nil
}

func (a acct) GroupAdd(ctx context.Context, payload *account.GroupAddPayload) (err error) {
	//TODO implement me
	return nil
}

func (a acct) GroupRemove(ctx context.Context, payload *account.GroupRemovePayload) (err error) {
	//TODO implement me
	return nil
}

func main() {
	a := &acct{}
	endpoints := account.NewEndpoints(a)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil)
	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
