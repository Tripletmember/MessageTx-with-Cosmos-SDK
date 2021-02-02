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
	switch v.Type() {
	case new(MsgSend).Type():
		docMsg := DocMsgSend{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgMultiSend).Type():
		docMsg := DocMsgMultiSend{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
