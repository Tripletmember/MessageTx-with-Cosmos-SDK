package tibc

import (
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
)

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence         uint64     `bson:"sequence"`
	Port             string     `bson:"port"`
	SourceChain      string     `bson:"source_chain"`
	DestinationChain string     `bson:"destination_chain"`
	RelayChain       string     `bson:"relay_chain"`
	Data             PacketData `bson:"data"`
}

type Height struct {
	RevisionNumber uint64 `bson:"revision_number"`
	RevisionHeight uint64 `bson:"revision_height"`
}

//NonFungibleTokenPacketData
type PacketData struct {
	Class          string `bson:"class"`
	Id             string `bson:"id"`
	Uri            string `bson:"uri"`
	Sender         string `bson:"sender"`
	Receiver       string `bson:"receiver"`
	AwayFromOrigin bool   `bson:"away_from_origin"`
}

func loadHeight(height tibcclient.Height) Height {
	return Height{
		RevisionNumber: height.RevisionNumber,
		RevisionHeight: height.RevisionHeight}
}

func loadPacket(packet tibcpacket.Packet) Packet {
	return Packet{
		Sequence:         packet.Sequence,
		Port:             packet.Port,
		SourceChain:      packet.SourceChain,
		DestinationChain: packet.DestinationChain,
		RelayChain:       packet.RelayChain,
		Data:             UnmarshalPacketData(packet.GetData()),
	}
}
