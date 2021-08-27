package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgUnstake struct {
	PoolName string      `bson:"pool_name"`
	Amount   models.Coin `bson:"amount"`
	Sender   string      `bson:"sender"`
}

func (m *DocTxMsgUnstake) GetType() string {
	return MsgTypeUnstake
}

func (m *DocTxMsgUnstake) BuildMsg(v interface{}) {
	msg := v.(*MsgUnstake)
	m.PoolName = msg.PoolName
	m.Amount = models.BuildDocCoin(msg.Amount)
	m.Sender = msg.Sender
}

func (m *DocTxMsgUnstake) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgUnstake
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
