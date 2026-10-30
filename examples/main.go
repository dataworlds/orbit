package main

import (
	"github.com/dataworlds/orbit"
	"github.com/dataworlds/orbit/player"
)

func test(p player.Player) {

}

func main() {
	server := orbit.New()

	server.OnSpawn = test

	server.Listen("19132")
}
