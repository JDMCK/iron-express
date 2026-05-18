package main

import (
	"fmt"
	"iron-express/config"
	"iron-express/core"
	"iron-express/gfx"
	"iron-express/input"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	Idling  = "idle"
	Running = "run"
	Jumping = "jump" // moving up
	Falling = "fall" // moving down
)

const runningSpeed float64 = 5.0

type Player struct {
	position   core.Vector2
	velocity   core.Vector2
	state      string
	animations gfx.AnimationMap
}

func NewPlayer() (*Player, error) {
	anims, err := config.LoadAnimationAtlas("player")
	if err != nil {
		return nil, err
	}
	return &Player{
		state:      Idling,
		animations: anims,
	}, nil
}

func (p *Player) Update(g *Game) {
	if g.Input.GetAction(input.Left).IsPressed {
		fmt.Println("RIGHT IS PRESSED")
		p.velocity.X = -runningSpeed
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.velocity.X = runningSpeed
	} else {
		p.velocity = core.Vector2{}
	}

	p.position = core.Add(p.position, p.velocity)
}

func (p *Player) Draw(screen *eb.Image) {
	p.animations[p.state].DrawAndUpdate(screen, int(p.position.X), int(p.position.Y))
}
