package core

import (
	"image/color"
	"math"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Collider struct {
	Position Vector2
	Width    int
	Height   int
	Enabled  bool
	img      *eb.Image
}

type CollisionDirection int

const (
	Up CollisionDirection = iota
	Down
	Left
	Right
	None
)

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

// Perform aabb collision and return and enum representing the
// direction of collision (from the perspective of c1)
func IntersectAABB(c1 Collider, c2 Collider) (CollisionDirection, float64) {
	c1Left := float64(c1.Position.X)
	c1Right := c1.Position.X + float64(c1.Width)
	c1Top := float64(c1.Position.Y)
	c1Bottom := c1.Position.Y + float64(c1.Height)

	c2Left := float64(c2.Position.X)
	c2Right := c2.Position.X + float64(c2.Width)
	c2Top := float64(c2.Position.Y)
	c2Bottom := c2.Position.Y + float64(c2.Height)

	if !(c1Left < c2Right) ||
		!(c1Right > c2Left) ||
		!(c1Top < c2Bottom) ||
		!(c1Bottom > c2Top) {
		return None, 0
	}

	dxLeft := math.Abs(c1Left - c2Right)
	dxRight := math.Abs(c1Right - c2Left)
	dyUp := math.Abs(c1Top - c2Bottom)
	dyDown := math.Abs(c1Bottom - c2Top)

	switch min(dxLeft, dxRight, dyUp, dyDown) {
	case dxLeft:
		return Left, dxLeft
	case dxRight:
		return Right, dxRight
	case dyUp:
		return Up, dyUp
	case dyDown:
		return Down, dyDown
	default:
		panic("How are we here")
	}
}

func (c *Collider) Draw(screen *eb.Image) {
	op := eb.DrawImageOptions{}
	op.GeoM.Translate(c.Position.X, c.Position.Y)
	screen.DrawImage(c.img, &op)
}
