package gfx

import eb "github.com/hajimehoshi/ebiten/v2"

type Parallax struct {
	layers           *Atlas
	speedMultipliers float64
}

func NewParallax(img *eb.Image, layerWidth int, layers int) *Parallax {
	// atlas := NewAtlas(img, 1, 4, 64, 32)

	return nil
}
