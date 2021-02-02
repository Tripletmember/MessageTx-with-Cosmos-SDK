package nft

import (
	"github.com/irisnet/irismod/modules/nft"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
