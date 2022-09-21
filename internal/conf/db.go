package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logger "github.com/sirupsen/logrus"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/internal/utils"
	"maxblog-be-template/src/model"
	"strings"
	"time"
)

func (cfg *Config) NewDB() (*gorm.DB, func(), error) {
	cfg.DB.DSN = cfg.Mysql.DSN()
	db, err := gorm.Open(cfg.DB.Type, cfg.DB.DSN)
	if err != nil {
		return nil, nil, err
	}
	if cfg.DB.Debug {
		db = db.Debug()
	}
	clean := func() {
		err := db.Close()
		if err != nil {
			logger.WithFields(logger.Fields{
				"失败方法": utils.GetFuncName(),
			}).Error(core.FormatError(800, err).Error())
		}
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, clean, err
	}
	db.SingularTable(true)
	db.SetLogger(&GormLogger{})
	db.DB().SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.DB.MaxLifetime) * time.Second)
	return db, clean, err
}

func (cfg *Config) AutoMigrate(db *gorm.DB) error {
	dbType := strings.ToLower(cfg.DB.Type)
	if dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	db = db.AutoMigrate(new(model.Data))
	cfg.createAdmin(db)
	return db.Error
}

func (cfg *Config) createAdmin(db *gorm.DB) {
	var data model.Data
	db.First(&data)
	if data.ID != 1 {
		data.Mobile = "130123456789"
		db.Create(&data)
	}
}
