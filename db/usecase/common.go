package usecase

var SearchUser *searchUserUseCase
var SearchChannel *searchChannelUseCase
var SearchMessage *searchMessageUseCase
var RegisterMessage *registerMessageUseCase
var UpdateMessage *updateMessageUseCase
var DeleteMessage *deleteMessageUseCase

func init() {
	SearchUser = &searchUserUseCase{}
	SearchChannel = &searchChannelUseCase{}
	SearchMessage = &searchMessageUseCase{}
	RegisterMessage = &registerMessageUseCase{}
	UpdateMessage = &updateMessageUseCase{}
	DeleteMessage = &deleteMessageUseCase{}
}
