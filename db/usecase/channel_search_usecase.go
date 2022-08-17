package usecase

import (
	"context"
	"db/common"
	"db/dao"
	"db/model"
)

type searchChannelUseCase struct{}

// Run - nameに合致したユーザーを返す
func (receiver *searchChannelUseCase) Run(ctx context.Context, userName string) ([]model.Channel, error) {
	if userName == "" {
		// 事前条件違反
		return nil, common.NewPreConditionError("fail: user name is empty")
	}

	return dao.ChannelDao.FindByUserName(ctx, userName)
}
