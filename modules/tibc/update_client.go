package tibc

import (
	cdc "github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgTIBCUpdateClient struct {
	// client unique identifier
	ChainName string `bson:"chain_name"`
	Header    string `bson:"header" yaml:"header"`
	Signer    string `bson:"signer" yaml:"signer"`
}

func (m *DocMsgTIBCUpdateClient) GetType() string {
	return MsgTypeTIBCUpdateClient
}

func (m *DocMsgTIBCUpdateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgTIBCUpdateClient)
	m.ChainName = msg.ChainName
	m.Signer = msg.Signer
	if msg.Header != nil {
		m.Header = utils.MarshalJsonIgnoreErr(v)
	}
}

func (m *DocMsgTIBCUpdateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTIBCUpdateClient
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
