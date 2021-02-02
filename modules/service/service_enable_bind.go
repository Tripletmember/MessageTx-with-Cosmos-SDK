package service

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type (
	DocMsgEnableServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgEnableServiceBinding) GetType() string {
	return MsgTypeEnableServiceBinding
}

func (m *DocMsgEnableServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgEnableServiceBinding)
	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = models.BuildDocCoins(msg.Deposit)
	m.Owner = msg.Owner
}

func (m *DocMsgEnableServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgEnableServiceBinding
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
