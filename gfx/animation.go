package gfx

import (
	"iron-express/core"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	timer core.Timer
	atlas *Atlas
	row   int
}

type AnimationMap map[string]*Animation

func NewAnimation(a *Atlas, row int, duration int, frames int, loop bool) *Animation {
	return &Animation{
		atlas: a,
		timer: core.NewTimer(duration, frames, loop, nil),
		row:   row,
	}
}

func (a *Animation) Play() {
	a.timer.Play()
}

func (a *Animation) Restart() {
	a.timer.Restart()
}

// Like restart, but defaults to paused
func (a *Animation) Reset() {
	a.timer.Reset()
}

func (a *Animation) Update() {
	a.timer.Update()
}

func (a *Animation) Pause() {
	a.timer.Pause()
}

func (a *Animation) TogglePause() {
	a.timer.TogglePause()
}

func (a *Animation) Draw(screen *eb.Image, x, y int) {
	op := &eb.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	frame := a.atlas.GetFrame(a.row, a.timer.Cycles).(*eb.Image)
	screen.DrawImage(frame, op)
}

func (a *Animation) DrawAndUpdate(screen *eb.Image, x, y int) {
	a.Draw(screen, x, y)
	a.Update()
}
