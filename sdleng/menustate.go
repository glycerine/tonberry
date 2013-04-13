package main

import (
	"fmt"
	"github.com/zeroshade/tonberry"
)

var (
	menu MenuState
)

type MenuState struct {
	test tonberry.Sprite
}

func (b *MenuState) Init() {
	b.test = tonberry.NewSprite("menustate.bmp")
	fmt.Println("MenuState Init")
}

func (b *MenuState) Clean() {
	fmt.Println("MenuState clean")
}

func (b *MenuState) Pause() {
	fmt.Println("MenuState Paused")
}

func (b *MenuState) Resume() {
	fmt.Println("MenuState Resumed")
}

func (b *MenuState) HandleEvents(ev tonberry.Event, g tonberry.Game) {
	switch e := ev.(type) {
	case tonberry.QuitEvent:
		g.Quit()
	case tonberry.KeyboardEvent:
		if e.Type == tonberry.KEYDOWN {
			switch e.Keysym.Sym {
			case tonberry.K_SPACE:
				g.ChangeState(&playState)
			case tonberry.K_ESCAPE:
				g.Quit()
			}
		}
	}
}

func (b *MenuState) Update(delta uint32, g tonberry.Game) {
}

func (b *MenuState) Draw(g tonberry.Game) {
	b.test.Show(g.GetScreen(), 0, 0)
}
