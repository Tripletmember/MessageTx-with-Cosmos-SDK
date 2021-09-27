package tibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgRecvCleanPacket struct {
	Packet Packet `bson:"packet"`
	Signer string `bson:"signer"`
}

func (m *DocMsgRecvCleanPacket) GetType() string {
	return MsgTypeTIBCRecvCleanPacket
}

func (m *DocMsgRecvCleanPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvCleanPacket)
	m.Signer = msg.Signer
	m.Packet = loadPacket(msg.Packet)
}

func (m *DocMsgRecvCleanPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvCleanPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
