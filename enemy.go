package main

import (
	"iron-express/config"
	"iron-express/core"
	"iron-express/gfx"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	position     core.Vector2
	velocity     core.Vector2
	acceleration core.Vector2
	state        string
	animations   gfx.AnimationMap
	facingRight  bool
	isGrounded   bool
	Collider     *core.Collider
}

func NewEnemy() (*Enemy, error) {
	anims, err := config.LoadAnimationAtlas("player")
	if err != nil {
		return nil, err
	}

	pos := core.Vector2{X: 100.0, Y: 0.0}
	collider := core.NewCollider(pos, playerWidth, playerHeight)

	return &Enemy{
		position:   pos,
		state:      Idling,
		animations: anims,
		Collider:   collider,
	}, nil
}

func (e *Enemy) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	e.Collider.Draw(screen)
	e.animations[e.state].Draw(screen, int(e.position.X), int(e.position.Y), e.facingRight, op)
}
