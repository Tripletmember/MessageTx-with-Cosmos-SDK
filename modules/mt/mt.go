package mt

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type mtClient struct {
}

func NewClient() Client {
	return mtClient{}
}

func (nft mtClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgMTMint:
		docMsg := DocMsgMTMint{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgMTEdit:
		docMsg := DocMsgMTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgMTTransfer:
		docMsg := DocMsgMTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgMTBurn:
		docMsg := DocMsgMTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgMTIssueDenom:
		docMsg := DocMsgMTIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgMTTransferDenom:
		docMsg := DocMsgMTTransferDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
