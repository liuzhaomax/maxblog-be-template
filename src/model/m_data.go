package model

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(MData), "*"))

type MData struct {
	Tx *gorm.DB
}

func (m *MData) QueryDataById(ctx context.Context, req *pb.IdRequest, data *Data) error {
	result := m.Tx.First(&data, req.Id)
	if result.RowsAffected == 0 {
		return errors.New("数据没找到") // TODO 错误写入core/constants.go
	}
	return nil
}
