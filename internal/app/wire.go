//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	dataModel "maxblog-be-template/src/model"
	dataService "maxblog-be-template/src/service"
)

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		dataService.ServiceSet,
		dataModel.ModelSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
