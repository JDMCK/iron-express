package core

import (
	"image/color"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Clamp(min, max, val float64) float64 {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

const BorderWidth = 2

func DrawBorder(screen *eb.Image, x, y int, width, height int, color color.Color, op *eb.DrawImageOptions) {
	newOp := eb.DrawImageOptions{}
	newOp.GeoM.Translate(float64(x), float64(y))
	newOp.GeoM.Concat(op.GeoM)

	x0, y0 := newOp.GeoM.Apply(0, 0)
	x1, y1 := newOp.GeoM.Apply(float64(width), float64(height))

	vector.StrokeRect(screen,
		float32(x0), float32(y0),
		float32(x1-x0), float32(y1-y0),
		BorderWidth, color, false)
}
