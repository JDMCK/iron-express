package core

import (
	"image/color"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Collider struct {
	Position Vector2
	Width    int
	Height   int
	Enabled  bool
}

func NewCollider() Collider {
	position := Vector2{
		X: 0.0,
		Y: 0.0,
	}
	return Collider{
		Position: position,
		Width:    1,
		Height:   1,
		Enabled:  true,
	}
}

func IsCollided(c1 Collider, c2 Collider) bool {
	c1Left := float64(c1.Position.X)
	c1Right := c1.Position.X + float64(c1.Width)
	c1Top := float64(c1.Position.Y)
	c1Bottom := c1.Position.Y - float64(c1.Height)

	c2Left := float64(c2.Position.X)
	c2Right := c2.Position.X + float64(c2.Width)
	c2Top := float64(c2.Position.Y)
	c2Bottom := c2.Position.Y - float64(c2.Height)

	return c1Left < c2Right &&
		c1Right > c2Left &&
		c1Top > c2Bottom &&
		c1Bottom < c2Top
}

func (c *Collider) Draw(screen *eb.Image) {
	op := eb.DrawImageOptions{}
	op.GeoM.Translate(c.Position.X, c.Position.Y)
	img := eb.NewImage(c.Width, c.Height)

	img.Fill(color.Black)
	screen.DrawImage(img, &op)
}
