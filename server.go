package orbit

import (
	"fmt"

	"github.com/dataworlds/orbit/player"
)

type server struct {
	OnSpawn func(p player.Player)
}

func New() *server {
	return &server{}
}

func (*server) Listen(addr string) {
	fmt.Println("working")

}
