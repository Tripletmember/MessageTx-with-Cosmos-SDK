package mt

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocMsgMTIssueDenom struct {
	Id     string `bson:"id"` // need get from events
	Name   string `bson:"name"`
	Data   string `bson:"data"`
	Sender string `bson:"sender"`
}

func (m *DocMsgMTIssueDenom) GetType() string {
	return MsgTypeMTIssueDenom
}

func (m *DocMsgMTIssueDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgMTIssueDenom)
	m.Name = msg.Name
	m.Data = string(msg.Data)
	m.Sender = msg.Sender
}

func (m *DocMsgMTIssueDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMTIssueDenom)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
