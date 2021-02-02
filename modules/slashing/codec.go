package slashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
