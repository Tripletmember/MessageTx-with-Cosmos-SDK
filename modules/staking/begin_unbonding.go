package staking

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

// MsgBeginUnbonding - struct for unbonding transactions
type DocTxMsgBeginUnbonding struct {
	DelegatorAddress string      `bson:"delegator_address"`
	ValidatorAddress string      `bson:"validator_address"`
	Amount           models.Coin `bson:"amount"`
}

func (doctx *DocTxMsgBeginUnbonding) GetType() string {
	return MsgTypeStakeBeginUnbonding
}

func (doctx *DocTxMsgBeginUnbonding) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeBeginUnbonding)
	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.Amount = models.BuildDocCoin(msg.Amount)
}
func (m *DocTxMsgBeginUnbonding) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeBeginUnbonding
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
