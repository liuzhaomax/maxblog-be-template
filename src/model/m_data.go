package model

import (
	"errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(MData), "*"))

type MData struct {
	DB *gorm.DB
}

func (mData *MData) QueryDataById(req *pb.IdRequest, data *Data) error {
	result := mData.DB.First(&data, req.Id)
	if result.RowsAffected == 0 {
		return errors.New("数据没找到") // TODO 错误写入core/constants.go
	}
	return nil
}
