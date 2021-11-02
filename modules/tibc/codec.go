package tibc

import (
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibccore "github.com/bianjieai/tibc-go/modules/tibc/core"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		tibctranfer.AppModuleBasic{},
		tibccore.AppModuleBasic{},
	)
}
