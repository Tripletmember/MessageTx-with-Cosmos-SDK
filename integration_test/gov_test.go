package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestGov() {
	cases := []SubTest{
		{
			"submitProposal",
			submitProposal,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func submitProposal(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0aa7010aa4010a252f636f736d6f732e676f762e763162657461312e4d73675375626d697450726f706f73616c127b0a390a202f636f736d6f732e676f762e763162657461312e5465787450726f706f73616c12150a057465737431120cc3a6c2b5c28bc3a8c2afc295120f0a04756269671207313030303030301a2d636f736d6f7331373471796c3032637570797171373763717174646c306672646136646c3372706c343970336412690a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102e570922aa1131475a67207eb1070cd730a7b66620878b967dde28c91aafbd8a412040a020801180312150a0e0a047562696712063235303030301080ade2041a40f4118f021834be325aea7fecf7ce725400f42c948949afdcc5ecc337ad9526a9253ee39958b09d310b84dc971826cda32fc8fc0bd689774b81c942f210f3280c")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Gov.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
