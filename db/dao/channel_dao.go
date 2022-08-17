package dao

import (
	"context"
	"database/sql"
	"db/model"
)

type channelDao struct{}

// FindByUserName - 指定したusernameに合致したchannelの配列を返す
func (receiver *channelDao) FindByUserName(ctx context.Context, userName string) ([]model.Channel, error) {
	return receiver.findByUserNameWithQueryExecutor(ctx, db, userName)
}

// FindByNameWithTx - 指定したnameに合致したユーザーを返します(Txを渡すことが出来る)
// 今回、このメソッドは使っていないが、トランザクションを使った検索も使わない検索も出来るように対応させた
// ここまで作り込まなくて大丈夫
func (receiver *channelDao) FindByNameWithTx(ctx context.Context, tx *sql.Tx, name string) ([]model.Channel, error) {
	return receiver.findByUserNameWithQueryExecutor(ctx, tx, name)
}

func (receiver *channelDao) findByUserNameWithQueryExecutor(ctx context.Context, executor QueryExecutor, userName string) ([]model.Channel, error) {
	query := "select * from channel where channel_id in (select user_channel_channel_id from user_channel where user_channel_user_id = (select user_id from user where user_name = ?));"
	rows, err := executor.QueryContext(ctx, query, userName)
	if err != nil {
		return nil, err
	}
	channels := make([]model.Channel, 0)
	for rows.Next() {
		var c channel
		if err := rows.Scan(&c.channelId, &c.channelName, &c.channelAdminUserId); err != nil {
			if err := rows.Close(); err != nil { // rowsのClose処理が必要
				return nil, err
			}
			return nil, err
		}
		channels = append(channels, c.convert())
	}

	return channels, nil
}

//// Add - ユーザーモデルの永続化
//// 更新系処理なのでトランザクションを必須にした
//func (receiver *channelDao) Add(ctx context.Context, tx *sql.Tx, c *model.Channel) error {
//	query := "INSERT INTO user (channel_id, channel_name, channel_adimn_user_id) VALUES(?, ?, ?)"
//	_, err := tx.ExecContext(ctx, query, u.UserId(), u.UserName())
//	return err
//}

//// QueryExecutor - *sql.Txと*sql.DBを同一視するためのインタフェース
//type QueryExecutor interface {
//	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
//}

type channel struct {
	channelId          string
	channelName        string
	channelAdminUserId string
}

func (receiver *channel) convert() model.Channel {
	return *model.NewChannel(receiver.channelId, receiver.channelName, receiver.channelAdminUserId)
}
