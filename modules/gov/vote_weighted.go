package gov

import . "github.com/kaifei-bianjie/msg-parser/modules"

// MsgVote
type DocTxMsgVoteWeighted struct {
	ProposalID int64                `bson:"proposal_id"` // ID of the proposal
	Voter      string               `bson:"voter"`       //  address of the voter
	Options    []WeightedVoteOption `bson:"options"`     //  option from OptionSet chosen by the voter
}

type WeightedVoteOption struct {
	Option int32  `bson:"option"`
	Weight string `bson:"weight"`
}

func (m *DocTxMsgVoteWeighted) GetType() string {
	return MsgTypeVoteWeighted
}

func (m *DocTxMsgVoteWeighted) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgVoteWeighted)
	m.Voter = msg.Voter

	for _, v := range msg.Options {
		m.Options = append(m.Options, WeightedVoteOption{
			Option: int32(v.Option),
			Weight: v.Weight.String(),
		})
	}
	m.ProposalID = int64(msg.ProposalId)
}

func (m *DocTxMsgVoteWeighted) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgVoteWeighted)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
