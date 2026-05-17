package main

import (
	"image/color"
	"iron-express/config"
	"iron-express/gfx"
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var animations gfx.AnimationMap
var animIndex int = 0
var currentAnim *gfx.Animation
var name string

func init() {
	anims, err := config.LoadAnimationAtlas("player")
	if err != nil {
		log.Fatal(err)
	}
	animations = anims
	currentAnim = animations["idle"]
}

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(eb.KeySpace) {
		animIndex = (animIndex + 1) % 4
	}
	names := []string{"run", "idle", "fall", "jump"}
	name = names[animIndex]
	currentAnim = animations[name]

	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	screen.Fill(color.RGBA{0, 100, 200, 0})
	ebitenutil.DebugPrint(screen, name)
	currentAnim.DrawAndUpdate(screen, 50, 50)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}

func main() {
	eb.SetWindowSize(400, 400)
	eb.SetWindowTitle("Iron Express")
	eb.SetWindowResizingMode(eb.WindowResizingModeEnabled)

	if err := eb.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
