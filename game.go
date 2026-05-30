package main

import (
	"iron-express/config"
	"iron-express/input"
)

type Game struct {
	paused bool
	debug  bool
	Input  input.System

	player    *Player
	levels    []Level
	currLevel int
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
