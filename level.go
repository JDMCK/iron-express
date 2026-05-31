package main

import (
	"image/color"
	"iron-express/core"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

const TileSize = 16

type Tile struct {
	img       *ebiten.Image
	collision bool
}

type Layer struct {
	tilesWidth  int // in tiles
	tilesHeight int // in tiles
	tiles       []Tile
	Collider    core.Collider
}

type Level struct {
	layers  []Layer
	enemies []Enemy
}

func NewTile(img *ebiten.Image, collision bool) Tile {

	return Tile{
		img:       img,
		collision: collision,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, x, y float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(t.img, &op)
}

func NewLayer(width, height int) Layer {
	tiles := make([]Tile, width*height)

	// give each tile in the layer a random color value
	for i := range tiles {
		img := ebiten.NewImage(TileSize, TileSize)
		randR := uint8(rand.UintN(256))
		randG := uint8(rand.UintN(256))
		randB := uint8(rand.UintN(256))
		img.Fill(color.RGBA{randR, randG, randB, 255})
		tiles[i] = NewTile(img, true)
	}

	layer := Layer{
		tilesWidth:  width,
		tilesHeight: height,
		tiles:       tiles,
	}

	// attach a collider to the newly made layer
	layerX, layerY := layer.getTilePos(0)
	layerTopLeft := core.Vector2{X: layerX, Y: layerY}
	collider := core.NewCollider(
		layerTopLeft, width*TileSize, height*TileSize)

	layer.Collider = collider

	return layer
}

// translates index to x y position of tile
func (l *Layer) getTilePos(i int) (float64, float64) {
	i %= l.tilesWidth * l.tilesHeight // wrap index to prevent overflow
	return float64(i % l.tilesWidth * TileSize), float64(i / l.tilesWidth * TileSize)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	for i, t := range l.tiles {
		x, y := l.getTilePos(i)
		t.Draw(screen, x, y)
	}
}

func NewLevel() Level {
	return Level{
		layers: []Layer{NewLayer(5, 5)},
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	for _, l := range l.layers {
		l.Draw(screen)
	}
	// TODO draw enemies
}
