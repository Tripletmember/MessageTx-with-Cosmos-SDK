package coinswap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgSwapOrder struct {
	Input      Input  `bson:"input"`        // the amount the sender is trading
	Output     Output `bson:"output"`       // the amount the sender is receiving
	Deadline   int64  `bson:"deadline"`     // deadline for the transaction to still be considered valid
	IsBuyOrder bool   `bson:"is_buy_order"` // boolean indicating whether the order should be treated as a buy or sell
}

type Input struct {
	Address string      `bson:"address"`
	Coin    models.Coin `bson:"coin"`
}

type Output struct {
	Address string      `bson:"address"`
	Coin    models.Coin `bson:"coin"`
}

func (doctx *DocTxMsgSwapOrder) GetType() string {
	return MsgTypeSwapOrder
}

func (doctx *DocTxMsgSwapOrder) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSwapOrder)
	doctx.Deadline = msg.Deadline
	doctx.IsBuyOrder = msg.IsBuyOrder
	doctx.Input = Input{
		Address: msg.Input.Address,
		Coin:    models.BuildDocCoin(msg.Input.Coin),
	}
	doctx.Output = Output{
		Address: msg.Output.Address,
		Coin:    models.BuildDocCoin(msg.Output.Coin),
	}
}
func (m *DocTxMsgSwapOrder) HandleTxMsg(v sdk.Msg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgSwapOrder
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Output.Address, msg.Input.Address)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
