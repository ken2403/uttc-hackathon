package controller

import (
	"db/common"
	"encoding/json"
	"errors"
	"net/http"
)

var SearchUser *searchUserController
var SearchChannel *searchChannelController
var SearchMessage *searchMessageController
var RegisterMessage *registerMessageController
var UpdateMessage *updateMessageController
var DeleteMessage *deleteMessageController

func init() {
	SearchUser = &searchUserController{}
	SearchChannel = &searchChannelController{}
	SearchMessage = &searchMessageController{}
	RegisterMessage = &registerMessageController{}
	UpdateMessage = &updateMessageController{}
	DeleteMessage = &deleteMessageController{}
}

// writeAsJson - レスポンスとしてJsonを返す時の共通の処理
func writeAsJson(w http.ResponseWriter, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	return err
}

// writeErrorResponse - 事前条件違反なら400、それ以外は500を返す
func writeErrorResponse(w http.ResponseWriter, err error) {
	var e *common.PreConditionError
	if errors.As(err, &e) {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
