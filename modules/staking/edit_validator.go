package staking

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

// MsgEditValidator - struct for editing a validator
type DocMsgEditValidator struct {
	Description       Description `bson:"description"`
	ValidatorAddress  string      `bson:"validator_address"`
	CommissionRate    string      `bson:"commission_rate"`
	MinSelfDelegation string      `bson:"min_self_delegation"`
}

func (doctx *DocMsgEditValidator) GetType() string {
	return MsgTypeStakeEditValidator
}

func (doctx *DocMsgEditValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeEdit)
	doctx.ValidatorAddress = msg.ValidatorAddress
	commissionRate := msg.CommissionRate
	if commissionRate == nil {
		doctx.CommissionRate = ""
	} else {
		doctx.CommissionRate = commissionRate.String()
	}
	doctx.Description = loadDescription(msg.Description)
	if msg.MinSelfDelegation != nil {
		doctx.MinSelfDelegation = msg.MinSelfDelegation.String()
	}
}
func (m *DocMsgEditValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeEdit
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
