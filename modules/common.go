package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	models "github.com/kaifei-bianjie/msg-parser/types"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func CreateMsgDocInfo(msg sdk.Msg, handler func() (Msg, []string)) MsgDocInfo {
	var (
		docTxMsg models.TxMsg
		signers  []string
		addrs    []string
	)

	m, addrcollections := handler()

	m.BuildMsg(msg)
	docTxMsg = models.TxMsg{
		Type: m.GetType(),
		Msg:  m,
	}

	_, signers = models.BuildDocSigners(msg.GetSigners())
	addrs = append(addrs, signers...)
	addrs = append(addrs, addrcollections...)

	return MsgDocInfo{
		DocTxMsg: docTxMsg,
		Signers:  signers,
		Addrs:    addrs,
	}
}

func ConvertMsg(v interface{}, msg interface{}) {
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
}

func ConvertAny(v *types.Any) string {
	if v != nil {
		return utils.MarshalJsonIgnoreErr(v)
	}
	return ""
}
