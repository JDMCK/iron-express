package main

import (
	"iron-express/gfx"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	background *gfx.Parallax
	layers     []Layer
	enemies    []Enemy
}

func NewLevel() Level {
	return Level{
		layers: []Layer{NewLayer(5, 100)},
	}
}

func (l *Level) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	for _, l := range l.layers {
		l.Draw(screen, op)
	}
	// TODO draw enemies
}
