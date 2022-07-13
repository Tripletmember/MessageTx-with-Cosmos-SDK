package mt

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type (
	DocMsgMTEdit struct {
		Id      string `bson:"id"`
		DenomId string `bson:"denom_id"`
		Data    string `bson:"data"`
		Sender  string `bson:"sender"`
	}
)

func (m *DocMsgMTEdit) GetType() string {
	return MsgTypeEditMT
}

func (m *DocMsgMTEdit) BuildMsg(v interface{}) {
	msg := v.(*MsgMTEdit)

	m.Id = strings.ToLower(msg.Id)
	m.DenomId = strings.ToLower(msg.DenomId)
	m.Data = string(msg.Data)
	m.Sender = msg.Sender
}

func (m *DocMsgMTEdit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMTEdit)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
