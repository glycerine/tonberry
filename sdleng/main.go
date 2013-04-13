package main

import (
	"github.com/zeroshade/tonberry"
)

const (
	LEVEL_WIDTH  = 640
	LEVEL_HEIGHT = 480
)

func main() {
	g := tonberry.NewGame("Test", LEVEL_WIDTH, LEVEL_HEIGHT, 32, false)

	time := tonberry.NewTimer()
	g.ChangeState(&menu)

	time.Start()

	for g.IsRunning() {
		g.HandleEvents()
		g.Update(time.GetTicks())
		time.Start()
		g.Draw()
	}

	g.Clean()
}
