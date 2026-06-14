package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	position     Vector2
	velocity     Vector2
	acceleration Vector2
	state        string
	animations   AnimationMap
	facingRight  bool
	isGrounded   bool
	Collider     *Collider
}

func NewEnemy() (*Enemy, error) {
	anims, err := LoadAnimationAtlas("player")
	if err != nil {
		return nil, err
	}

	pos := Vector2{X: 100.0, Y: 0.0}
	collider := NewCollider(pos, playerWidth, playerHeight)

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
