package tonberry

import (
	"github.com/zeroshade/Go-SDL/sdl"
	"image"
	"runtime"
)

type spr struct {
	sur        *sdl.Surface
	box        image.Rectangle
	xVel, yVel int
	velInc     int
}

type Sprite interface {
	Move(deltaTicks uint32)
	HandleInput(*sdl.KeyboardEvent)
	Show(*sdl.Surface)
}

func NewSprite(file string, bounds image.Rectangle) Sprite {
	s := &spr{
		sur:    load_image(file),
		box:    bounds,
		velInc: 500,
	}

	runtime.SetFinalizer(s, func(s *spr) {
		s.sur.Free()
	})

	return s
}

func (s *spr) Move(deltaTicks uint32) {
	var multiplier float64 = (float64(deltaTicks) / 1000.0)

	s.box = s.box.Add(image.Pt(int(float64(s.xVel)*multiplier), 0))
	if (s.box.Min.X < 0) || (s.box.Max.X > levelWidth) {
		s.box = s.box.Sub(image.Pt(int(float64(s.xVel)*multiplier), 0))
	}

	s.box = s.box.Add(image.Pt(0, int(float64(s.yVel)*multiplier)))
	if (s.box.Min.Y < 0) || (s.box.Max.Y > levelHeight) {
		s.box = s.box.Sub(image.Pt(0, int(float64(s.yVel)*multiplier)))
	}
}

func (s *spr) HandleInput(ev *sdl.KeyboardEvent) {
	if ev.Type == sdl.KEYDOWN {
		switch ev.Keysym.Sym {
		case sdl.K_UP:
			s.yVel -= s.velInc
		case sdl.K_DOWN:
			s.yVel += s.velInc
		case sdl.K_LEFT:
			s.xVel -= s.velInc
		case sdl.K_RIGHT:
			s.xVel += s.velInc
		}
	} else if ev.Type == sdl.KEYUP {
		switch ev.Keysym.Sym {
		case sdl.K_UP:
			s.yVel += s.velInc
		case sdl.K_DOWN:
			s.yVel -= s.velInc
		case sdl.K_LEFT:
			s.xVel += s.velInc
		case sdl.K_RIGHT:
			s.xVel -= s.velInc
		}
	}
}

func (s *spr) Show(sc *sdl.Surface) {
	apply_surface(int16(s.box.Min.X), int16(s.box.Min.Y), s.sur, sc)
}
