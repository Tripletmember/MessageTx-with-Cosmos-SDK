package service

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type serviceClient struct {
}

func NewClient() Client {
	return serviceClient{}
}

func (service serviceClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgDefineService:
		docMsg := DocMsgDefineService{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgBindService:
		docMsg := DocMsgBindService{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgCallService:
		docMsg := DocMsgCallService{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgRespondService:
		docMsg := DocMsgServiceResponse{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgUpdateServiceBinding:
		docMsg := DocMsgUpdateServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgSetWithdrawAddress:
		docMsg := DocMsgSetWithdrawAddress{}
		msgData := MsgSetWithdrawAddress{}
		ConvertMsg(msg, &msgData)
		if msgData.Owner == "" {
			return MsgDocInfo{}, false
		}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgDisableServiceBinding:
		docMsg := DocMsgDisableServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgEnableServiceBinding:
		docMsg := DocMsgEnableServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgRefundServiceDeposit:
		docMsg := DocMsgRefundServiceDeposit{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgUpdateRequestContext:
		docMsg := DocMsgUpdateRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgPauseRequestContext:
		docMsg := DocMsgPauseRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgStartRequestContext:
		docMsg := DocMsgStartRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgKillRequestContext:
		docMsg := DocMsgKillRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case *MsgWithdrawEarnedFees:
		docMsg := DocMsgWithdrawEarnedFees{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
