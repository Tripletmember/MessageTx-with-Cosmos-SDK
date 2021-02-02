package upgrade

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type upgradeClient struct {
}

func NewClient() Client {
	return upgradeClient{}
}

func (upgrade upgradeClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
