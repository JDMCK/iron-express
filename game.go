package main

import (
	"image/color"
	"iron-express/gui"
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
	Input  Input

	cam *Camera

	player *Player

	levels    []Level
	currLevel int
	enemies   []*Enemy

	gui []gui.Element
}

var frame = 0 // to keep track of the number of frames since the game started

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}

func NewGame() (*Game, error) {
	input, err := LoadInput()
	if err != nil {
		return nil, err
	}
	player, err := NewPlayer()
	if err != nil {
		return nil, err
	}

	// enemy, err := NewEnemy()
	// if err != nil {
	// 	return nil, err
	// }

	enemies := make([]*Enemy, 0, 10)
	// enemies = append(enemies, enemy)

	// Make the initial level
	levels := make([]Level, 0, 1)
	level := NewLevel("level00")
	levels = append(levels, level)

	// Make the GUI
	guiEls := make([]gui.Element, 0)

	cam := NewCamera(1)
	cam.CenterScreenOffset(WorldWidth-player.Collider.Width, 0)

	return &Game{
		Input:   *input,
		player:  player,
		levels:  levels,
		gui:     guiEls,
		cam:     cam,
		enemies: enemies,
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

	g.cam.SetFocusX(g.player.position.X, 1)

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

	for _, enemy := range g.enemies {
		enemy.Draw(screen, camOp)
	}
}

func (g *Game) GetCurrLevel() *Level {
	return &g.levels[g.currLevel]
}
