package dao

import (
	"context"
	"database/sql"
	"db/model"
)

type userDao struct{}

// FindByName - 指定したnameに合致したユーザーを返す
func (receiver *userDao) FindByName(ctx context.Context, name string) ([]model.User, error) {
	return receiver.findByNameWithQueryExecutor(ctx, db, name)
}

// FindByNameWithTx - 指定したnameに合致したユーザーを返します(Txを渡すことが出来る)
// 今回、このメソッドは使っていないが、トランザクションを使った検索も使わない検索も出来るように対応させた
// ここまで作り込まなくて大丈夫
func (receiver *userDao) FindByNameWithTx(ctx context.Context, tx *sql.Tx, name string) ([]model.User, error) {
	return receiver.findByNameWithQueryExecutor(ctx, tx, name)
}

func (receiver *userDao) findByNameWithQueryExecutor(ctx context.Context, executor QueryExecutor, name string) ([]model.User, error) {
	query := "SELECT user_id, user_name FROM user WHERE user_name = ?"
	rows, err := executor.QueryContext(ctx, query, name)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0)
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			if err := rows.Close(); err != nil { // rowsのClose処理が必要
				return nil, err
			}
			return nil, err
		}
		users = append(users, u.convert())
	}

	return users, nil
}

// FindById - 指定したidに合致したユーザーを返す
func (receiver *userDao) FindById(ctx context.Context, id string) ([]model.User, error) {
	return receiver.findByIdWithQueryExecutor(ctx, db, id)
}

// FindByIdWithTx - 指定したidに合致したユーザーを返します(Txを渡すことが出来る)
func (receiver *userDao) FindByIdWithTx(ctx context.Context, tx *sql.Tx, id string) ([]model.User, error) {
	return receiver.findByIdWithQueryExecutor(ctx, tx, id)
}

func (receiver *userDao) findByIdWithQueryExecutor(ctx context.Context, executor QueryExecutor, id string) ([]model.User, error) {
	query := "SELECT user_id, user_name FROM user WHERE user_id = ?"
	rows, err := executor.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0)
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			if err := rows.Close(); err != nil { // rowsのClose処理が必要
				return nil, err
			}
			return nil, err
		}
		users = append(users, u.convert())
	}

	return users, nil
}

// user - Userテーブルに対応
// model.Userの各プロパティを大文字にして上書き可能にすればこの定義は不要になる。直接 rows.Scan でmodel.Userが使えるようになるため。
// なお、UserテーブルとUserモデルが全く同じなので冗長に見えるが、
// DB世界とオブジェクト世界とのインピーダンスミスマッチが発生することを前提にすると必要
// また、なんらかのORマッパーを使い出すと、model.Userにメタタグを埋め込みがち。意図的にそれでよしとするなら使ってもいいが、
// DB関連の技術やHTTP関連の技術をmodelパッケージに持ち込まずクリーンな状態を保ちたければ、これを定義しておくことで回避可能になる
// controllerでレスポンス用にjsonタグを持った構造体を定義しているのも同じ理由。
type user struct {
	Id   string
	Name string
}

func (receiver *user) convert() model.User {
	return *model.NewUser(receiver.Id, receiver.Name)
}
