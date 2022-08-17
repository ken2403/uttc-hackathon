package usecase

import (
	"context"
	"db/common"
	"db/dao"
	"db/model"
)

type searchUserUseCase struct{}

// Run - nameに合致したユーザーを返す
func (receiver *searchUserUseCase) Run(ctx context.Context, name string) ([]model.User, error) {
	if name == "" {
		// 事前条件違反
		return nil, common.NewPreConditionError("fail: name is empty")
	}

	return dao.UserDao.FindByName(ctx, name)
}

// RunById - idに合致したユーザーを返す
func (receiver *searchUserUseCase) RunById(ctx context.Context, id string) ([]model.User, error) {
	if id == "" {
		// 事前条件違反
		return nil, common.NewPreConditionError("fail: id is empty")
	}

	return dao.UserDao.FindById(ctx, id)
}
