package tonberry

import (
	"github.com/zeroshade/Go-SDL/sdl"
)

type Timer struct {
	startT  uint32
	pauseT  uint32
	started bool
	paused  bool
}

func NewTimer() *Timer {
	return &Timer{0, 0, false, false}
}

func (t *Timer) Start() {
	t.started = true
	t.paused = false
	t.startT = sdl.GetTicks()
}

func (t *Timer) Stop() {
	t.started = false
	t.paused = false
}

func (t *Timer) GetTicks() uint32 {
	if t.started {
		if t.paused {
			return t.pauseT
		} else {
			return sdl.GetTicks() - t.startT
		}
	}

	return 0
}

func (t *Timer) Pause() {
	if t.started && !t.paused {
		t.paused = true

		t.pauseT = sdl.GetTicks() - t.startT
	}
}

func (t *Timer) Unpause() {
	if t.paused {
		t.paused = false
		t.startT = sdl.GetTicks() - t.pauseT
		t.pauseT = 0
	}
}

func (t *Timer) IsStarted() bool {
	return t.started
}

func (t *Timer) IsPaused() bool {
	return t.paused
}
