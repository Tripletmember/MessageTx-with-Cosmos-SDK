package params

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
