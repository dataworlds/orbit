package main

import (
	"fmt"

	"github.com/dataworlds/orbit"
	"github.com/dataworlds/orbit/player"
)

func test(p player.Player) {
	fmt.Println("Onspawn working")
}

func main() {
	server := orbit.New()

	server.OnSpawn = test

	server.Listen(":19132")
}
