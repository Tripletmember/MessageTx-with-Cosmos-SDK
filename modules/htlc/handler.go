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

func (htlc htlcClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgClaimHTLC:
		docMsg := DocTxMsgClaimHTLC{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgCreateHTLC:
		docMsg := DocTxMsgCreateHTLC{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
