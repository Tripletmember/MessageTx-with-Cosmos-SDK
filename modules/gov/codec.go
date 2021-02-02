package gov

import (
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/kaifei-bianjie/msg-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
		upgrade.AppModuleBasic{},
	)
}
