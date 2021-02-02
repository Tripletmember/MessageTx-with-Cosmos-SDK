package staking

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

// MsgDelegate - struct for bonding transactions
type DocTxMsgDelegate struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
	Amount           Coin   `bson:"amount"`
}

func (doctx *DocTxMsgDelegate) GetType() string {
	return MsgTypeStakeDelegate
}

func (doctx *DocTxMsgDelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeDelegate)
	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.Amount = Coin(models.BuildDocCoin(msg.Amount))
}
func (m *DocTxMsgDelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeDelegate
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
