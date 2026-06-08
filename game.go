package main

import (
	"image/color"
	"iron-express/config"
	"iron-express/gui"
	"iron-express/input"
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var WindowWidth = 1280
var WindowHeight = 720
var WorldWidth = 320
var WorldHeight = 180

type Game struct {
	paused bool
	debug  bool
	Input  input.System

	cam *Camera

	player    *Player
	levels    []Level
	currLevel int

	gui []gui.Element
}

var frame = 0 // to keep track of the number of frames since the game started

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
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

	cam := NewCamera(1)
	cam.CenterScreenOffset(WorldWidth-player.Collider.Width, WorldHeight*2-player.Collider.Height*2)

	return &Game{
		Input:  *input,
		player: player,
		levels: levels,
		gui:    guiEls,
		cam:    cam,
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
	g.levels[g.currLevel].Update()

	g.cam.SetFocus(g.player.position.X, g.player.position.Y, 1)

	frame += 1
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})

	camOp := g.cam.DrawOptions()

	g.GetCurrLevel().Draw(screen, g.cam.focusX, g.cam.focusY, camOp)

	g.player.Draw(screen, camOp, g.debug)

	for _, e := range g.gui {
		e.Draw(screen)
	}
}

func (g *Game) GetCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
