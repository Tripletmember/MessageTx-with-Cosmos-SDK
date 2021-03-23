package token

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocMsgEditToken struct {
	Symbol    string `bson:"symbol"`
	Name      string `bson:"name"`
	MaxSupply string `bson:"max_supply"`
	Mintable  bool   `bson:"mintable"`
	Owner     string `bson:"owner"`
}

func (m *DocMsgEditToken) GetType() string {
	return MsgTypeEditToken
}

func (m *DocMsgEditToken) BuildMsg(v interface{}) {
	msg := v.(*MsgEditToken)

	m.Symbol = msg.Symbol
	m.Owner = msg.Owner
	m.Name = msg.Name
	m.MaxSupply = fmt.Sprint(msg.MaxSupply)
	m.Mintable = msg.Mintable.ToBool()
}

func (m *DocMsgEditToken) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgEditToken
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
