package htlc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type htlcClient struct {
}

func NewClient() Client {
	return htlcClient{}
}

func (htlc htlcClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgClaimHTLC).Type():
		docMsg := DocTxMsgClaimHTLC{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgCreateHTLC).Type():
		docMsg := DocTxMsgCreateHTLC{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
