package staking

import (
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
	cdc "github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

// MsgCreateValidator defines an SDK message for creating a new validator.
type DocTxMsgCreateValidator struct {
	Description       Description     `bson:"description"`
	Commission        CommissionRates `bson:"commission"`
	MinSelfDelegation string          `bson:"min_self_delegation"`
	DelegatorAddress  string          `bson:"delegator_address"`
	ValidatorAddress  string          `bson:"validator_address"`
	Pubkey            string          `bson:"pubkey"`
	Value             Coin            `bson:"value"`
}

func (doctx *DocTxMsgCreateValidator) GetType() string {
	return MsgTypeStakeCreateValidator
}

func (doctx *DocTxMsgCreateValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeCreate)
	if msg.Pubkey != nil {
		doctx.Pubkey = utils.MarshalJsonIgnoreErr(msg.Pubkey)
	}

	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.MinSelfDelegation = msg.MinSelfDelegation.String()
	doctx.Commission = CommissionRates{
		Rate:          msg.Commission.Rate.String(),
		MaxChangeRate: msg.Commission.MaxChangeRate.String(),
		MaxRate:       msg.Commission.MaxRate.String(),
	}
	doctx.Description = loadDescription(msg.Description)
	doctx.Value = Coin{Denom: msg.Value.Denom, Amount: msg.Value.Amount.String()}
}
func (m *DocTxMsgCreateValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeCreate
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

func loadDescription(description stake.Description) Description {
	return Description{
		Moniker:         description.Moniker,
		Details:         description.Details,
		Identity:        description.Identity,
		Website:         description.Website,
		SecurityContact: description.SecurityContact,
	}
}
