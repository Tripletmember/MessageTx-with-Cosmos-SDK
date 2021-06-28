package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/modules/ibc"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestIbc() {
	cases := []SubTest{
		{
			"CreateClient",
			CreateClient,
		},
		{
			"GetIbcPacketDenom",
			GetIbcPacketDenom,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func CreateClient(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0a90030a8d030a232f6962632e636f72652e636c69656e742e76312e4d7367437265617465436c69656e7412e5020aaa010a2b2f6962632e6c69676874636c69656e74732e74656e6465726d696e742e76312e436c69656e745374617465127b0a09626966726f73742d321204080110031a03088c0622030884072a0308d80432003a06080210a38c0d42190a090801180120012a0100120c0a02000110211804200c300142190a090801180120012a0100120c0a02000110201801200130014a07757067726164654a10757067726164656449424353746174651286010a2e2f6962632e6c69676874636c69656e74732e74656e6465726d696e742e76312e436f6e73656e737573537461746512540a0c08aadeb9800610bbadf08a0112220a2051c77a4f5ac9d60247b465bb42d81f3837bbe54a5e01e394dd2369fe089c26141a20b8c97bc5436f53aac003e94d194e75bc4167e6874ebbc2a8b327912bd6ee2f551a2d636f736d6f73313675726e79677a3472726a786c68793578386b6e336167377030756668756b6c79377464613212640a4e0a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a210265facc37ebf82d3732778610528478e402832ec5d1913efe13b363b8ec58054312040a02080112120a0c0a047562696712043234373410ff84061a40138e01c63715448b3a6a2546075c46edb138bbcda005fe9914c1fd45eccfdc175eeb5b4191faacd45674da4aa568e654e1d1bc338a9deef0623421c4cecc1922")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Ibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func GetIbcPacketDenom(s IntegrationTestSuite) {
	packet := ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-9",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
		Data: ibc.PacketData{
			Denom:  "transfer/channel-9/umuon",
			Amount: 1,
		},
	}

	fmt.Println("Atom Iris => Cosmos: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))
	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-9",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
		Data: ibc.PacketData{
			Denom:  "unyan",
			Amount: 1,
		},
	}

	fmt.Println("Iris Iris => Cosmos: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))

	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-1",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-9",
		Data: ibc.PacketData{
			Denom:  "umuon",
			Amount: 1,
		},
	}

	fmt.Println("Atom Cosmos => Iris: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))
	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-1",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-9",
		Data: ibc.PacketData{
			Denom:  "transfer/channel-1/unyan",
			Amount: 1,
		},
	}

	fmt.Println("Iris Cosmos => Iris: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))

}
