package main

import (
	"image/color"
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
	g.player.Update(g)
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{10, 180, 255, 255})
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}

func main() {
	eb.SetWindowSize(800, 800)
	eb.SetWindowTitle("Iron Express")

	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
