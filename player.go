package tonberry

import (
	"image"
)

const (
	CHKX = iota
	CHKY
)

type player struct {
	spr        Sprite
	box        image.Rectangle
	xVel, yVel int
	velInc     int
	boundsChk  func(int, image.Rectangle) bool
}

func NewMoveable(file string, bounds image.Rectangle, bcheck func(int, image.Rectangle) bool) GameObject {
	return &player{
		spr:       NewSprite(file),
		velInc:    500,
		box:       bounds,
		boundsChk: bcheck,
	}
}

func (d *player) ResetMovement() {
	d.xVel = 0
	d.yVel = 0
}

func (d *player) Update(deltaTicks uint32) {
	var multiplier float64 = (float64(deltaTicks) / 1000.0)

	d.box = d.box.Add(image.Pt(int(float64(d.xVel)*multiplier), 0))
	if d.boundsChk(CHKX, d.box) {
		d.box = d.box.Sub(image.Pt(int(float64(d.xVel)*multiplier), 0))
	}

	d.box = d.box.Add(image.Pt(0, int(float64(d.yVel)*multiplier)))
	if d.boundsChk(CHKY, d.box) {
		d.box = d.box.Sub(image.Pt(0, int(float64(d.yVel)*multiplier)))
	}
}

func (d *player) HandleInput(ev KeyboardEvent) {
	if ev.Type == KEYDOWN {
		switch ev.Keysym.Sym {
		case K_UP:
			d.yVel -= d.velInc
		case K_DOWN:
			d.yVel += d.velInc
		case K_LEFT:
			d.xVel -= d.velInc
		case K_RIGHT:
			d.xVel += d.velInc
		}
	} else if ev.Type == KEYUP && (d.xVel != 0 || d.yVel != 0) {
		switch ev.Keysym.Sym {
		case K_UP:
			d.yVel += d.velInc
		case K_DOWN:
			d.yVel -= d.velInc
		case K_LEFT:
			d.xVel += d.velInc
		case K_RIGHT:
			d.xVel -= d.velInc
		}
	}
}

func (d *player) Draw(sc Screen) {
	d.spr.Show(sc, d.box.Min.X, d.box.Min.Y)
}

func (d *player) DrawCam(sc Screen, c Camera) {
	d.spr.Show(sc, d.box.Min.X-int(c.X), d.box.Min.Y-int(c.Y))
}

func (d *player) SetCamera(c *Camera, level string) {
	c.X = int16((d.box.Min.X + d.box.Dx()/2) - int(c.W/2))
	c.Y = int16((d.box.Min.Y + d.box.Dy()/2) - int(c.H/2))

	if c.X < 0 {
		c.X = 0
	}
	if c.Y < 0 {
		c.Y = 0
	}
	if c.X > int16(maps[level].LWidth-c.W) {
		c.X = int16(maps[level].LWidth - c.W)
	}
	if c.Y > int16(maps[level].LHeight-c.H) {
		c.Y = int16(maps[level].LHeight - c.H)
	}
}
