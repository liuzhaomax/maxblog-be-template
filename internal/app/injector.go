package app

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	DB *gorm.DB
}
