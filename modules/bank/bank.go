package bank

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type bankClient struct {
}

func NewClient() Client {
	return bankClient{}
}

func (bank bankClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgSend:
		docMsg := DocMsgSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgMultiSend:
		docMsg := DocMsgMultiSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
