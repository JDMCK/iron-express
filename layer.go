package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	rows      int
	cols      int
	tiles     []*Tile
	Colliders []*Collider
	parallax  *Parallax
}

func NewLayer(cols, rows int, tileWidth, tileHeight int, atlas *Atlas, atlasIndices []int) *Layer {
	tiles := make([]*Tile, rows*cols)

	for i, _ := range tiles {
		atlasIndex := atlasIndices[i]
		if atlasIndex == -1 {
			continue
		}
		x, y := calcTilePos(i, cols, tileWidth, tileHeight)
		tiles[i] = NewTile(x, y, atlas.Frames[atlasIndex], true)
	}

	layer := Layer{
		rows:  rows,
		cols:  cols,
		tiles: tiles,
	}

	layer.Colliders = nil
	return &layer
}

// translates index to x y position of tile
func calcTilePos(i int, width, tileWidth, tileHeight int) (int, int) {
	return (i % width) * tileWidth, (i / width) * tileHeight
}
func (l *Layer) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	for _, t := range l.tiles {
		if t == nil {
			continue
		}
		t.Draw(screen, op)
	}
}
