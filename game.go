package tonberry

import (
	"fmt"
	"github.com/zeroshade/Go-SDL/sdl"
	"image"
)

type game struct {
	running, fullscreen bool
	screen              *sdl.Surface
	test                Sprite
}

var (
	levelWidth  int
	levelHeight int
)

type Game interface {
	Init(title string, w, h, bpp int, fullscreen bool)
	IsRunning() bool
	HandleEvents()
	Draw()
	Update(deltaTicks uint32)
	Quit()
	Clean()
}

func NewGame(title string, w, h, bpp int, fullscreen bool) Game {
	g := &game{running: true, screen: nil}
	g.Init(title, w, h, bpp, fullscreen)
	levelHeight = h
	levelWidth = w
	return g
}

func (g *game) IsRunning() bool {
	return g.running
}

func (g *game) Init(title string, w, h, bpp int, fullscreen bool) {
	sdl.Init(sdl.INIT_EVERYTHING)

	sdl.WM_SetCaption(title, "")

	flags := uint32(0)

	if fullscreen {
		flags = sdl.FULLSCREEN
	}
	g.fullscreen = fullscreen

	g.screen = sdl.SetVideoMode(w, h, bpp, flags)

	g.test = NewSprite("dot.png", image.Rect(0, 0, 20, 20))

	fmt.Println("Game Initialized successfully")
}

func (g *game) Clean() {
	g.screen.Free()
	sdl.Quit()
}

func (g *game) HandleEvents() {
	for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
		switch e := ev.(type) {
		case *sdl.QuitEvent:
			g.Quit()
		case *sdl.KeyboardEvent:
			g.test.HandleInput(e)
			switch e.Keysym.Sym {
			case sdl.K_ESCAPE:
				g.Quit()
			}
		}
	}
}

func (g *game) Quit() {
	g.running = false
}

func (g *game) Update(deltaTicks uint32) {
	g.test.Move(deltaTicks)
}

func (g *game) Draw() {
	g.screen.FillRect(&g.screen.Clip_rect, sdl.MapRGB(g.screen.Format, 0xFF, 0xFF, 0xFF))

	g.test.Show(g.screen)

	g.screen.Flip()
}
