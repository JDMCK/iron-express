package gfx

import (
	"iron-express/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type Clip struct {
	timer core.Timer
	atlas Atlas
}

func NewClip(a Atlas, duration int, loop bool) Clip {
	return Clip{
		atlas: a,
		timer: core.NewTimer(duration, a.FrameCount, loop, nil),
	}
}

func (c *Clip) Play() {
	c.timer.Play()
}

func (c *Clip) Restart() {
	c.timer.Restart()
}

// Like restart, but defaults to paused
func (c *Clip) Reset() {
	c.timer.Reset()
}

func (c *Clip) Update() {
	c.timer.Update()
}

func (c *Clip) Pause() {
	c.timer.Pause()
}

func (c *Clip) TogglePause() {
	c.timer.TogglePause()
}

func (c *Clip) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	frame := c.atlas.GetFrame(c.timer.Cycles).(*ebiten.Image)
	screen.DrawImage(frame, op)
}

func (c *Clip) DrawAndUpdate(screen *ebiten.Image, x, y int) {
	c.Draw(screen, x, y)
	c.Update()
}
