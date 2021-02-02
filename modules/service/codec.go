package service

import (
	"github.com/irisnet/irismod/modules/service"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
