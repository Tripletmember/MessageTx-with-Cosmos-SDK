package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type stakingClient struct {
}

func NewClient() Client {
	return stakingClient{}
}

func (staking stakingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgBeginRedelegate:
		docMsg := DocTxMsgBeginRedelegate{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgStakeBeginUnbonding:
		docMsg := DocTxMsgBeginUnbonding{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgStakeDelegate:
		docMsg := DocTxMsgDelegate{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgStakeEdit:
		docMsg := DocMsgEditValidator{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgStakeCreate:
		docMsg := DocTxMsgCreateValidator{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
