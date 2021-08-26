package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestFarm() {
	cases := []SubTest{
		{
			"createPool",
			createPool,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func createPool(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0a89010a86010a1b2f697269736d6f642e6661726d2e4d7367437265617465506f6f6c12670a09627573642d697269731a057374616b6520ac022a0f0a0a6e6f646530746f6b656e12013132130a0a6e6f646530746f6b656e12053130303030422a69616131737278306c707535346b766372353774683476687a6b76393276336b6d6b363730723776743212640a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a210384228c1b55c6af8d4299e5be90d4a0b539b4cbeeffb67145e6a15d08eb65ba7d12040a020801180712100a0a0a057374616b6512013210c09a0c1a408aee9a1afb96272d02e7a92a4e058761aaa3ca5923364014bd4a5e2da3f766fa17afe369404977ab6c8c57a70dd290e819ce5c8f710e08545b79a57f4fd5ff82")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Farm.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
