package random

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type randomClient struct {
}

func NewClient() Client {
	return randomClient{}
}

func (random randomClient) HandleTxMsg(msgData sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msgData.Type() {
	case new(MsgRequestRandom).Type():

		txMsg := DocTxMsgRequestRand{}
		return txMsg.HandleTxMsg(msgData), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
