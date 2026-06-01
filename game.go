package main

import (
	"image/color"
	"iron-express/config"
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

var frame = 0 // to keep track of the number of frames since the game started

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

	// Make the GUI
	guiEls := make([]gui.Element, 0)

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

func (g *Game) Update() error {
	g.player.Update(g)

	for _, e := range g.gui {
		e.Update()
	}

	frame += 1
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.getCurrLevel().Draw(screen)
	g.player.Draw(screen)
	for _, e := range g.gui {
		e.Draw(screen)
	}
}

func (g *Game) getCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
