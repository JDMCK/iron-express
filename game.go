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
	return &Game{
		Input:  *input,
		player: player,
	}, nil
}
