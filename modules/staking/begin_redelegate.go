package staking

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

// MsgDelegate - struct for bonding transactions
type DocTxMsgBeginRedelegate struct {
	DelegatorAddress    string      `bson:"delegator_address"`
	ValidatorSrcAddress string      `bson:"validator_src_address"`
	ValidatorDstAddress string      `bson:"validator_dst_address"`
	Amount              models.Coin `bson:"amount"`
}

func (doctx *DocTxMsgBeginRedelegate) GetType() string {
	return MsgTypeBeginRedelegate
}

func (doctx *DocTxMsgBeginRedelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgBeginRedelegate)
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.ValidatorSrcAddress = msg.ValidatorSrcAddress
	doctx.ValidatorDstAddress = msg.ValidatorDstAddress
	doctx.Amount = models.BuildDocCoin(msg.Amount)
}
func (m *DocTxMsgBeginRedelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgBeginRedelegate
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorDstAddress, msg.ValidatorSrcAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
