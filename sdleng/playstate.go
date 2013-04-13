package main

import (
	"fmt"
	"github.com/zeroshade/tonberry"
	"image"
)

var (
	playState PlayState
)

const (
	TILE_RED = iota
	TILE_GREEN
	TILE_BLUE
	TILE_CENTER
	TILE_TOP
	TILE_TOPRIGHT
	TILE_RIGHT
	TILE_BOTTOMRIGHT
	TILE_BOTTOM
	TILE_BOTTOMLEFT
	TILE_LEFT
	TILE_TOPLEFT
)

type PlayState struct {
	test tonberry.GameObject
	cam  tonberry.Camera
}

func TouchWall(r image.Rectangle, t []tonberry.Tile) bool {
	for _, v := range t {
		if v.Type >= TILE_CENTER && v.Type <= TILE_TOPLEFT && r.Overlaps(v.Box) {
			return true
		}
	}
	return false
}

func (b *PlayState) Init() {
	tonberry.InitTileMap("map.json", "base")
	mHeight := tonberry.MapHeight("base")
	mWidth := tonberry.MapWidth("base")
	tiles := tonberry.MapTiles("base")
	b.test = tonberry.NewMoveable("dot.png", image.Rect(0, 0, 20, 20), func(ty int, box image.Rectangle) bool {
		if ty == tonberry.CHKX {
			return (box.Min.X < 0 || box.Max.X > mWidth || TouchWall(box, tiles))
		} else if ty == tonberry.CHKY {
			return (box.Min.Y < 0 || box.Max.Y > mHeight || TouchWall(box, tiles))
		}
		return true
	})
	b.cam = tonberry.NewCamera(image.Rect(0, 0, LEVEL_WIDTH, LEVEL_HEIGHT))
	fmt.Println("PlayState Init")
}

func (b *PlayState) Clean() {
	fmt.Println("PlayState clean")
}

func (b *PlayState) Pause() {
	fmt.Println("PlayState Paused")
	b.test.ResetMovement()
}

func (b *PlayState) Resume() {
	fmt.Println("PlayState Resumed")
}

func (b *PlayState) HandleEvents(ev tonberry.Event, g tonberry.Game) {
	switch e := ev.(type) {
	case tonberry.QuitEvent:
		g.Quit()
	case tonberry.KeyboardEvent:
		b.test.HandleInput(e)
		if e.Type == tonberry.KEYDOWN {
			switch e.Keysym.Sym {
			case tonberry.K_SPACE:
				g.PushState(&pausedState)
			case tonberry.K_ESCAPE:
				g.Quit()
			}
		}
	}
}

func (b *PlayState) Update(delta uint32, g tonberry.Game) {
	b.test.Update(delta)
}

func (b *PlayState) Draw(g tonberry.Game) {
	// g.GetScreen().FillRect(&g.GetScreen().Clip_rect, tonberry.MapRGB(g, 0xFF, 0xFF, 0xFF))
	b.test.SetCamera(&b.cam, "base")
	tonberry.DrawMap("base", b.cam, g.GetScreen())
	b.test.DrawCam(g.GetScreen(), b.cam)
}
