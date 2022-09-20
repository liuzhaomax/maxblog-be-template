package service

import (
	"context"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"maxblog-be-template/internal/core"
	"maxblog-be-template/internal/utils"
	"maxblog-be-template/src/model"
	"maxblog-be-template/src/pb"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct {
	MData *model.MData
	Tx    *core.Trans
}

func (bData *BData) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	var data *model.Data
	err := bData.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := bData.MData.QueryDataById(req, data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": utils.GetFuncName(),
		}).Info(core.FormatError(803, err).Error())
		return nil, err
	}
	res := model.Model2PB(data)
	return res, nil
}
