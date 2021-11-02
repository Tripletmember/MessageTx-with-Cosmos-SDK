package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type distributionClient struct {
}

func NewClient() Client {
	return distributionClient{}
}

func (distribution distributionClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgStakeSetWithdrawAddress:
		docMsg := DocTxMsgSetWithdrawAddress{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgWithdrawDelegatorReward:
		docMsg := DocTxMsgWithdrawDelegatorReward{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgWithdrawValidatorCommission:
		docMsg := DocTxMsgWithdrawValidatorCommission{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgFundCommunityPool:
		docMsg := DocTxMsgFundCommunityPool{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
