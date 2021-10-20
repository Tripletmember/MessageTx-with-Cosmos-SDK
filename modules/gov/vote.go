package gov

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

// MsgVote
type DocTxMsgVote struct {
	ProposalID int64  `bson:"proposal_id"` // ID of the proposal
	Voter      string `bson:"voter"`       //  address of the voter
	Option     int32  `bson:"option"`      //  option from OptionSet chosen by the voter
}

func (doctx *DocTxMsgVote) GetType() string {
	return MsgTypeVote
}

func (doctx *DocTxMsgVote) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgVote)
	doctx.Voter = msg.Voter
	doctx.Option = int32(msg.Option)
	doctx.ProposalID = int64(msg.ProposalId)
}

func (m *DocTxMsgVote) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgVote
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
