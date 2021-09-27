package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestTibc() {
	cases := []SubTest{
		{
			"NftTransfer",
			NftTransfer,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func NftTransfer(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0aaa010aa7010a292f746962632e617070732e6e66745f7472616e736665722e76312e4d73674e66745472616e73666572127a0a046c616e671204693030311a2a69616131636871306e637434353066336e3974767a6a3234706563796335337333776664773334307735222a69616131347361786e6c7664667463647037737539366335773932386e66767570367638667673646c6a2a09746962632d746573743209746962632d7465737412640a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102585b0763171ff17f3c1d0ccb906fac6591169df6ec2bf17ef7cd136001f29eea12040a020801181312100a0a0a057374616b6512013210c09a0c1a406f0a3cebbf23cd32480f0dd4becd7ffa28d1c68c92e851c3db32582c7b5dd34d3ccd037da0fee0edf3320436b019996c016d6e78467dfda59460338a9f3e6d46")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Tibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
