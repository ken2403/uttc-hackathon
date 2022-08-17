package controller

import (
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// registerUserController - これを定義せず関数だけでもよかったが、privateメソッドを定義できて便利なので定義することにした
type updateMessageController struct{}

func (receiver *updateMessageController) Run(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var m responseForMessageUpdate
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("fail: json.Decoder.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message, err := usecase.UpdateMessage.Run(ctx, m.MessageId, m.MessageSendUserId, m.MessageSendDate, m.MessageChannelId, m.MessageEditFlag, m.MessageMessage)
	if err != nil {
		log.Printf("fail: usecase.UpdateMessage, %v\n", err)
		writeErrorResponse(w, err)
		return
	}

	if err := writeAsJson(w, receiver.convert(message)); err != nil {
		log.Printf("fail: writeAsJson, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// controllerのメソッドにするか、convert関数を作るか迷ったけど、どちらを選んでもさほど悪くはないのでこちらで。
// なお、model.Userに変換メソッドは持たせてはいけない。
// 情報エキスパートではあるが、疎結合・高凝集に違反する(Goでは循環参照するとコンパイルエラーになるので実装したくても出来ないが)
func (receiver *updateMessageController) convert(m *model.Message) responseForMessageUpdate {
	return responseForMessageUpdate{
		MessageId:         m.MessageId(),
		MessageSendUserId: m.MessageSendUserId(),
		MessageSendDate:   m.MessageSendDate(),
		MessageEditFlag:   m.MessageEditFlag(),
		MessageChannelId:  m.MessageChannelId(),
		MessageMessage:    m.MessageMessage(),
	}
}

type responseForMessageUpdate struct {
	MessageId         string `json:"message_id"`
	MessageSendUserId string `json:"message_send_user_id"`
	MessageSendDate   string `json:"message_send_date"`
	MessageChannelId  string `json:"message_channel_id"`
	MessageEditFlag   bool   `json:"message_edit_flag"`
	MessageMessage    string `json:"message_message"`
}
