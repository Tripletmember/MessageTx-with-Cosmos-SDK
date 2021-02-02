package wasm

//
//import (
//	"github.com/cosmos/cosmos-sdk/types"
//	. "github.com/kaifei-bianjie/msg-parser/modules"
//)
//
//type wasmClient struct {
//}
//
//func NewClient() Client {
//	return wasmClient{}
//}
//
//func (wasm wasmClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
//	var (
//		msgDocInfo MsgDocInfo
//	)
//	ok := true
//	switch v.Type() {
//	case new(MsgStoreCode).Type():
//		docMsg := DocMsgStoreCode{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	case new(MsgInstantiateContract).Type():
//		docMsg := DocMsgInstantiateContract{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	case new(MsgExecuteContract).Type():
//		docMsg := DocMsgExecuteContract{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	case new(MsgMigrateContract).Type():
//		docMsg := DocMsgMigrateContract{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	case new(MsgUpdateAdmin).Type():
//		docMsg := DocMsgUpdateAdmin{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	case new(MsgClearAdmin).Type():
//		docMsg := DocMsgClearAdmin{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		break
//	default:
//		ok = false
//	}
//	return msgDocInfo, ok
//}
