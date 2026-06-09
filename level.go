package main

import (
	"iron-express/config"
	"iron-express/gfx"
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	background *gfx.Parallax
	layers     []Layer
	enemies    []Enemy
}

func NewLevel() Level {
	backAtlas, err := config.LoadAtlas("background")
	if err != nil {
		log.Fatal("Failed to load background atlas.")
	}
	parallax := gfx.NewParallax(backAtlas, []float64{0.99, 0.95, 0.9, 0.8, 0.4})
	return Level{
		background: parallax,
		layers:     []Layer{},
	}
}

func (l *Level) Update() {
	l.background.Update()
}

func (l *Level) Draw(screen *eb.Image, x, y float64, op *eb.DrawImageOptions) {
	l.background.Draw(screen, x, y, op)

	for _, l := range l.layers {
		l.Draw(screen, op)
	}
	// TODO draw enemies
}
