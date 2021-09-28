package tibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgRecvCleanPacket struct {
	CleanPacket     CleanPacket `bson:"clean_packet"`
	ProofCommitment string      `bson:"proof_commitment"`
	ProofHeight     Height      `bson:"proof_height"`
	Signer          string      `bson:"signer"`
}

func (m *DocMsgRecvCleanPacket) GetType() string {
	return MsgTypeTIBCRecvCleanPacket
}

func (m *DocMsgRecvCleanPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvCleanPacket)
	m.Signer = msg.Signer
	m.CleanPacket = loadCleanPacket(msg.CleanPacket)
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(msg.ProofCommitment)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgRecvCleanPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvCleanPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
