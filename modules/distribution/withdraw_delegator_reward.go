package distribution

import . "github.com/kaifei-bianjie/msg-parser/modules"

// msg struct for delegation withdraw from a single validator
type DocTxMsgWithdrawDelegatorReward struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
}

func (doctx *DocTxMsgWithdrawDelegatorReward) GetType() string {
	return MsgTypeWithdrawDelegatorReward
}

func (doctx *DocTxMsgWithdrawDelegatorReward) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawDelegatorReward)
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.ValidatorAddress = msg.ValidatorAddress
}
func (m *DocTxMsgWithdrawDelegatorReward) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgWithdrawDelegatorReward
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
