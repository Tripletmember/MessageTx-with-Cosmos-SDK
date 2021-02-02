package record

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type recordClient struct {
}

func NewClient() Client {
	return recordClient{}
}

func (record recordClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecordCreate).Type():
		docMsg := DocMsgRecordCreate{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
