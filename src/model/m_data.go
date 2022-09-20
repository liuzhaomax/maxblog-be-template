package model

import (
	"errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(MData), "*"))

type MData struct {
	DB *gorm.DB
}

func (mData *MData) QueryDataById(req *pb.IdRequest, data *Data) error {
	result := mData.DB.First(data, req.Id)
	if result.RowsAffected == 0 {
		return errors.New(core.FormatInfo(803))
	}
	return nil
}
