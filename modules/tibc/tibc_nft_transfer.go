package tibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgTIBCNftTransfer struct {
	// the class to which the NFT to be transferred belongs
	Class        string `bson:"class"`
	Id           string `bson:"id"`
	Sender       string `bson:"sender"`
	Receiver     string `bson:"receiver"`
	DestChain    string `bson:"dest_chain"`
	RealayChain  string `bson:"realay_chain"`
	DestContract string `bson:"dest_contract"`
}

func (m *DocMsgTIBCNftTransfer) GetType() string {
	return MsgTypeTIBCNftTransfer
}

func (m *DocMsgTIBCNftTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgTIBCNftTransfer)
	m.Class = msg.Class
	m.Id = msg.Id
	m.Sender = msg.Sender
	m.Receiver = msg.Receiver
	m.DestChain = msg.DestChain
	m.RealayChain = msg.RealayChain
	m.DestContract = msg.DestContract
}

func (m *DocMsgTIBCNftTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTIBCNftTransfer
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
