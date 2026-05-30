package main

import (
	"image/color"
	"iron-express/config"
	"iron-express/input"

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

func (g *Game) Update() error {
	g.player.Update(g)
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.player.Draw(screen)
	g.GetCurrLevel().Draw(screen)
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
	levels := make([]Level, 0, 1)
	level := NewLevel()
	levels = append(levels, level)
	return &Game{
		Input:  *input,
		player: player,
		levels: levels,
	}, nil
}

func (g *Game) GetCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
