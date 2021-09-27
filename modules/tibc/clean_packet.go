package tibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgCleanPacket struct {
	Packet Packet `bson:"packet"`
	Signer string `bson:"signer"`
}

func (m *DocMsgCleanPacket) GetType() string {
	return MsgTypeTIBCCleanPacket
}

func (m *DocMsgCleanPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgCleanPacket)
	m.Signer = msg.Signer
	m.Packet = loadPacket(msg.Packet)
}

func (m *DocMsgCleanPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCleanPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
