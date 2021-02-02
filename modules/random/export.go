package random

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
