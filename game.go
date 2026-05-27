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

	player       *Player
	testCollider core.Collider
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

	// temporary
	testCollider := core.NewCollider(
		core.Vector2{X: 0, Y: 90},
		5,
		5,
	)

	return &Game{
		Input:        *input,
		player:       player,
		testCollider: testCollider,
	}, nil
}

func initGame() {
	var err error
	game, err = NewGame()
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.player.Update(g)
	colliding := core.IsCollided(g.player.Collider, g.testCollider)
	if colliding {
		fmt.Println("collided!")
	}
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.player.Draw(screen)
	g.testCollider.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}
