package mt

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type (
	DocMsgMTransfer struct {
		Id        string `bson:"id"`
		DenomId   string `bson:"denom_id"`
		Amount    uint64 `bson:"amount"`
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
	}
)

func (m *DocMsgMTransfer) GetType() string {
	return MsgTypeTransferMT
}

func (m *DocMsgMTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgMTTransfer)

	m.Id = strings.ToLower(msg.Id)
	m.DenomId = strings.ToLower(msg.DenomId)
	m.Amount = msg.Amount
	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
}

func (m *DocMsgMTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMTTransfer)
	addrs = append(addrs, msg.Sender, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
