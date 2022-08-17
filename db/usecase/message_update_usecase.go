package usecase

import (
	"context"
	"database/sql"
	"db/dao"
	"db/model"
)

type updateMessageUseCase struct{}

// Run - messageを更新する
func (receiver *updateMessageUseCase) Run(ctx context.Context, messageId, messageSendUserId string, messageSendDate string, messageChannelId string, messageEditFlag bool, message string) (*model.Message, error) {
	newMessage, err := model.NewMessageWithValidationNoId(messageId, messageSendUserId, messageSendDate, messageChannelId, messageEditFlag, message)
	if err != nil {
		return nil, err
	}

	err = dao.WithTx(ctx, func(tx *sql.Tx) error {
		return dao.MessageDao.Update(ctx, tx, newMessage)
	})
	return newMessage, err
}
