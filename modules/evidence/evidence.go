package evidence

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type evidenceClient struct {
}

func NewClient() Client {
	return evidenceClient{}
}

func (evidence evidenceClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgSubmitEvidence:
		docMsg := DocMsgSubmitEvidence{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
