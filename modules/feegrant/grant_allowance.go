package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/golang/protobuf/proto"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgGrantAllowance struct {
	Granter   string     `bson:"granter"`
	Grantee   string     `bson:"grantee"`
	Allowance *Allowance `bson:"allowance"`
}

type Allowance struct {
	SpendLimit []models.Coin `bson:"spend_limit"`
	Expiration int64         `bson:"expiration"`
}

func (m *DocTxMsgGrantAllowance) GetType() string {
	return MsgTypeGrantAllowance
}

func (m *DocTxMsgGrantAllowance) BuildMsg(v interface{}) {
	msg := v.(*MsgGrantAllowance)
	m.Granter = msg.Granter
	m.Grantee = msg.Grantee

	if msg.Allowance.Value != nil {
		var cc feegrant.BasicAllowance
		_ = proto.Unmarshal(msg.Allowance.Value, &cc)

		m.Allowance = &Allowance{
			SpendLimit: models.BuildDocCoins(cc.SpendLimit),
			Expiration: cc.Expiration.Unix(),
		}
	}
}

func (m *DocTxMsgGrantAllowance) HandleTxMsg(v sdk.Msg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgGrantAllowance)
	addrs = append(addrs, msg.Granter, msg.Grantee)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
