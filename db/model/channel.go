package model

type Channel struct {
	channelId          string
	channelName        string
	channelAdimnUserId string
}

func NewChannel(id string, name string, adminUserId string) *Channel {
	return &Channel{
		channelId:          id,
		channelName:        name,
		channelAdimnUserId: adminUserId,
	}
}

func NewChannelWithValidation(name string, adminUserId string) (*Channel, error) {
	return &Channel{
		channelId:          idGenerator.Run(), // IDは外部から渡してOKという世界観よりはmodelパッケージ側で生成されるのがいいだろう（PC出荷時にMACアドレスが採番されているイメージ）
		channelName:        name,
		channelAdimnUserId: adminUserId,
	}, nil
}

func (c Channel) ChannelId() string {
	return c.channelId
}

func (c Channel) ChannelName() string {
	return c.channelName
}

func (c Channel) ChannelAdminUserId() string {
	return c.channelAdimnUserId
}
