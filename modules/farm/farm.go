package farm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type farmClient struct {
}

func NewClient() Client {
	return farmClient{}
}

func (farm farmClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgCreatePool).Type():
		docMsg := DocTxMsgCreatePool{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgDestroyPool).Type():
		docMsg := DocTxMsgDestroyPool{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgAdjustPool).Type():
		docMsg := DocTxMsgAdjustPool{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgHarvest).Type():
		docMsg := DocTxMsgHarvest{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgStake).Type():
		docMsg := DocTxMsgStake{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgUnstake).Type():
		docMsg := DocTxMsgUnstake{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
