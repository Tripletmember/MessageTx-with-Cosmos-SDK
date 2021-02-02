package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestToken() {
	cases := []SubTest{
		{
			"IssueToken",
			IssueToken,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func IssueToken(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0a790a770a1c2f697269736d6f642e746f6b656e2e4d73674973737565546f6b656e12570a056b69747479120b4b6974747920546f6b656e22056b697474792880d0dbc3f4023080a094a58d1d3801422a69616131373664643074676e33386772706338687078666d776c36736c386a666d6b6e6567386d6b787212630a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a21024a97355aa9722c455b269376c805ab5d966d55934e7860565a48bdf58470d7c812040a020801182a120f0a090a047562696612013210c09a0c1a40621285916b58ed958b1ccd28b4d5ef35838ac5c6997e27deec6bc886980851a0407430c5c25084bdb704f34836d3c396a71810c21677b8590c25af2c47737a19")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Token.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
