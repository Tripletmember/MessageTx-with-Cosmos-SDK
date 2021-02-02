package gov

import (
	cdc "github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocTxMsgSubmitProposal struct {
	Proposer       string        `bson:"proposer"`        //  Address of the proposer
	InitialDeposit []models.Coin `bson:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Content        interface{}   `bson:"content"`
}

func (doctx *DocTxMsgSubmitProposal) GetType() string {
	return MsgTypeSubmitProposal
}

func (doctx *DocTxMsgSubmitProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSubmitProposal)
	doctx.Content = CovertContent(msg.GetContent())
	doctx.Proposer = msg.Proposer
	doctx.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
}

func CovertContent(content GovContent) interface{} {
	switch content.ProposalType() {
	case ProposalTypeCancelSoftwareUpgrade:
		var data ContentCancelSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeSoftwareUpgrade:
		var data ContentSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeCommunityPoolSpend:
		var data ContentCommunityPoolSpendProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeClientUpdate:
		var data ContentClientUpdateProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeText:
		var data ContentTextProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeParameterChange:
		var data ContentParameterChangeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	}
	return content
}

func (m *DocTxMsgSubmitProposal) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgSubmitProposal
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)
	content := msg.GetContent()
	if content != nil && ProposalTypeCommunityPoolSpend == content.ProposalType() {
		communityPoolSpend := CovertContent(content).(ContentCommunityPoolSpendProposal)
		addrs = append(addrs, communityPoolSpend.Recipient)
	}
	addrs = append(addrs, msg.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
