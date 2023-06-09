package oracle

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocMsgEditFeed struct {
	FeedName          string        `bson:"feed_name" yaml:"feed_name"`
	LatestHistory     int64         `bson:"latest_history" yaml:"latest_history"`
	Description       string        `bson:"description"`
	Creator           string        `bson:"creator"`
	Providers         []string      `bson:"providers"`
	Timeout           int64         `bson:"timeout"`
	ServiceFeeCap     []models.Coin `bson:"service_fee_cap" yaml:"service_fee_cap"`
	RepeatedFrequency int64         `bson:"repeated_frequency" yaml:"repeated_frequency"`
	ResponseThreshold uint32        `bson:"response_threshold" yaml:"response_threshold"`
}

func (m *DocMsgEditFeed) GetType() string {
	return TxTypeEditFeed
}

func (m *DocMsgEditFeed) BuildMsg(v interface{}) {
	msg := v.(*MsgEditFeed)

	m.FeedName = msg.FeedName
	m.LatestHistory = int64(msg.LatestHistory)
	m.Description = msg.Description
	m.Creator = msg.Creator
	m.Providers = msg.GetProviders()
	m.Timeout = msg.Timeout
	m.ServiceFeeCap = models.BuildDocCoins(msg.ServiceFeeCap)
	m.RepeatedFrequency = int64(msg.RepeatedFrequency)
	m.ResponseThreshold = msg.ResponseThreshold
}

func (m *DocMsgEditFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgEditFeed
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	addrs = append(addrs, msg.GetProviders()...)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
