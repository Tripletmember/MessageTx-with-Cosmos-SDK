package tibc

import (
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
)

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence         int64      `bson:"sequence"`
	Port             string     `bson:"port"`
	SourceChain      string     `bson:"source_chain"`
	DestinationChain string     `bson:"destination_chain"`
	RelayChain       string     `bson:"relay_chain"`
	Data             PacketData `bson:"data"`
}

type Height struct {
	RevisionNumber int64 `bson:"revision_number"`
	RevisionHeight int64 `bson:"revision_height"`
}

//NonFungibleTokenPacketData
type PacketData struct {
	Class          string `bson:"class"`
	Id             string `bson:"id"`
	Uri            string `bson:"uri"`
	Sender         string `bson:"sender"`
	Receiver       string `bson:"receiver"`
	AwayFromOrigin bool   `bson:"away_from_origin"`
	DestContract   string `bson:"dest_contract"`
}

// CleanPacket defines a type that carries data across different chains through TIBC
type CleanPacket struct {
	Sequence         int64  `bson:"sequence"`
	SourceChain      string `bson:"source_chain"`
	DestinationChain string `bson:"destination_chain"`
	RelayChain       string `bson:"relay_chain"`
}

func loadHeight(height tibcclient.Height) Height {
	return Height{
		RevisionNumber: int64(height.RevisionNumber),
		RevisionHeight: int64(height.RevisionHeight)}
}

func loadPacket(packet tibcpacket.Packet) Packet {
	return Packet{
		Sequence:         int64(packet.Sequence),
		Port:             packet.Port,
		SourceChain:      packet.SourceChain,
		DestinationChain: packet.DestinationChain,
		RelayChain:       packet.RelayChain,
		Data:             UnmarshalPacketData(packet.GetData()),
	}
}

func loadCleanPacket(packet tibcpacket.CleanPacket) CleanPacket {
	return CleanPacket{
		Sequence:         int64(packet.Sequence),
		SourceChain:      packet.SourceChain,
		DestinationChain: packet.DestinationChain,
		RelayChain:       packet.RelayChain,
	}
}
