package main

import (
	"github.com/zeroshade/tonberry"
)

func main() {
	g := tonberry.NewGame("Test", 640, 480, 32, false)

	time := tonberry.NewTimer()
	time.Start()

	for g.IsRunning() {
		g.HandleEvents()
		g.Update(time.GetTicks())
		time.Start()
		g.Draw()
	}

	g.Clean()
}
