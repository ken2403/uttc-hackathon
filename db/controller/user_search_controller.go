package controller

import (
	"db/model"
	"db/usecase"
	"log"
	"net/http"
)

// searchUserController - これを定義せず関数だけでもよかったが、privateメソッドを定義できて便利なので定義することにした
type searchUserController struct{}

func (receiver *searchUserController) Run(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("name")
	var users []model.User
	switch name {
	case "":
		id := r.URL.Query().Get("id")
		var err error
		users, err = usecase.SearchUser.RunById(ctx, id)
		if err != nil {
			log.Printf("fail: SearchUser.RunById, %v", err)
			writeErrorResponse(w, err)
			return
		}

	default:
		var err error
		users, err = usecase.SearchUser.Run(ctx, name)
		if err != nil {
			log.Printf("fail: SearchUser.Run, %v", err)
			writeErrorResponse(w, err)
			return
		}
	}

	// model.Userにjsonタグを付与すればこの変換は不要にはなるが、以下の理由からそれは避けている。
	// jsonを通信フォーマットに使うというのはHTTP関連技術の事情であり、アプリケーション固有の事情とは切り離しておくべき（密結合、低凝集を避けクリーンな状態にしている）
	// registerの方のcontrollerを見てもらうと分かるが、あちらはmodel.Userにjsonタグを付与しても使えない。model.Userの形とcontrollerで期待する形は必ずしも一致しない。
	// ここで一致しているのはたまたま。パッと見が同じだから使える！とやるのではなく、そうする意味を考えていくとよい。
	// DAOでuser構造体を作っている理由も同じ。
	// なお、カリキュラムでも述べたように過剰な設計とも言えるかも知れないので少なくとも課題を解く上でこの形に強制はしなくて良いです。
	results := make([]responseForUserSearch, 0, len(users))
	for _, user := range users {
		results = append(results, receiver.convert(user))
	}
	if err := writeAsJson(w, results); err != nil {
		log.Printf("fail: writeAsJson, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// controllerのメソッドにするか、convert関数を作るか迷ったけど、どちらを選んでもさほど悪くはないのでこちらで。
// なお、model.Userに変換メソッドは持たせてはいけない。
// 情報エキスパートではあるが、疎結合・高凝集に違反する(Goでは循環参照するとコンパイルエラーになるので実装したくても出来ないが)
func (receiver *searchUserController) convert(user model.User) responseForUserSearch {
	return responseForUserSearch{
		UserId:   user.UserId(),
		UserName: user.UserName(),
	}
}

// model.Userにjsonタグを付与せずにこれを別途定義した理由は上記のRunメソッド内で説明している
type responseForUserSearch struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
