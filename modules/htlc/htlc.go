package htlc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgCreateHTLC struct {
	Sender               string        `bson:"sender"`                  // the initiator address
	To                   string        `bson:"to"`                      // the destination address
	ReceiverOnOtherChain string        `bson:"receiver_on_other_chain"` // the claim receiving address on the other chain
	SenderOnOtherChain   string        `bson:"sender_on_other_chain"`
	Amount               []models.Coin `bson:"amount"`    // the amount to be transferred
	HashLock             string        `bson:"hash_lock"` // the hash lock generated from secret (and timestamp if provided)
	Timestamp            int64         `bson:"timestamp"` // if provided, used to generate the hash lock together with secret
	TimeLock             int64         `bson:"time_lock"` // the time span after which the HTLC will expire
	Transfer             bool          `bson:"transfer"`
}

func (doctx *DocTxMsgCreateHTLC) GetType() string {
	return MsgTypeCreateHTLC
}

func (doctx *DocTxMsgCreateHTLC) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCreateHTLC)
	doctx.Sender = msg.Sender
	doctx.To = msg.To
	doctx.Amount = models.BuildDocCoins(msg.Amount)
	doctx.Timestamp = int64(msg.Timestamp)
	doctx.HashLock = msg.HashLock
	doctx.TimeLock = int64(msg.TimeLock)
	doctx.ReceiverOnOtherChain = msg.ReceiverOnOtherChain
	doctx.SenderOnOtherChain = msg.SenderOnOtherChain
	doctx.Transfer = msg.Transfer
}

func (m *DocTxMsgCreateHTLC) HandleTxMsg(v sdk.Msg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgCreateHTLC
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender, msg.To)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

type DocTxMsgClaimHTLC struct {
	Sender string `bson:"sender"` // the initiator address
	Id     string `bson:"id"`     // the id identifying the HTLC to be claimed
	Secret string `bson:"secret"` // the secret with which to claim
}

func (doctx *DocTxMsgClaimHTLC) GetType() string {
	return MsgTypeClaimHTLC
}

func (doctx *DocTxMsgClaimHTLC) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgClaimHTLC)
	doctx.Sender = msg.Sender
	doctx.Secret = msg.Secret
	doctx.Id = msg.Id
}

func (m *DocTxMsgClaimHTLC) HandleTxMsg(v sdk.Msg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgClaimHTLC
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
