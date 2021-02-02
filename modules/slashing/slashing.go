package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type slashingClient struct {
}

func NewClient() Client {
	return slashingClient{}
}

func (slashing slashingClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgUnjail).Type():
		docMsg := DocTxMsgUnjail{}
		return docMsg.HandleTxMsg(msg), ok

	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
