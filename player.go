package main

import (
	"iron-express/config"
	"iron-express/core"
	"iron-express/gfx"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	Idle    = "idle"
	Running = "run"
	Jumping = "jump" // moving up
	Falling = "fall" // moving down
)

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
		state:      Idle,
		animations: anims,
	}, nil
}

func (p *Player) Update(g *Game) {

}

func (p *Player) Draw(screen *eb.Image) {
	p.animations[p.state].DrawAndUpdate(screen, int(p.position.X), int(p.position.Y))
}
