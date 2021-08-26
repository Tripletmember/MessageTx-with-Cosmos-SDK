package farm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgAdjustPool struct {
	PoolName         string        `bson:"pool_name"`
	AdditionalReward []models.Coin `bson:"additional_reward"`
	RewardPerBlock   []models.Coin `bson:"reward_per_block"`
	Creator          string        `bson:"creator"`
}

func (m *DocTxMsgAdjustPool) GetType() string {
	return MsgTypeAdjustPool
}

func (m *DocTxMsgAdjustPool) BuildMsg(v interface{}) {
	msg := v.(*MsgAdjustPool)
	m.PoolName = msg.PoolName
	m.RewardPerBlock = models.BuildDocCoins(msg.RewardPerBlock)
	m.AdditionalReward = models.BuildDocCoins(msg.AdditionalReward)
	m.Creator = msg.Creator

}

func (m *DocTxMsgAdjustPool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgAdjustPool
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
