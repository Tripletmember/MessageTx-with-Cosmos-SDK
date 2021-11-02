package token

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type tokenClient struct {
}

func NewClient() Client {
	return tokenClient{}
}
func (token tokenClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgMintToken:
		docMsg := DocMsgMintToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnToken:
		docMsg := DocMsgBurnToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditToken:
		docMsg := DocMsgEditToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueToken:
		docMsg := DocMsgIssueToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwner:
		docMsg := DocMsgTransferTokenOwner{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
