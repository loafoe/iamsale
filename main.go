package main

import (
	"github.com/glebarez/sqlite"
	"github.com/loafoe/iamsale/gen/account"
	"github.com/loafoe/iamsale/gen/aggregate"
	accs "github.com/loafoe/iamsale/gen/http/account/server"
	aggs "github.com/loafoe/iamsale/gen/http/aggregate/server"
	ts "github.com/loafoe/iamsale/gen/http/test/server"
	"github.com/loafoe/iamsale/gen/test"
	"github.com/loafoe/iamsale/service"
	g "github.com/loafoe/iamsale/storage/gorm"
	"github.com/spf13/viper"
	goahttp "goa.design/goa/v3/http"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type Config struct {
	IAM service.IAMConfig
	API service.AuthConfig
}

func main() {
	// AuthConfig
	viper.SetEnvPrefix("iamsale")
	viper.SetConfigName("iamsale")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/iamsale")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("error reading config", "error", err)
		return
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		slog.Error("error unmarshalling config", "error", err)
		return
	}

	// Storage
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}) // file::memory:?cache=shared
	if err != nil {
		slog.Error("error opening memory database", "error", err)
		return
	}
	dbStorage, err := g.New(db)
	if err != nil {
		slog.Error("error provisioning storage layer", "error", err)
		return
	}

	// Server
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	a, err := service.NewAccount(config.API, config.IAM, dbStorage)
	if err != nil {
		slog.Error("error creating aggregate service", "error", err)
		return
	}
	accEndpoint := account.NewEndpoints(a)
	accServer := accs.New(accEndpoint, mux, dec, enc, nil, nil)
	accs.Mount(mux, accServer)

	b, err := service.NewAggregate(config.API, config.IAM, dbStorage)
	aggEndpoint := aggregate.NewEndpoints(b)
	aggServer := aggs.New(aggEndpoint, mux, dec, enc, nil, nil)
	aggs.Mount(mux, aggServer)

	t := &service.Test{}
	testEndpoint := test.NewEndpoints(t)
	testServer := ts.New(testEndpoint, mux, dec, enc, nil, nil)
	ts.Mount(mux, testServer)

	httpsvr := &http.Server{
		Addr:    "localhost:8088",
		Handler: mux,
	}
	slog.Info("starting server")
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
