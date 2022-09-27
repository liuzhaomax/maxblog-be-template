//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/service"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		core.TransSet,
		core.LoggerSet,
		model.ModelSet,
		service.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
