package mt

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type DocMsgMTTransferDenom struct {
	Id        string `bson:"id"`
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
}

func (m *DocMsgMTTransferDenom) GetType() string {
	return MsgTypeMTTransferDenom
}

func (m *DocMsgMTTransferDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgMTTransferDenom)

	m.Id = strings.ToLower(msg.Id)
	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
}

func (m *DocMsgMTTransferDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgMTTransferDenom)

	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
