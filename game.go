package tonberry

import (
	"fmt"
	"github.com/zeroshade/Go-SDL/sdl"
)

type game struct {
	running, fullscreen bool
	screen              *sdl.Surface
	states              []GameState
}

type Game interface {
	Init(title string, w, h, bpp int, fullscreen bool)
	IsRunning() bool
	HandleEvents()
	Draw()
	Update(deltaTicks uint32)
	Quit()
	Clean()
	ChangeState(GameState)
	PushState(GameState)
	PopState()
	GetScreen() Screen
}

func NewGame(title string, w, h, bpp int, fullscreen bool) Game {
	g := &game{running: true, screen: nil}
	g.Init(title, w, h, bpp, fullscreen)
	return g
}

func (g *game) GetScreen() Screen {
	return Screen{g.screen}
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

	// g.test = NewSprite("dot.png", image.Rect(0, 0, 20, 20))

	g.states = make([]GameState, 0)

	fmt.Println("Game Initialized successfully")
}

func (g *game) Clean() {
	for _, v := range g.states {
		v.Clean()
	}
	g.states = []GameState{}
	g.screen.Free()
	sdl.Quit()
}

func (g *game) HandleEvents() {
	var p Event
	for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
		switch e := ev.(type) {
		case *sdl.QuitEvent:
			p = QuitEvent{e}
		case *sdl.KeyboardEvent:
			p = KeyboardEvent{e}
		}
		g.states[len(g.states)-1].HandleEvents(p, g)
	}
}

func (g *game) Quit() {
	g.running = false
}

func (g *game) Update(deltaTicks uint32) {
	// g.test.Move(deltaTicks)
	g.states[len(g.states)-1].Update(deltaTicks, g)
}

func (g *game) Draw() {
	g.states[len(g.states)-1].Draw(g)
	// g.screen.FillRect(&g.screen.Clip_rect, sdl.MapRGB(g.screen.Format, 0xFF, 0xFF, 0xFF))

	// g.test.Show(g.screen)

	g.screen.Flip()
}

func (g *game) ChangeState(st GameState) {
	l := len(g.states)
	if l != 0 {
		g.states[l-1].Clean()
		g.states = g.states[:l-1]
	}

	st.Init()
	g.states = append(g.states, st)
}

func (g *game) PushState(st GameState) {
	l := len(g.states)
	if l != 0 {
		g.states[l-1].Pause()
	}

	st.Init()
	g.states = append(g.states, st)
}

func (g *game) PopState() {
	l := len(g.states)
	if l != 0 {
		g.states[l-1].Clean()
		g.states = g.states[:l-1]
	}

	if l-1 != 0 {
		g.states[l-2].Resume()
	}
}
