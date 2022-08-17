package controller

import (
	"db/model"
	"db/usecase"
	"log"
	"net/http"
)

type searchMessageController struct{}

func (receiver *searchMessageController) Run(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("channel")
	messages, err := usecase.SearchMessage.Run(ctx, name)
	if err != nil {
		log.Printf("fail: SearchMessage.Run, %v", err)
		writeErrorResponse(w, err)
		return
	}

	// jsonを通信フォーマットに使うというのはHTTP関連技術の事情であり、アプリケーション固有の事情とは切り離しておくべき（密結合、低凝集を避けクリーンな状態にしている）
	// registerの方のcontrollerを見てもらうと分かるが、あちらはmodel.Userにjsonタグを付与しても使えない。model.Userの形とcontrollerで期待する形は必ずしも一致しない。
	// ここで一致しているのはたまたま。パッと見が同じだから使える！とやるのではなく、そうする意味を考えていくとよい。
	// DAOでuser構造体を作っている理由も同じ。
	// なお、カリキュラムでも述べたように過剰な設計とも言えるかも知れないので少なくとも課題を解く上でこの形に強制はしなくて良いです。
	results := make([]responseForMessageSearch, 0, len(messages))
	for _, message := range messages {
		results = append(results, receiver.convert(message))
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
func (receiver *searchMessageController) convert(message model.Message) responseForMessageSearch {
	return responseForMessageSearch{
		MessageId:         message.MessageId(),
		MessageSendUserId: message.MessageSendUserId(),
		MessageSendDate:   message.MessageSendDate(),
		MessageChannelID:  message.MessageChannelId(),
		MessageEditFlag:   message.MessageEditFlag(),
		MessageMessage:    message.MessageMessage(),
	}
}

// model.Userにjsonタグを付与せずにこれを別途定義した理由は上記のRunメソッド内で説明している
type responseForMessageSearch struct {
	MessageId         string `json:"message_id"`
	MessageSendUserId string `json:"message_send_user_id"`
	MessageSendDate   string `json:"message_send_date"`
	MessageChannelID  string `json:"message_channel_id"`
	MessageEditFlag   bool   `json:"message_edit_flag"`
	MessageMessage    string `json:"message_message"`
}
