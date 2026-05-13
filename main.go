package main

import (
	"iron-express/gfx"
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var explosion gfx.Clip

func init() {
	var err error
	img, _, err := ebitenutil.NewImageFromFile("assets/animations/explosion.png")
	if err != nil {
		log.Fatal(err)
	}
	atlas := gfx.NewAtlas(img, 14, 14)
	explosion = gfx.NewClip(atlas, 10, true)
}

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(eb.KeySpace) {
		explosion.TogglePause()
	}

	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	explosion.DrawAndUpdate(screen, 100, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	eb.SetWindowSize(640, 480)
	eb.SetWindowTitle("~Gome~")
	eb.SetWindowResizingMode(eb.WindowResizingModeEnabled)

	if err := eb.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
