package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type feegrantClient struct {
}

func NewClient() Client {
	return feegrantClient{}
}

func (feegrant feegrantClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgGrantAllowance:
		docMsg := DocTxMsgGrantAllowance{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgRevokeAllowance:
		docMsg := DocTxMsgRevokeAllowance{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
