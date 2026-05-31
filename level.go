package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

const TileSize = 16

type Tile struct {
	img       *ebiten.Image
	collision bool
}

type Layer struct {
	width  int // in tiles
	height int // in tiles
	tiles  []Tile
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

	for i := range tiles {
		img := ebiten.NewImage(TileSize, TileSize)
		randR := uint8(rand.UintN(256))
		randG := uint8(rand.UintN(256))
		randB := uint8(rand.UintN(256))
		img.Fill(color.RGBA{randR, randG, randB, 255})
		tiles[i] = NewTile(img, true)
	}

	return Layer{
		width:  width,
		height: height,
		tiles:  tiles,
	}
}

// translates index to x y position of tile
func (l *Layer) getTilePos(i int) (float64, float64) {
	i %= l.width * l.height // wrap index to prevent overflow
	return float64(i % l.width * TileSize), float64(i / l.width * TileSize)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	for i, t := range l.tiles {
		x, y := l.getTilePos(i)
		t.Draw(screen, x, y)
	}
}

func NewLevel() Level {
	return Level{
		layers: []Layer{NewLayer(20, 5)},
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	for _, l := range l.layers {
		l.Draw(screen)
	}
	// TODO draw enemies
}
