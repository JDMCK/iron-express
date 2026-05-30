package main

import (
	"iron-express/config"
	"iron-express/input"
)

type Game struct {
	paused bool
	debug  bool
	Input  input.System

	player *Player
	level  *Level
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
	level := NewLevel()
	return &Game{
		Input:  *input,
		player: player,
		level:  &level,
	}, nil
}
