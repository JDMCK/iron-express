package main

import (
	"fmt"
	"image/color"
	"iron-express/config"
	"iron-express/core"
	"iron-express/gui"
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

	gui []gui.Element
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
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

	guiEls := make([]gui.Element, 0)
	guiEls = append(guiEls, gui.NewBasicButton("Save", 50, 50, gui.Large, gui.Primary, func() { fmt.Println("Saved") }))
	guiEls = append(guiEls, gui.NewBasicButton("Cancel", 50, 100, gui.Small, gui.Secondary, func() { fmt.Println("Cancel") }))
	guiEls = append(guiEls, gui.NewBasicButton("Delete", 50, 150, gui.Medium, gui.Danger, func() { fmt.Println("Delete") }))
	guiEls = append(guiEls, gui.NewNumberPicker(1, 5, 0, 50, 200))
	guiEls = append(guiEls, gui.NewCheckbox(50, 250))

	return &Game{
		Input:  *input,
		player: player,
		levels: levels,
		gui:    guiEls,
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

	for _, e := range g.gui {
		e.Update()
	}

	handleCollisions(g)

	frame += 1
	return nil
}

func handleCollisions(g *Game) {
	// Collision detection
	level := g.GetCurrLevel()
	for _, layer := range level.layers {
		if dir, amt := core.IntersectAABB(g.player.Collider, layer.Collider); dir != core.None {
			fmt.Println("player collided with layer by %d amount", amt)
			// TODO: player function to adjust its position
		}
	}
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.GetCurrLevel().Draw(screen)
	g.player.Draw(screen)
	for _, e := range g.gui {
		e.Draw(screen)
	}
}

func (g *Game) GetCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
