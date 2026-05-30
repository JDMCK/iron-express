package main

import (
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

func initGame() *Game {
	var err error
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	return game
}

func main() {
	var game *Game = initGame()
	eb.SetWindowSize(800, 800)
	eb.SetWindowTitle("Iron Express")

	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
