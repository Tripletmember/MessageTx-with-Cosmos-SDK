package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type oracleClient struct {
}

func NewClient() Client {
	return oracleClient{}
}

func (oracle oracleClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {

	ok := true
	switch msg := v.(type) {
	case *MsgStartFeed:
		docMsg := DocMsgStartFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgPauseFeed:
		docMsg := DocMsgPauseFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgEditFeed:
		docMsg := DocMsgEditFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgCreateFeed:
		docMsg := DocMsgCreateFeed{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
