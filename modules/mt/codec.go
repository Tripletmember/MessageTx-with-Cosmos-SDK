package mt

import (
	"github.com/irisnet/irismod/modules/mt"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(mt.AppModuleBasic{})
}
