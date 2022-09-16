package model

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Mobile string `gorm:"index:idx_mobile;unique;varchar(11);not null"`
}
