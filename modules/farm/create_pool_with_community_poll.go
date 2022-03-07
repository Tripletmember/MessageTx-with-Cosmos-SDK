package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgCreatePoolWithCommunityPool struct {
	Content        Content       `bson:"content"`
	InitialDeposit []models.Coin `bson:"initial_deposit"`
	Proposer       string        `bson:"proposer"`
}

// CommunityPoolCreateFarmProposal is a gov Content type for creating a farm pool with community pool
type Content struct {
	Title           string        `bson:"title"`
	Description     string        `bson:"description"`
	PoolDescription string        `bson:"pool_description"`
	LptDenom        string        `bson:"lpt_denom"`
	RewardPerBlock  []models.Coin `bson:"reward_per_block"`
	FundApplied     []models.Coin `bson:"fund_applied"`
	FundSelfBond    []models.Coin `bson:"fund_self_bond"`
}

func (m *DocTxMsgCreatePoolWithCommunityPool) GetType() string {
	return MsgTypeCreateProposal
}

func (m *DocTxMsgCreatePoolWithCommunityPool) BuildMsg(v interface{}) {
	msg := v.(*MsgCreatePoolWithCommunityPool)
	m.Content = loadContent(msg.Content)
	m.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
	m.Proposer = msg.Proposer

}

func (m *DocTxMsgCreatePoolWithCommunityPool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreatePoolWithCommunityPool
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

func loadContent(content CommunityPoolCreateFarmProposal) Content {
	return Content{
		Title:           content.Title,
		Description:     content.Description,
		PoolDescription: content.PoolDescription,
		LptDenom:        content.LptDenom,
		RewardPerBlock:  models.BuildDocCoins(content.RewardPerBlock),
		FundApplied:     models.BuildDocCoins(content.FundApplied),
		FundSelfBond:    models.BuildDocCoins(content.FundSelfBond),
	}
}
