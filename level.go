package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	background *Parallax
	layers     []*Layer
	enemies    []Enemy
}

func NewLevel(levelConfig string) Level {
	backAtlas, err := LoadAtlas("background")
	if err != nil {
		log.Fatal("Failed to load background atlas.")
	}
	parallax := NewParallax(backAtlas, []float64{1, 0.95, 0.9, 0.8, 0.4})
	layers := LoadLevelLayers(levelConfig)
	return Level{
		background: parallax,
		layers:     layers,
	}
}

func (l *Level) Update() {
	l.background.Update()
}

func (l *Level) Draw(screen *eb.Image, x, y float64, op *eb.DrawImageOptions) {
	l.background.Draw(screen, x, op)

	for _, l := range l.layers {
		l.Draw(screen, op)
	}
	// TODO draw enemies
}
