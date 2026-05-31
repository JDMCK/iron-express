package main

import (
	"fmt"
	"image/color"
	"iron-express/config"
	"iron-express/core"
	"iron-express/input"
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	paused bool
	debug  bool
	Input  input.System

	player    *Player
	levels    []Level
	currLevel int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}

func NewGame() (*Game, error) {
	input, err := config.LoadInput()
	if err != nil {
		return nil, err
	}
	player, err := NewPlayer()
	if err != nil {
		return nil, err
	}

	// Make the initial level
	levels := make([]Level, 0, 1)
	level := NewLevel()
	levels = append(levels, level)

	return &Game{
		Input:  *input,
		player: player,
		levels: levels,
	}, nil
}

func initGame() {
	var err error
	game, err = NewGame()
	if err != nil {
		log.Fatal(err)
	}
}

var frame = 0

func (g *Game) Update() error {
	g.player.Update(g)

	level := g.GetCurrLevel()
	for _, layer := range level.layers {
		if core.IsCollided(g.player.Collider, layer.Collider) {
			// do something about this
			fmt.Printf("collided on frame %d\n", frame)
		}
	}

	frame += 1
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.GetCurrLevel().Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) GetCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
