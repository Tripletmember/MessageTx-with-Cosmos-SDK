package token

import (
	"github.com/irisnet/irismod/modules/token"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
