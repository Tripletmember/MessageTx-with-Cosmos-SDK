package farm

import (
	"github.com/irisnet/irismod/modules/farm"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(farm.AppModuleBasic{})
}
