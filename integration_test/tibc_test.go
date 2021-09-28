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
		}, {
			"RecvPacket",
			RecvPacket,
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

func RecvPacket(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0ade050adb050a222f746962632e636f72652e7061636b65742e76312e4d7367526563765061636b657412b4050a8901080412034e46541a0a746962632d74657374312209746962632d7465737432690a0574657374311206733232337931222a69616131777577726a77746c76797736366e703530717367766867797a36777566327339337432766b382a2a6961613134613436396b70713476797978786d6877676c6574773432767930633476783068356c787364300112f4030ac1020abe020a2c636f6d6d69746d656e74732f746962632d74657374312f746962632d746573742f73657175656e6365732f341220e2de2e3b50d049d2c401678b700c1679556d5079e6e51c95fe4386cb5e4f7bcd1a0c0801180120012a0400029419222a080112260204941920a41a78fe238aae077475e154e90b285d7a46ccbdeaef56d4d6c810efded21c5620222c0801120504069419201a2120acea3fef9fd556593a2fb5374bae5913c006afe2cc53cbc6fdf47ff3ba764c98222c08011205060a9419201a2120faa8ba33a7f0f3bbcf3989a7c3284a716279566642d2cef571b3e35b49e4251a222a0801122608169419200c6f4c6553b0d76bb52a3412a4dea09a453a1929be011d60fbae3a884631500720222a080112260a2e941920c9245187ebdb5a0f80d918ef257b7560466247f192eae4db2a3e54db1e0ace41200aad010aaa010a047469626312203362fa7095860f2e060fd1a443a6ab541c03bc71ba8c8b6d5c9cd4cc2e70b2e21a090801180120012a0100222708011201011a20c9c8849ed125cc7681329c4d27b83b1fc8acf7a865c9d1d1df575cca56f48dbe22250801122101b1fa483555c512177f8a63e199748a3666b25c8960325cf411fd0108ff65c012222508011221013c80938e278137394006b8240ff425818c7d33cf9b240015df0927fc20a918e01a0310cb0c222a69616131636871306e637434353066336e3974767a6a323470656379633533733377666477333430773512660a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102585b0763171ff17f3c1d0ccb906fac6591169df6ec2bf17ef7cd136001f29eea12040a020801182112120a0c0a057374616b65120331303010c09a0c1a40021bcd3a6d5e4ea217e9d8e63128961f1eacc9a71d2eb81aa8931e2ebb9f52791872b7b2e26d5e2f86cc20477d9c56a7dc45d533bc1979ed2769707f8fdf3471")
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
