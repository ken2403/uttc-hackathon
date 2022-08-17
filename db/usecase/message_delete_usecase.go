package usecase

import (
	"context"
	"database/sql"
	"db/dao"
)

type deleteMessageUseCase struct{}

// Run - messageをdelete
func (receiver *deleteMessageUseCase) Run(ctx context.Context, messageId string) error {
	err := dao.WithTx(ctx, func(tx *sql.Tx) error {
		return dao.MessageDao.Delete(ctx, tx, messageId)
	})
	return err
}
