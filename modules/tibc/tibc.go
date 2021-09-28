package tibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type tibcClient struct {
}

func NewClient() Client {
	return tibcClient{}
}
func (ibc tibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgTIBCNftTransfer).Type():
		docMsg := DocMsgTIBCNftTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgTIBCUpdateClient).Type():
		docMsg := DocMsgTIBCUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgTIBCRecvPacket).Type():
		docMsg := DocMsgTIBCRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgTIBCAcknowledgement).Type():
		docMsg := DocMsgTIBCAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgCleanPacket).Type():
		docMsg := DocMsgCleanPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgRecvCleanPacket).Type():
		docMsg := DocMsgRecvCleanPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
