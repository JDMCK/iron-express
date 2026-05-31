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
	img      *eb.Image
}

func NewCollider(position Vector2, width int, height int) Collider {

	img := eb.NewImage(width, height)
	img.Fill(color.Black)

	return Collider{
		Position: position,
		Width:    width,
		Height:   height,
		Enabled:  true,
		img:      img,
	}
}

func IsCollided(c1 Collider, c2 Collider) bool {
	c1Left := float64(c1.Position.X)
	c1Right := c1.Position.X + float64(c1.Width)
	c1Top := float64(c1.Position.Y)
	c1Bottom := c1.Position.Y + float64(c1.Height)

	c2Left := float64(c2.Position.X)
	c2Right := c2.Position.X + float64(c2.Width)
	c2Top := float64(c2.Position.Y)
	c2Bottom := c2.Position.Y + float64(c2.Height)

	return c1Left < c2Right &&
		c1Right > c2Left &&
		c1Top < c2Bottom &&
		c1Bottom > c2Top
}

func (c *Collider) Draw(screen *eb.Image) {
	op := eb.DrawImageOptions{}
	op.GeoM.Translate(c.Position.X, c.Position.Y)
	screen.DrawImage(c.img, &op)
}
