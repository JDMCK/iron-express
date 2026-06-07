package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	focusX  float64
	focusY  float64
	offsetX float64
	offsetY float64
	zoom    float64
}

func NewCamera(zoom float64) *Camera {
	return &Camera{
		zoom: zoom,
	}
}

func (c *Camera) CenterScreenOffset(screenWidth, screenHeight int) {
	c.offsetX = float64(screenWidth) / 2
	c.offsetY = float64(screenHeight) / 2
}

func (c *Camera) DrawOptions() *eb.DrawImageOptions {
	op := eb.DrawImageOptions{}
	op.GeoM.Translate(-c.focusX, -c.focusY)
	op.GeoM.Scale(c.zoom, c.zoom)
	op.GeoM.Translate(c.offsetX, c.offsetY)
	return &op
}
