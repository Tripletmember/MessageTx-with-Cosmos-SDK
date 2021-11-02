package ibc

import (
	ibctransfer "github.com/cosmos/ibc-go/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/modules/core"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
	)
}
