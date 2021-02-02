package distribution

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

// msg struct for delegation withdraw for all of the delegator's delegations
type DocTxMsgFundCommunityPool struct {
	Amount    []models.Coin `bson:"amount"`
	Depositor string        `bson:"depositor"`
}

func (doctx *DocTxMsgFundCommunityPool) GetType() string {
	return MsgTypeMsgFundCommunityPool
}

func (doctx *DocTxMsgFundCommunityPool) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgFundCommunityPool)
	doctx.Depositor = msg.Depositor
	doctx.Amount = models.BuildDocCoins(msg.Amount)
}
func (m *DocTxMsgFundCommunityPool) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgFundCommunityPool
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
