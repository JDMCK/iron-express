package main

import (
	"image/color"
	"iron-express/core"
	"iron-express/gfx"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	rows     int
	cols     int
	tiles    []*Tile
	Collider *core.Collider
	parallax *gfx.Parallax
}

func NewLayer(width, height int) Layer {
	tiles := make([]*Tile, width*height)

	// give each tile in the layer a random color value
	for i := range tiles {
		img := ebiten.NewImage(TileSize, TileSize)
		randR := uint8(rand.UintN(256))
		randG := uint8(rand.UintN(256))
		randB := uint8(rand.UintN(256))
		img.Fill(color.RGBA{randR, randG, randB, 255})
		x, y := calcTilePos(i, width, TileSize, TileSize)
		tiles[i] = NewTile(x, y, img, true)
	}

	layer := Layer{
		rows:  height,
		cols:  width,
		tiles: tiles,
	}

	// attach a collider to the newly made layer
	layerX, layerY := layer.tiles[0].x, layer.tiles[0].y
	layerTopLeft := core.Vector2{X: float64(layerX), Y: float64(layerY)}
	collider := core.NewCollider(
		layerTopLeft, width*TileSize, height*TileSize)

	layer.Collider = collider

	return layer
}

// translates index to x y position of tile
func calcTilePos(i int, width, tileWidth, tileHeight int) (int, int) {
	return (i % width) * tileWidth, (i / width) * tileHeight
}
func (l *Layer) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	for _, t := range l.tiles {
		t.Draw(screen, op)
	}
}
