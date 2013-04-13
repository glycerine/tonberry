package tonberry

import (
	"github.com/zeroshade/Go-SDL/sdl"
	"image"
)

type Camera struct {
	sdl.Rect
}

func NewCamera(bounds image.Rectangle) Camera {
	var c Camera
	c.Rect = sdl.RectFromGoRect(bounds)
	return c
}

type Screen struct {
	*sdl.Surface
}

type Event interface {
	sdl.Event
}

type QuitEvent struct {
	*sdl.QuitEvent
}

type KeyboardEvent struct {
	*sdl.KeyboardEvent
}

func MapRGB(game Game, r, g, b uint8) uint32 {
	return sdl.MapRGB(game.GetScreen().Format, r, g, b)
}

func load_image(file string) *sdl.Surface {
	loadedImage := sdl.Load(file)
	var optimizedImg *sdl.Surface
	if loadedImage != nil {
		optimizedImg = sdl.DisplayFormat(loadedImage)
		loadedImage.Free()
	}

	if optimizedImg != nil {
		colorKey := sdl.MapRGB(optimizedImg.Format, 0x00, 0xFF, 0xFF)
		optimizedImg.SetColorKey(sdl.SRCCOLORKEY, colorKey)
	}
	return optimizedImg
}

func apply_surface(x, y int16, src, dst *sdl.Surface) {
	apply_surface_clip(x, y, src, dst, nil)
}

func apply_surface_clip(x, y int16, src, dst *sdl.Surface, clip *sdl.Rect) {
	offset := sdl.Rect{X: x, Y: y}
	dst.Blit(&offset, src, clip)
}
