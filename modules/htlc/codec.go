package htlc

import (
	"github.com/irisnet/irismod/modules/htlc"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
