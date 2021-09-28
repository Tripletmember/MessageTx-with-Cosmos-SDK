package tibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgCleanPacket struct {
	CleanPacket CleanPacket `bson:"clean_packet"`
	Signer      string      `bson:"signer"`
}

func (m *DocMsgCleanPacket) GetType() string {
	return MsgTypeTIBCCleanPacket
}

func (m *DocMsgCleanPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgCleanPacket)
	m.Signer = msg.Signer
	m.CleanPacket = loadCleanPacket(msg.CleanPacket)
}

func (m *DocMsgCleanPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCleanPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
