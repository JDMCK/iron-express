package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

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

func (p *Parallax) Draw(screen *eb.Image, x float64, op *eb.DrawImageOptions) {
	for i := range p.layers.Frames {
		shiftSpeed := p.speedMultipliers[i]
		layerOffset := x * shiftSpeed
		frameWidth := float64(p.layers.FrameWidth)
		centerZone := int((x - layerOffset) / frameWidth) // the multiple of the frame width where the target currently is
		centerOffset := float64(centerZone) * frameWidth

		beforeGeoM := eb.DrawImageOptions{}
		beforeGeoM.GeoM.Translate(centerOffset-frameWidth+layerOffset, 0)
		beforeGeoM.GeoM.Concat(op.GeoM)

		centerOp := eb.DrawImageOptions{}
		centerOp.GeoM.Translate(centerOffset+layerOffset, 0)
		centerOp.GeoM.Concat(op.GeoM)

		afterGeoM := eb.DrawImageOptions{}
		afterGeoM.GeoM.Translate(centerOffset+frameWidth+layerOffset, 0)
		afterGeoM.GeoM.Concat(op.GeoM)

		screen.DrawImage(p.layers.GetFrameFromCoords(i, 0).(*eb.Image), &beforeGeoM)
		screen.DrawImage(p.layers.GetFrameFromCoords(i, 0).(*eb.Image), &centerOp)
		screen.DrawImage(p.layers.GetFrameFromCoords(i, 0).(*eb.Image), &afterGeoM)
	}
}
