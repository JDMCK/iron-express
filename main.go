package main

import (
	"log"

	"net/http"
	_ "net/http/pprof"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var game *Game

func main() {
	// Run profiling server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	initGame()

	eb.SetWindowSize(800, 800)
	eb.SetWindowTitle("Iron Express")

	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
