package app

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/service"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	DB      *gorm.DB
	Service *service.BData
	Model   *model.MData
}
