package gov

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

// MsgDeposit
type DocTxMsgDeposit struct {
	ProposalID uint64        `bson:"proposal_id"` // ID of the proposal
	Depositor  string        `bson:"depositor"`   // Address of the depositor
	Amount     []models.Coin `bson:"amount"`      // Coins to add to the proposal's deposit
}

func (doctx *DocTxMsgDeposit) GetType() string {
	return MsgTypeDeposit
}

func (doctx *DocTxMsgDeposit) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgDeposit)
	doctx.Depositor = msg.Depositor
	doctx.Amount = models.BuildDocCoins(msg.Amount)
	doctx.ProposalID = msg.ProposalId
}

func (m *DocTxMsgDeposit) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgDeposit
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
