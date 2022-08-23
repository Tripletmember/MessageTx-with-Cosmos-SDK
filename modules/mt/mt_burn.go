package mt

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type (
	DocMsgMTBurn struct {
		Sender  string `bson:"sender"`
		Id      string `bson:"id"`
		Amount  string `bson:"amount"`
		DenomId string `bson:"denom_id"`
	}
)

func (m *DocMsgMTBurn) GetType() string {
	return MsgTypeBurnMT
}

func (m *DocMsgMTBurn) BuildMsg(v interface{}) {
	msg := v.(*MsgMTBurn)

	m.Sender = msg.Sender
	m.Id = strings.ToLower(msg.Id)
	m.Amount = fmt.Sprint(msg.Amount)
	m.DenomId = strings.ToLower(msg.DenomId)
}

func (m *DocMsgMTBurn) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMTBurn)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
