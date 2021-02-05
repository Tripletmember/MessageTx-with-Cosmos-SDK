package bank

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type (
	DocMsgSend struct {
		FromAddress string        `bson:"from_address"`
		ToAddress   string        `bson:"to_address"`
		Amount      []models.Coin `bson:"amount"`
	}
)

func (m *DocMsgSend) GetType() string {
	return MsgTypeSend
}

func (m *DocMsgSend) BuildMsg(v interface{}) {
	msg := v.(*MsgSend)
	m.FromAddress = msg.FromAddress
	m.ToAddress = msg.ToAddress
	m.Amount = models.BuildDocCoins(msg.Amount)
}

func (m *DocMsgSend) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgSend
	)
	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.FromAddress, msg.ToAddress)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
