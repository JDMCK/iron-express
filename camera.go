package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	// the x coord in world space to point the 0, 0 of the camera (top left corner)
	focusX float64
	focusY float64

	offsetX float64
	offsetY float64

	zoom float64
}

func NewCamera(zoom float64) *Camera {
	return &Camera{
		zoom: zoom,
	}
}

func (c *Camera) SetFocus(x, y float64, drag float64) {
	c.focusX = x
	c.focusY = y
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
