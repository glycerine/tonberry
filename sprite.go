package tonberry

import (
	"github.com/zeroshade/Go-SDL/sdl"
	"image"
	"runtime"
)

type spr struct {
	sur *sdl.Surface
}

type Sprite interface {
	Show(Screen, int, int)
}

func NewSprite(file string) Sprite {
	s := &spr{
		sur: load_image(file),
	}
	if s.sur == nil {
		panic("Sprite unable to load image")
	}
	runtime.SetFinalizer(s, func(s *spr) {
		s.sur.Free()
	})

	return s
}

func (s *spr) Show(sc Screen, x, y int) {
	apply_surface(int16(x), int16(y), s.sur, sc.Surface)
}

type multiSprite struct {
	sur   *sdl.Surface
	clips []sdl.Rect
}

type ClipSprite interface {
	Show(int, Screen, int, int)
}

func NewClipSprite(file string, clips []image.Rectangle) ClipSprite {
	ms := &multiSprite{
		sur: load_image(file),
	}
	ms.clips = make([]sdl.Rect, len(clips))
	for i, j := range clips {
		ms.clips[i] = sdl.RectFromGoRect(j)
	}
	return ms
}

func (ms *multiSprite) Show(clip int, sc Screen, x, y int) {
	apply_surface_clip(int16(x), int16(y), ms.sur, sc.Surface, &ms.clips[clip])
}
