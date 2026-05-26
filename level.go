package main

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	img *ebiten.Image
}

type Layer struct {
	width  int
	height int
	tiles  []Tile
}

type Level struct {
	layers  []Layer
	enemies []Enemy
}
