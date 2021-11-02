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

func (slashing slashingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgUnjail:
		docMsg := DocTxMsgUnjail{}
		return docMsg.HandleTxMsg(msg), ok

	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
