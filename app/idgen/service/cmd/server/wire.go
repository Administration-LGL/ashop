//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ashop/app/idgen/service/internal/biz"
	"ashop/app/idgen/service/internal/conf"
	"ashop/app/idgen/service/internal/server"
	"ashop/app/idgen/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Node, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}