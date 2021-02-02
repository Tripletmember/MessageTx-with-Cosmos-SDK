package distribution

import . "github.com/kaifei-bianjie/msg-parser/modules"

// msg struct for validator withdraw
type DocTxMsgWithdrawValidatorCommission struct {
	ValidatorAddress string `bson:"validator_address"`
}

func (doctx *DocTxMsgWithdrawValidatorCommission) GetType() string {
	return MsgTypeMsgWithdrawValidatorCommission
}

func (doctx *DocTxMsgWithdrawValidatorCommission) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawValidatorCommission)
	doctx.ValidatorAddress = msg.ValidatorAddress
}

func (m *DocTxMsgWithdrawValidatorCommission) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgWithdrawValidatorCommission
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
