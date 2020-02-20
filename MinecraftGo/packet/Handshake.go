package packet

import (
	"../dataTypes"
	"encoding/hex"
	"fmt"
)

type Handshaking struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      uint16
	NextState       int
}

func (h *Handshaking) Handle(packet []byte, connection *Connection) {
	fmt.Println(hex.Dump(packet))
	fmt.Println("We are now handling a handshake packet")

	protocolVersion, cursor := dataTypes.ReadVarInt(packet)
	h.ProtocolVersion = protocolVersion
	fmt.Printf("Protocol Version %d (end %d)\n", protocolVersion, cursor)

	somethingElse, end := dataTypes.ReadVarInt(packet[cursor:])
	cursor += end
	fmt.Println("Mysterious value", somethingElse)

	serverAddress, end := dataTypes.ReadString(packet[cursor:])
	cursor += end
	h.ServerAddress = serverAddress
	fmt.Printf("Server Address lives at %d: '%s'\n", cursor, h.ServerAddress)

	fmt.Println(packet[cursor:])
	serverPort, end := dataTypes.ReadUnsignedShort(packet[cursor:])
	cursor += end
	fmt.Println("Server port", serverPort)

	nextState, end := dataTypes.ReadVarInt(packet[cursor:])
	h.NextState = nextState

	fmt.Println("Next State", nextState)
}
