package packet

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"

	"github.com/Ocelotworks/MinecraftGo/entity"
	"github.com/gofrs/uuid"
)

type LoginStart struct {
	Username string `proto:"string"`
}

func (ls *LoginStart) GetPacketId() int {
	return 0x00
}

//var testKey = []byte{0x30, 0x81, 0x9f, 0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x81, 0x8d, 0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00, 0xcc, 0x61, 0xf9, 0xef, 0x5a, 0xd0, 0xbc, 0x21, 0xde, 0x5b, 0x3c, 0xa6, 0x9e, 0xe7, 0x25, 0xd2, 0xc5, 0x04, 0xed, 0xf9, 0x9a, 0x6e, 0x97, 0xa0, 0x27, 0x9d, 0x95, 0x9f, 0x23, 0x64, 0xed, 0x21, 0xaa, 0xef, 0x8c, 0x70, 0x45, 0x52, 0xd5, 0xd1, 0xa3, 0xde, 0xb2, 0xee, 0x1a, 0x0b, 0xe1, 0x55, 0x8e, 0x3c, 0xa1, 0x82, 0xb1, 0x1a, 0x8c, 0x14, 0x39, 0x2b, 0x6d, 0xe5, 0x23, 0x46, 0xbc, 0x7c, 0x88, 0xbf, 0x75, 0xe3, 0xfb, 0x2b, 0x9f, 0x27, 0xfa, 0xb2, 0x1f, 0x5c, 0xf1, 0x99, 0x34, 0xe3, 0x11, 0x0e, 0xa4, 0xad, 0x72, 0xa6, 0xf0, 0x73, 0xf8, 0xab, 0x38, 0xb9, 0x93, 0x9b, 0xb6, 0x39, 0xe7, 0x8a, 0xd7, 0xf4, 0x34, 0x11, 0x9d, 0x8c, 0xc2, 0xb1, 0xbd, 0xe4, 0xea, 0xf9, 0x51, 0x3b, 0x05, 0x65, 0x08, 0xc9, 0x08, 0xed, 0x43, 0xc9, 0x9b, 0x0c, 0xeb, 0x22, 0x2c, 0xbb, 0xe1, 0xef, 0x02, 0x03, 0x01, 0x00, 0x01}

func (ls *LoginStart) Handle(packet []byte, connection *Connection) {
	fmt.Println("Username ", ls.Username)

	if connection.Minecraft.EnableEncryption {
		fmt.Println(connection.Key.PublicKey)

		publicKey, exception := x509.MarshalPKIXPublicKey(&connection.Key.PublicKey)

		if exception != nil {
			fmt.Println("Marshalling public key", exception)
			return
		}

		connection.VerifyToken = make([]byte, 4)
		rand.Read(connection.VerifyToken)

		connection.Player = &entity.Player{
			Username:       ls.Username,
			Properties:     []entity.PlayerProperty{},
			Gamemode:       0,
			Ping:           0,
			HasDisplayName: true,
			DisplayName: entity.ChatMessageComponent{
				Text: ls.Username,
			},
			EntityID: connection.Minecraft.ConnectedPlayers,
			X:        5,
			Y:        255,
			Z:        5,
			Yaw:      0,
			Pitch:    0,
		}

		encryptionPacket := Packet(&EncryptionRequest{
			ServerID:          "",
			PublicKeyLength:   len(publicKey),
			PublicKey:         publicKey,
			VerifyTokenLength: len(connection.VerifyToken),
			VerifyToken:       connection.VerifyToken,
		})

		connection.SendPacket(&encryptionPacket)
	} else {
		connection.Player = &entity.Player{
			UUID:           uuid.NewV3(uuid.Nil, "OfflinePlayer:"+ls.Username).Bytes(),
			Username:       ls.Username,
			Properties:     []entity.PlayerProperty{},
			Gamemode:       0,
			Ping:           0,
			HasDisplayName: true,
			DisplayName: entity.ChatMessageComponent{
				Text: ls.Username,
			},
			EntityID: connection.Minecraft.ConnectedPlayers,
			X:        5,
			Y:        255,
			Z:        5,
			Yaw:      0,
			Pitch:    0,
		}

		connection.Minecraft.StartPlayerJoin(connection)
	}
}