package usecase

import (
	"context"
	"database/sql"
	"db/dao"
	"db/model"
)

type registerMessageUseCase struct{}

// Run - messageを登録する
func (receiver *registerMessageUseCase) Run(ctx context.Context, messageSendUserId string, messageSendDate string, messageChannelId string, messageEditFlag bool, message string) (*model.Message, error) {
	newMessage, err := model.NewMessageWithValidation(messageSendUserId, messageSendDate, messageChannelId, messageEditFlag, message)
	if err != nil {
		return nil, err
	}

	err = dao.WithTx(ctx, func(tx *sql.Tx) error {
		return dao.MessageDao.Add(ctx, tx, newMessage)
	})
	return newMessage, err
}
