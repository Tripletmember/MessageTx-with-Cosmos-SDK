package coinswap

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type coinswapClient struct {
}

func NewClient() Client {
	return coinswapClient{}
}

func (coinswap coinswapClient) HandleTxMsg(msg types.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgAddLiquidity).Type():
		docMsg := DocTxMsgAddLiquidity{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgRemoveLiquidity).Type():
		docMsg := DocTxMsgRemoveLiquidity{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgSwapOrder).Type():
		docMsg := DocTxMsgSwapOrder{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
