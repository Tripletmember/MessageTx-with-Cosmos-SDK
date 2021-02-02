package coinswap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgRemoveLiquidity struct {
	MinToken          string      `bson:"min_token"`          // coin to be withdrawn with a lower bound for its amount
	WithdrawLiquidity models.Coin `bson:"withdraw_liquidity"` // amount of UNI to be burned to withdraw liquidity from a reserve pool
	MinIrisAmt        string      `bson:"min_iris_amt"`       // minimum amount of the native asset the sender is willing to accept
	Deadline          int64       `bson:"deadline"`
	Sender            string      `bson:"sender"`
}

func (doctx *DocTxMsgRemoveLiquidity) GetType() string {
	return MsgTypeRemoveLiquidity
}

func (doctx *DocTxMsgRemoveLiquidity) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgRemoveLiquidity)
	doctx.Sender = msg.Sender
	doctx.MinIrisAmt = msg.MinStandardAmt.String()
	doctx.MinToken = msg.MinToken.String()
	doctx.Deadline = msg.Deadline
	doctx.WithdrawLiquidity = models.BuildDocCoin(msg.WithdrawLiquidity)
}
func (m *DocTxMsgRemoveLiquidity) HandleTxMsg(v sdk.Msg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgRemoveLiquidity
	)
	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
