package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type govClient struct {
}

func NewClient() Client {
	return govClient{}
}

func (gov govClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgSubmitProposal:
		docMsg := DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgVote:
		docMsg := DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgDeposit:
		docMsg := DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgVoteWeighted:
		docMsg := DocTxMsgVoteWeighted{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
