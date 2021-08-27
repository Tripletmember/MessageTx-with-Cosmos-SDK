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

func (nft nftClient) HandleTxMsg(msg types.Msg) (MsgDocInfo, bool) {

	switch msg.Type() {
	case new(MsgNFTMint).Type():
		docMsg := DocMsgNFTMint{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTEdit).Type():
		docMsg := DocMsgNFTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTTransfer).Type():
		docMsg := DocMsgNFTTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTBurn).Type():
		docMsg := DocMsgNFTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgIssueDenom).Type():
		docMsg := DocMsgIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgTransferDenom).Type():
		docMsg := DocMsgTransferDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
