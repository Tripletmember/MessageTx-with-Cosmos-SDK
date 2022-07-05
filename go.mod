module github.com/kaifei-bianjie/msg-parser

go 1.15

require (
	github.com/bianjieai/tibc-go v0.3.0
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/cosmos/ibc-go v1.1.0
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/irismod v1.5.2
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.13
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
)
