package orbit

import (
	"github.com/dataworlds/orbit/player"
)

type server struct {
	OnSpawn func(p player.Player)
}

func New() *server {
	return &server{}
}

func (server *server) Listen(addr string) {
	if server.OnSpawn != nil {
		server.OnSpawn(player.Player{})
	}

}
