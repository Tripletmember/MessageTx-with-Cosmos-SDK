package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestStaking() {
	cases := []SubTest{
		{
			"Delegate",
			Delegate,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func Delegate(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0aa0010a9d010a232f636f736d6f732e7374616b696e672e763162657461312e4d736744656c656761746512760a2d636f736d6f73316571766b667468747272393367347039717370703534773664746a74726e32376770376a726c1234636f736d6f7376616c6f706572313836716874633632636636656a6c7433657277367a6b32386d6777386e653767376c726836741a0f0a047562696712073530303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a21038be4539785f0621a19066eb2da45c11dcc5fccff5d58e89c670bb80d251cc1b712040a020801180012130a0d0a04756269671205313030303010a0c21e1a40b16b015e1ed8284fa04881b852c8c2f714f6991ad756cae00bae2ee460aa2e826865cae6536755a90e2e3aedf73c4466563eb0c49d0144c921e071fac5b929e6")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Staking.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
