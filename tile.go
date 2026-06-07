package main

import eb "github.com/hajimehoshi/ebiten/v2"

const TileSize = 16

type Tile struct {
	img       *eb.Image
	x, y      int
	collision bool
}

func NewTile(x, y int, img *eb.Image, collision bool) *Tile {
	return &Tile{
		x:         x,
		y:         y,
		img:       img,
		collision: collision,
	}
}
func (t *Tile) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	newOp := eb.DrawImageOptions{}
	newOp.GeoM.Translate(float64(t.x), float64(t.y))
	newOp.GeoM.Concat(op.GeoM)
	screen.DrawImage(t.img, &newOp)
}
