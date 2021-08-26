package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgStake struct {
	PoolName string      `bson:"pool_name"`
	Amount   models.Coin `bson:"amount"`
	Sender   string      `bson:"sender"`
}

func (m *DocTxMsgStake) GetType() string {
	return MsgTypeStake
}

func (m *DocTxMsgStake) BuildMsg(v interface{}) {
	msg := v.(*MsgStake)
	m.PoolName = msg.PoolName
	m.Amount = models.BuildDocCoin(msg.Amount)
	m.Sender = msg.Sender
}

func (m *DocTxMsgStake) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgStake
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
