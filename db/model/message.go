package model

import (
	"db/common"
	"fmt"
)

type Message struct {
	messageId         string
	messageSendUserId string
	messageSendDate   string
	messageChannelID  string
	messageEditFlag   bool
	messageMessage    string
}

func NewMessage(id string, messageSendUserId string, messageSendDate string, messageChannelId string, messageEditFlag bool, messageMessage string) *Message {
	return &Message{
		messageId:         id,
		messageSendUserId: messageSendUserId,
		messageSendDate:   messageSendDate,
		messageChannelID:  messageChannelId,
		messageEditFlag:   messageEditFlag,
		messageMessage:    messageMessage,
	}
}

func NewMessageWithValidation(messageSendUserId string, messageSendDate string, messageChannelId string, messageEditFlag bool, messageMessage string) (*Message, error) {
	if len(messageMessage) > 140 {
		// 事前条件違反
		return nil, common.NewPreConditionError(fmt.Sprintf("fail: invalid message(%s)", messageMessage))
	}
	return &Message{
		messageId:         idGenerator.Run(),
		messageSendUserId: messageSendUserId,
		messageSendDate:   messageSendDate,
		messageChannelID:  messageChannelId,
		messageEditFlag:   messageEditFlag,
		messageMessage:    messageMessage,
	}, nil
}

func NewMessageWithValidationNoId(messageId string, messageSendUserId string, messageSendDate string, messageChannelId string, messageEditFlag bool, messageMessage string) (*Message, error) {
	if len(messageMessage) > 140 {
		// 事前条件違反
		return nil, common.NewPreConditionError(fmt.Sprintf("fail: invalid message(%s)", messageMessage))
	}
	return &Message{
		messageId:         messageId,
		messageSendUserId: messageSendUserId,
		messageSendDate:   messageSendDate,
		messageChannelID:  messageChannelId,
		messageEditFlag:   messageEditFlag,
		messageMessage:    messageMessage,
	}, nil
}

func (m Message) MessageId() string {
	return m.messageId
}

func (m Message) MessageSendUserId() string {
	return m.messageSendUserId
}

func (m Message) MessageSendDate() string {
	return m.messageSendDate
}

func (m Message) MessageChannelId() string {
	return m.messageChannelID
}

func (m Message) MessageEditFlag() bool {
	return m.messageEditFlag
}

func (m Message) MessageMessage() string {
	return m.messageMessage
}
