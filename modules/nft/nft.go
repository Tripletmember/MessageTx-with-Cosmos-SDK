package nft

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type nftClient struct {
}

func NewClient() Client {
	return nftClient{}
}

func (nft nftClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {

	switch msg := v.(type) {
	case *MsgNFTMint:
		docMsg := DocMsgNFTMint{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTEdit:
		docMsg := DocMsgNFTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTTransfer:
		docMsg := DocMsgNFTTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTBurn:
		docMsg := DocMsgNFTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgIssueDenom:
		docMsg := DocMsgIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgTransferDenom:
		docMsg := DocMsgTransferDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
