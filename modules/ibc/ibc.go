package ibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type ibcClient struct {
}

func NewClient() Client {
	return ibcClient{}
}

func (ibc ibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecvPacket).Type():
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTransfer).Type():
		docMsg := DocMsgTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgCreateClient).Type():
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpdateClient).Type():
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpgradeClient).Type():
		docMsg := DocMsgUpgradeClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgSubmitMisbehaviour).Type():
		docMsg := DocMsgSubmitMisbehaviour{}
		msgDocInfo = docMsg.HandleTxMsg(v)

	case new(MsgConnectionOpenInit).Type():
		docMsg := DocMsgConnectionOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenTry).Type():
		docMsg := DocMsgConnectionOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenAck).Type():
		docMsg := DocMsgConnectionOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenConfirm).Type():
		docMsg := DocMsgConnectionOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)

	case new(MsgChannelOpenInit).Type():
		docMsg := DocMsgChannelOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenTry).Type():
		docMsg := DocMsgChannelOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenAck).Type():
		docMsg := DocMsgChannelOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenConfirm).Type():
		docMsg := DocMsgChannelOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelCloseInit).Type():
		docMsg := DocMsgChannelCloseInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelCloseConfirm).Type():
		docMsg := DocMsgChannelCloseConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTimeout).Type():
		docMsg := DocMsgTimeout{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTimeoutOnClose).Type():
		docMsg := DocMsgTimeoutOnClose{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgAcknowledgement).Type():
		docMsg := DocMsgAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	default:
		ok = false
	}
	return msgDocInfo, ok
}
