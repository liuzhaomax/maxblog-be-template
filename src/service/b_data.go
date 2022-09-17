package service

import (
	"context"
	"github.com/google/wire"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct {
	MData *model.MData
}

func (d *BData) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	data, err := d.MData.QueryDataById(ctx, req)
	if err != nil {
		return nil, err
	}
	res := model.Model2PB(data)
	return res, nil
}
