package service

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"strings"
)

type (
	DocMsgKillRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgKillRequestContext) GetType() string {
	return MsgTypeKillRequestContext
}

func (m *DocMsgKillRequestContext) BuildMsg(v interface{}) {
	msg := v.(*MsgKillRequestContext)

	m.RequestContextID = strings.ToUpper(msg.RequestContextId)
	m.Consumer = msg.Consumer
}

func (m *DocMsgKillRequestContext) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgKillRequestContext
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
