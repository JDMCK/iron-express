package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var game *Game

func init() {
	var err error
	game, err = NewGame()
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}

func main() {
	eb.SetWindowSize(400, 400)
	eb.SetWindowTitle("Iron Express")
	// eb.SetWindowResizingMode(eb.WindowResizingModeEnabled)

	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
