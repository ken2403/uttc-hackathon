package usecase

import (
	"context"
	"db/common"
	"db/dao"
	"db/model"
)

type searchMessageUseCase struct{}

// Run - nameに合致したユーザーを返す
func (receiver *searchMessageUseCase) Run(ctx context.Context, channelName string) ([]model.Message, error) {
	if channelName == "" {
		// 事前条件違反
		return nil, common.NewPreConditionError("fail: channel name is empty")
	}

	return dao.MessageDao.FindByChannelName(ctx, channelName)
}
