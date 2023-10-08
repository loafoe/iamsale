package main

import (
	"github.com/loafoe/iamsale/gen/account"
	"github.com/loafoe/iamsale/gen/aggregate"
	accs "github.com/loafoe/iamsale/gen/http/account/server"
	aggs "github.com/loafoe/iamsale/gen/http/aggregate/server"
	"github.com/loafoe/iamsale/service"
	goahttp "goa.design/goa/v3/http"
	"net/http"
)

func main() {
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	a := &service.Account{}
	accEndpoint := account.NewEndpoints(a)
	accServer := accs.New(accEndpoint, mux, dec, enc, nil, nil)
	accs.Mount(mux, accServer)

	b := &service.Aggregate{}
	aggEndpoint := aggregate.NewEndpoints(b)
	aggServer := aggs.New(aggEndpoint, mux, dec, enc, nil, nil)
	aggs.Mount(mux, aggServer)

	httpsvr := &http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
