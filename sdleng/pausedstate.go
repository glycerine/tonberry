package main

import (
	"fmt"
	"github.com/zeroshade/tonberry"
)

var (
	pausedState PausedState
)

type PausedState struct {
	test tonberry.Sprite
}

func (b *PausedState) Init() {
	b.test = tonberry.NewSprite("paused.bmp")
	fmt.Println("PausedState Init")
}

func (b *PausedState) Clean() {
	fmt.Println("PausedState clean")
}

func (b *PausedState) Pause() {
	fmt.Println("PausedState Paused")
}

func (b *PausedState) Resume() {
	fmt.Println("PausedState Resumed")
}

func (b *PausedState) HandleEvents(ev tonberry.Event, g tonberry.Game) {
	switch e := ev.(type) {
	case tonberry.QuitEvent:
		g.Quit()
	case tonberry.KeyboardEvent:
		if e.Type == tonberry.KEYDOWN {
			switch e.Keysym.Sym {
			case tonberry.K_SPACE:
				g.PopState()
			case tonberry.K_ESCAPE:
				g.Quit()
			}
		}
	}
}

func (b *PausedState) Update(delta uint32, g tonberry.Game) {
}

func (b *PausedState) Draw(g tonberry.Game) {
	b.test.Show(g.GetScreen(), 0, 0)
}
