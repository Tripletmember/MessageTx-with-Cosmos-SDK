package integration

import (
	"encoding/hex"
	"fmt"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/codec"
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
		}, {
			"MsgAcknowledgement",
			MsgAcknowledgement,
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
func MsgAcknowledgement(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0adf070adc070a272f6962632e636f72652e6368616e6e656c2e76312e4d736741636b6e6f776c656467656d656e7412b0070ae001080112087472616e736665721a0a6368616e6e656c2d333122087472616e736665722a096368616e6e656c2d3032a1017b22616d6f756e74223a223130222c2264656e6f6d223a227472616e736665722f6368616e6e656c2d33312f7374616b65222c227265636569766572223a22696161316a7a386b74346e307566733475796730736d6536367368756c32716b68776476397a77643676222c2273656e646572223a22696161317376616e6e6876327a617865667138336d37747265673037387564666b33376c35726b787466227d3a0310d95940dcd396d0b1b2e0d81612117b22726573756c74223a2241513d3d227d1a86050a82030aff020a3261636b732f706f7274732f7472616e736665722f6368616e6e656c732f6368616e6e656c2d302f73657175656e6365732f31122008f7557ed51826fe18d84512bf24ec75001edbaf2123a477df72a0a9f3640a7c1a0d0801180120012a050002f6a301222d080112060204f6a301201a212073a0508aab36e7909b65372846f953a88a8d5aaccadfe83efda509fe09de41b7222d080112060406f6a301201a212066de425ace458527d718e2357799410d8c1603a978ef54a5ce0e24b18dcf9e30222d08011206060ef6a301201a2120840b5018133113e4b452d94ec6643d3f7c3402f82cca1ac1a8c9ad3571afcdce222d08011206081af6a301201a21205ff733bc3ddca5e265d7de48a72aa6842650d160202b17373080be57cadb5ab7222d080112060a2af6a301201a212087babc6e6d3f1a473af42a986bce716a1cd07dd00bf66bcc88b380f720d8ee84222d080112060e7cf6a301201a2120ae211f123ec9d5b21dec5ceaad2a6005bbe6624433ecd61c0686c5531a0a4e650afe010afb010a036962631220d1646e332046dcdcdd61ea2285fe24113b42e17a01288cea9a0c8049c5769a561a090801180120012a0100222508011221016c2870e99eaa2453e238dd6af9c7660109a8a9667cff72a26c031cb1594f7353222708011201011a205dfd19ed456064093ac720743ac192a01a3cd3687ae9b88a0fd5e134019d6a16222708011201011a20fa3672853e691148a59a717c6fb7c3f4c27ae5adf3d10f1609a4c9a3a9a37e4322250801122101128f4bee7ae00fd3197f860dcbece39442d264f8fae3fddee86c925656baa159222708011201011a2023a7499b9940678a0fc560d0b9b9d749919ab67ad81896765bf73b9fda33bb44220310fc512a2a696161316a7a386b74346e307566733475796730736d6536367368756c32716b68776476397a7764367612660a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103adbd43a6a6d2e451d0378dc0bf765a21ebc3eac27c2605f1107794b64be4568312040a020801181a12120a0c0a05756e79616e120336303010e0a7121a40ee9e69fa27ae3d8e32001d365feaae3000afbbdb9e93c58bdd4c220104662ece1f86292deed96004bb611b6efe34556a15cdbf800271fb45557a43d586d58159")
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
