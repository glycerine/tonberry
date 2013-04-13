package tonberry

type GameState interface {
	Init()
	Clean()
	Pause()
	Resume()
	HandleEvents(Event, Game)
	Update(uint32, Game)
	Draw(Game)
}

type GameObject interface {
	Update(deltaTicks uint32)
	ResetMovement()
	HandleInput(KeyboardEvent)
	SetCamera(*Camera, string)
	Draw(Screen)
	DrawCam(Screen, Camera)
}
