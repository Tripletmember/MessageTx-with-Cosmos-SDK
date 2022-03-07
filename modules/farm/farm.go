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

func (farm farmClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgCreatePool:
		docMsg := DocTxMsgCreatePool{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgCreatePoolWithCommunityPool:
		docMsg := DocTxMsgCreatePoolWithCommunityPool{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgDestroyPool:
		docMsg := DocTxMsgDestroyPool{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgAdjustPool:
		docMsg := DocTxMsgAdjustPool{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgHarvest:
		docMsg := DocTxMsgHarvest{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgStake:
		docMsg := DocTxMsgStake{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgUnstake:
		docMsg := DocTxMsgUnstake{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
