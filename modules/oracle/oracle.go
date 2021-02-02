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

func (oracle oracleClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {

	ok := true
	switch msg.Type() {
	case new(MsgStartFeed).Type():
		docMsg := DocMsgStartFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgPauseFeed).Type():
		docMsg := DocMsgPauseFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgEditFeed).Type():
		docMsg := DocMsgEditFeed{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgCreateFeed).Type():
		docMsg := DocMsgCreateFeed{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
