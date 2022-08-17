package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const dbName = "/test_database"

var db *sql.DB

var UserDao *userDao
var ChannelDao *channelDao
var MessageDao *messageDao

func init() {
	// get envs
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	// open db
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: sql.Ping, %v\n", err)
	}

	db = _db

	log.Println("success: init dao package")

	UserDao = &userDao{}
	ChannelDao = &channelDao{}
	MessageDao = &messageDao{}
}

// WithTx - トランザクションを使った一連のお決まりパターンの処理をまとめています
func WithTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil) // txOptionsを外部から受け取るように変えても良い（サボってnilにしている）
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			// エラーが2つも起きている状況でもなんとか情報を拾うため独自エラーを使っている
			return NewTxRollbackError(err, err2)
		}
		return err
	}
	return tx.Commit()
}

func CloseDB() error {
	return db.Close()
}

// TxRollbackError - ロールバック失敗時のエラーメッセージも確認できるようにするためのエラー
// ここまでやらなくて大丈夫ですw
type TxRollbackError struct {
	message string
	err     error
}

func NewTxRollbackError(originalErr, rollbackErr error) *TxRollbackError {
	return &TxRollbackError{
		message: fmt.Sprintf("%s + %s", originalErr.Error(), rollbackErr.Error()), // ラップしているエラーのメッセージも含める
		err:     originalErr,
	}
}

// Error - errorインタフェースを満たすために定義
func (receiver *TxRollbackError) Error() string {
	return receiver.message
}

// UnWrap - errors.As/Is関数で正しく動作するために定義
func (receiver *TxRollbackError) UnWrap() error {
	return receiver.err
}
