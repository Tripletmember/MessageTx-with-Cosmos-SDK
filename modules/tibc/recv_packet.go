package tibc

import (
	cdc "github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgTIBCRecvPacket struct {
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgTIBCRecvPacket) GetType() string {
	return MsgTypeTIBCRecvPacket
}

func (m *DocMsgTIBCRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgTIBCRecvPacket)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(msg.ProofCommitment)
	m.Packet = loadPacket(msg.Packet)
	//m.PacketId = fmt.Sprintf("%v%v%v", msg.Packet.SourceChain,
	//	msg.Packet.DestinationChain, msg.Packet.Sequence)

}

func (m *DocMsgTIBCRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTIBCRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
func UnmarshalPacketData(bytesdata []byte) PacketData {
	var (
		packetData NonFungibleTokenPacketData
	)
	cdc.GetMarshaler().MustUnmarshalBinaryBare(bytesdata, &packetData)
	return PacketData{
		Class:          packetData.Class,
		Receiver:       packetData.Receiver,
		Sender:         packetData.Sender,
		Id:             packetData.Id,
		Uri:            packetData.Uri,
		AwayFromOrigin: packetData.AwayFromOrigin,
	}
}
