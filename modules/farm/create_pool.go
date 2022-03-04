package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgCreatePool struct {
	Description    string        `bson:"description"`
	LptDenom       string        `bson:"lpt_denom"`
	StartHeight    int64         `bson:"start_height"`
	RewardPerBlock []models.Coin `bson:"reward_per_block"`
	TotalReward    []models.Coin `bson:"total_reward"`
	Editable       bool          `bson:"editable"`
	Creator        string        `bson:"creator"`
}

func (m *DocTxMsgCreatePool) GetType() string {
	return MsgTypeCreatePool
}

func (m *DocTxMsgCreatePool) BuildMsg(v interface{}) {
	msg := v.(*MsgCreatePool)
	m.Description = msg.Description
	m.LptDenom = msg.LptDenom
	m.StartHeight = msg.StartHeight
	m.RewardPerBlock = models.BuildDocCoins(msg.RewardPerBlock)
	m.TotalReward = models.BuildDocCoins(msg.TotalReward)
	m.Editable = msg.Editable
	m.Creator = msg.Creator

}

func (m *DocTxMsgCreatePool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreatePool
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
