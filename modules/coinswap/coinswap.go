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

func (coinswap coinswapClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgAddLiquidity:
		docMsg := DocTxMsgAddLiquidity{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgRemoveLiquidity:
		docMsg := DocTxMsgRemoveLiquidity{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgSwapOrder:
		docMsg := DocTxMsgSwapOrder{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
