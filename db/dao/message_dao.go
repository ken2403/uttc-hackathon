package dao

import (
	"context"
	"database/sql"
	"db/model"
)

type messageDao struct{}

// FindByUChannelName - 指定したchannelNameに合致したmessageの配列を返す
func (receiver *messageDao) FindByChannelName(ctx context.Context, channelName string) ([]model.Message, error) {
	return receiver.findByUserNameWithQueryExecutor(ctx, db, channelName)
}

// FindByNameWithTx - 指定したchannelnameに合致したmessagesを返します(Txを渡すことが出来る)
// 今回、このメソッドは使っていないが、トランザクションを使った検索も使わない検索も出来るように対応させた
// ここまで作り込まなくて大丈夫
func (receiver *messageDao) FindByNameWithTx(ctx context.Context, tx *sql.Tx, channelName string) ([]model.Message, error) {
	return receiver.findByUserNameWithQueryExecutor(ctx, tx, channelName)
}

func (receiver *messageDao) findByUserNameWithQueryExecutor(ctx context.Context, executor QueryExecutor, channelName string) ([]model.Message, error) {
	query := "select * from message where message_channel_id = (select channel_id from channel where channel_name = ?);"
	rows, err := executor.QueryContext(ctx, query, channelName)
	if err != nil {
		return nil, err
	}
	messages := make([]model.Message, 0)
	for rows.Next() {
		var m message
		if err := rows.Scan(&m.messageId, &m.messageSendUserId, &m.messageSendDate, &m.messageChannelID, &m.messageEditFlag, &m.messageMessage); err != nil {
			if err := rows.Close(); err != nil { // rowsのClose処理が必要
				return nil, err
			}
			return nil, err
		}
		messages = append(messages, m.convert())
	}

	return messages, nil
}

// Add - ユーザーモデルの永続化
// 更新系処理なのでトランザクションを必須にした
func (receiver *messageDao) Add(ctx context.Context, tx *sql.Tx, m *model.Message) error {
	query := "INSERT INTO message (message_id, message_send_user_id, message_send_date, message_channel_id, message_edit_flag, message_message) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, m.MessageId(), m.MessageSendUserId(), m.MessageSendDate(), m.MessageChannelId(), m.MessageEditFlag(), m.MessageMessage())
	return err
}

// Update
// 更新系処理なのでトランザクションを必須にした
func (receiver *messageDao) Update(ctx context.Context, tx *sql.Tx, m *model.Message) error {
	query := "UPDATE message SET message_message=?, message_edit_flag=? where message_id=?;"
	_, err := tx.ExecContext(ctx, query, m.MessageMessage(), m.MessageEditFlag(), m.MessageId())
	return err
}

// Delete
// 更新系処理なのでトランザクションを必須にした
func (receiver *messageDao) Delete(ctx context.Context, tx *sql.Tx, messageId string) error {
	query := "DELETE from message where message_id=?;"
	_, err := tx.ExecContext(ctx, query, messageId)
	return err
}

// QueryExecutor - *sql.Txと*sql.DBを同一視するためのインタフェース
type QueryExecutor interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type message struct {
	messageId         string
	messageSendUserId string
	messageSendDate   string
	messageChannelID  string
	messageEditFlag   bool
	messageMessage    string
}

func (receiver *message) convert() model.Message {
	return *model.NewMessage(receiver.messageId, receiver.messageSendUserId, receiver.messageSendDate, receiver.messageChannelID, receiver.messageEditFlag, receiver.messageMessage)
}
