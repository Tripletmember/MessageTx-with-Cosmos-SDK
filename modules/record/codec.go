package record

import (
	"github.com/irisnet/irismod/modules/record"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
