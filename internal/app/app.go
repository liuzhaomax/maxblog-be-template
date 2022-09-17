package app

import (
	"github.com/jinzhu/gorm"
	logger "github.com/sirupsen/logrus"
	"maxblog-be-template/internal/conf"
	"maxblog-be-template/internal/core"
)

func InitDB() (*gorm.DB, func(), error) {
	cfg := conf.GetInstanceOfConfig()
	logger.Info(core.DB_Connection_Started)
	db, clean, err := cfg.NewDB()
	if err != nil {
		logger.Fatal(core.DB_Connection_Failed, err)
		return nil, clean, err
	}
	err = cfg.AutoMigrate(db)
	if err != nil {
		logger.Fatal(core.DB_Auto_Migration_Failed, err)
		return nil, clean, err
	}
	return db, clean, err
}
