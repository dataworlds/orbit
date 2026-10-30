// Experimental transport
package transport

import (
	"log"

	"github.com/dataworlds/orbit/player"
	"github.com/df-mc/go-nethernet"
	"github.com/df-mc/go-nethernet/endpoint"
)

// Start listening for connections
func Run(addr string, OnSpawn func(p player.Player)) {
	signal, err := endpoint.Serve(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer signal.Close()

	listener, err := (nethernet.ListenConfig{
		DisableTrickleICE: signal.DisableTrickleICE(),
	}).Listen(signal)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		if OnSpawn != nil {
			OnSpawn(player.Player{})
		}

		_ = conn
	}
}
