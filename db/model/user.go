package model

import (
	"db/common"
	"fmt"
)

type User struct {
	userId   string
	userName string
}

// NewUser - 渡ってきた値で単に生成する(DAOでDBの値から復元する時に使うことを想定)
func NewUser(id string, name string) *User {
	return &User{
		userId:   id,
		userName: name,
	}
}

// NewUserWithValidation - validation付きのコンストラクタ(新規生成時に使う)
// 実現したかった世界観としては、変な状態のモデルの生成を許さないというもの。例えばAgeがマイナスの値のUserを一度生成してその後「大丈夫？」と調べるのではなく。
// ※ usecaseで実行する、でも良いと思う。チーム内のポリシー次第。
//   usecaseでやったほうがいい点としては、「アプリケーションでその名前は1つだけであること」のような仕様を満たそうと思うとDAOを使う必要が出てくるのでmodelがDAOは使えない（usecaseが使う）
//   ただしその場合でも、モデルの構成要素の妥当性チェックはこのようにモデル側でやって、モデル単体のチェック以上のこと（アプリケーション内で名前のユニーク性を保証するなど）はusecase層で行うというのもアリ。
// ※ 長いので、単に生成する方をNewUserとし、こちらはCreateUserなどのように命名規則で分けても良い
func NewUserWithValidation(name string, age int) (*User, error) {
	if name == "" || len(name) > 50 {
		// 事前条件違反
		return nil, common.NewPreConditionError(fmt.Sprintf("fail: invalid name(%s)", name))
	}
	if age < 20 || age > 80 {
		// 事前条件違反
		return nil, common.NewPreConditionError(fmt.Sprintf("fail: invalid age(%d)", age))
	}
	return &User{
		userId:   idGenerator.Run(), // IDは外部から渡してOKという世界観よりはmodelパッケージ側で生成されるのがいいだろう（PC出荷時にMACアドレスが採番されているイメージ）
		userName: name,
	}, nil
}

func (u User) UserId() string {
	return u.userId
}

func (u User) UserName() string {
	return u.userName
}
