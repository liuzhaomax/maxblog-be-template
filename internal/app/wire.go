//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/service"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		service.ServiceSet,
		model.ModelSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
