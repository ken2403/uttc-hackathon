package controller

import (
	"db/usecase"
	"log"
	"net/http"
)

// registerUserController - これを定義せず関数だけでもよかったが、privateメソッドを定義できて便利なので定義することにした
type deleteMessageController struct{}

func (receiver *deleteMessageController) Run(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")

	err := usecase.DeleteMessage.Run(ctx, id)
	if err != nil {
		log.Printf("fail: usecase.DeleteMessage, %v\n", err)
		writeErrorResponse(w, err)
		return
	}
}

type requestForMessageDelete struct {
	MessageId string `json:"message_id"`
}
