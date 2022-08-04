package mt

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type DocMsgMTMint struct {
	Id        string `bson:"id"` // need get from events
	DenomId   string `bson:"denom_id"`
	Amount    string `bson:"amount"`
	Data      string `bson:"data"`
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
}

func (m *DocMsgMTMint) GetType() string {
	return MsgTypeMintMT
}

func (m *DocMsgMTMint) BuildMsg(v interface{}) {
	msg := v.(*MsgMTMint)

	m.DenomId = strings.ToLower(msg.DenomId)
	m.Amount = fmt.Sprint(msg.Amount)
	m.Data = string(msg.Data)
	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
}

func (m *DocMsgMTMint) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMTMint)
	addrs = append(addrs, msg.Sender, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
