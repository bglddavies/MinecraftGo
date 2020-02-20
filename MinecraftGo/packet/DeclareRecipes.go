package packet

type DeclareRecipes struct {
	NumRecipes int      `proto:"varInt"`
	Recipes    []Recipe `proto:"array"`
}

func (dr *DeclareRecipes) GetPacketId() int {
	return 0x5b
}

func (dr *DeclareRecipes) Handle(packet []byte, connection *Connection) {
	//Client Only Packet
}

type Recipe struct {
	ID   string `proto:"string"`
	Type string `proto:"string"`
	Data []byte `proto:"raw"`
}
