package packet

type UpdateViewPosition struct {
	ChunkX int `proto:"varInt"`
	ChunkZ int `proto:"varInt"`
}

func (uvp *UpdateViewPosition) GetPacketId() int {
	return 0x41
}
