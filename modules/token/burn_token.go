package token

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

// MsgBurnToken defines an SDK message for burning some tokens.
type DocMsgBurnToken struct {
	Symbol string `bson:"symbol"`
	Amount string `bson:"amount"`
	Sender string `bson:"sender"`
}

func (m *DocMsgBurnToken) GetType() string {
	return MsgTypeBurnToken
}

func (m *DocMsgBurnToken) BuildMsg(v interface{}) {
	msg := v.(*MsgBurnToken)

	m.Symbol = msg.Symbol
	m.Amount = fmt.Sprint(msg.Amount)
	m.Sender = msg.Sender
}

func (m *DocMsgBurnToken) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgBurnToken)

	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
