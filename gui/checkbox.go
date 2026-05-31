package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Checkbox struct {
	checked bool
	btn     *Button
}

func NewCheckbox(x, y int) *Checkbox {
	cb := Checkbox{
		checked: false,
	}
	btn := NewButton("", 32, 32, x, y, 0, color.Gray{150}, func() {
		cb.checked = !cb.checked
	})
	cb.btn = btn
	return &cb
}

func (c *Checkbox) Update() {
	if c.checked {
		c.btn.SetColor(color.RGBA{0, 255, 0, 255})
	} else {
		c.btn.SetColor(color.Gray{150})
	}
	c.btn.Update()
}

func (c *Checkbox) Draw(screen *ebiten.Image) {
	c.btn.Draw(screen)
}
