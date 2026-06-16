package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	timer Timer
	atlas *Atlas
	row   int
}

type AnimationMap map[string]*Animation

func NewAnimation(a *Atlas, row int, duration int, frames int, loop bool) *Animation {
	return &Animation{
		atlas: a,
		timer: NewTimer(duration, frames, loop, nil),
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

func (a *Animation) Draw(screen *eb.Image, x, y int, facingRight bool, op *eb.DrawImageOptions) {
	newOp := &eb.DrawImageOptions{}
	if facingRight == false {
		newOp.GeoM.Scale(-1, 1)
		newOp.GeoM.Translate(float64(a.atlas.FrameWidth), 0)
	}
	newOp.GeoM.Translate(float64(x), float64(y))
	newOp.GeoM.Concat(op.GeoM)

	frame := a.atlas.GetFrameFromCoords(a.timer.Cycles, a.row).(*eb.Image)
	screen.DrawImage(frame, newOp)
}
