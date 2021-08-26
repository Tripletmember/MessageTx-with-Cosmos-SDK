package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocTxMsgDestroyPool struct {
	PoolName string `bson:"pool_name"`
	Creator  string `bson:"creator"`
}

func (m *DocTxMsgDestroyPool) GetType() string {
	return MsgTypeDestroyPool
}

func (m *DocTxMsgDestroyPool) BuildMsg(v interface{}) {
	msg := v.(*MsgDestroyPool)
	m.PoolName = msg.PoolName
	m.Creator = msg.Creator
}

func (m *DocTxMsgDestroyPool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgDestroyPool
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
