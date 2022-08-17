package model

type UserChannel struct {
	userChannelId        string
	userChannelUserId    string
	userChannelChannelId string
}

func (uc UserChannel) UserChannelId() string {
	return uc.userChannelId
}

func (uc UserChannel) UserChannelUserId() string {
	return uc.userChannelUserId
}

func (uc UserChannel) UserChannelChannelId() string {
	return uc.UserChannelChannelId()
}
