package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocTxMsgHarvest struct {
	PoolId string `bson:"pool_id"`
	Sender string `bson:"sender"`
}

func (m *DocTxMsgHarvest) GetType() string {
	return MsgTypeHarvest
}

func (m *DocTxMsgHarvest) BuildMsg(v interface{}) {
	msg := v.(*MsgHarvest)
	m.PoolId = msg.PoolId
	m.Sender = msg.Sender
}

func (m *DocTxMsgHarvest) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgHarvest
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
