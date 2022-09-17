package service

import (
	"context"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/pb"
)

type DataServer struct {
}

func (d *DataServer) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	var data model.Data

	res := model.Model2PB(data)
	return res, nil
}
