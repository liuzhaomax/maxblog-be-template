package service

import (
	"context"
	"github.com/google/wire"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct {
	MData *model.MData
}

func (b *BData) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	var data *model.Data
	err := core.ExecTrans(ctx, b.MData.Tx, func(ctx context.Context) error {
		err := b.MData.QueryDataById(ctx, req, data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := model.Model2PB(data)
	return res, nil
}
