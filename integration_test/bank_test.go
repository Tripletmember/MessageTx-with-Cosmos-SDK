package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestBank() {
	cases := []SubTest{
		{
			"TestSend",
			send,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func send(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0a94010a91010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412710a2d636f736d6f7331373471796c3032637570797171373763717174646c306672646136646c3372706c3439703364122d636f736d6f73313675726e79677a3472726a786c68793578386b6e336167377030756668756b6c7937746461321a110a0475626967120931303030303030303012580a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102e570922aa1131475a67207eb1070cd730a7b66620878b967dde28c91aafbd8a412040a0208011802120410c09a0c1a40aff0ce2672a3f87fa0872d3d3a19dcabef631dcb5ea442ee77e552be055cf54d18b127a61e416aaccb550adc600aa1cede1fb618d67f2da27f5b844e32548492")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Bank.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
