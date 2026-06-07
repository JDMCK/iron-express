package gfx

import eb "github.com/hajimehoshi/ebiten/v2"

type Parallax struct {
	layers           *Atlas
	speedMultipliers []float64
}

func NewParallax(atlas *Atlas, speedMultipliers []float64) *Parallax {
	return &Parallax{
		layers:           atlas,
		speedMultipliers: speedMultipliers,
	}
}

func (p *Parallax) Update() {

}

func (p *Parallax) Draw(screen *eb.Image, op *eb.DrawImageOptions) {

}
