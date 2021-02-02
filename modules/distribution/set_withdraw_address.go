package distribution

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocTxMsgSetWithdrawAddress struct {
	DelegatorAddress string `bson:"delegator_address"`
	WithdrawAddress  string `bson:"withdraw_address"`
}

func (doctx *DocTxMsgSetWithdrawAddress) GetType() string {
	return MsgTypeSetWithdrawAddress
}

func (doctx *DocTxMsgSetWithdrawAddress) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeSetWithdrawAddress)
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.WithdrawAddress = msg.WithdrawAddress
}

func (m *DocTxMsgSetWithdrawAddress) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeSetWithdrawAddress
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.WithdrawAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
