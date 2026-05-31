package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var game *Game

func main() {
	initGame()

	eb.SetWindowSize(800, 800)
	eb.SetWindowTitle("Iron Express")

	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
