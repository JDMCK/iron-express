package main

import (
	"fmt"
	"iron-express/config"
	"iron-express/core"
	"iron-express/gfx"
	"iron-express/input"
	"math"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	Idling  = "idle"
	Running = "run"
	Jumping = "jump" // moving up
	Falling = "fall" // moving down
)

type Player struct {
	position     core.Vector2
	velocity     core.Vector2
	acceleration core.Vector2
	state        string
	animations   gfx.AnimationMap
	isGrounded   bool
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
	move(p, g)
}

func (p *Player) Draw(screen *eb.Image) {
	p.animations[p.state].DrawAndUpdate(screen, int(p.position.X), int(p.position.Y))
}

// Horizontal
const runAcceleration float64 = 2
const runDeceleration float64 = 2 // friction
const maxRunSpeed float64 = 2
const horizontalEpsilon float64 = 0.01

// Vertical
const jumpSpeed float64 = 3
const maxJumpSpeed float64 = 4
const maxFallSpeed float64 = 2
const gravityAcceleration float64 = 2

const TEMPGround float64 = 100

func move(p *Player, g *Game) {
	// ----- Horizontal -----
	if g.Input.GetAction(input.Left).IsPressed {
		p.acceleration.X = -runAcceleration
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.acceleration.X = runAcceleration
	} else {
		p.acceleration.X = 0

		// apply friction when moving (but no key presses)
		if p.velocity.X != 0 {
			p.velocity.X /= runDeceleration
		}
	}

	if g.Input.GetAction(input.Left).IsPressed {
		fmt.Println("Left")
	}
	if g.Input.GetAction(input.Jump).IsPressed {
		fmt.Println("jump")
	}
	p.velocity.X = p.acceleration.X + p.velocity.X
	p.position.X = p.velocity.X + p.position.X

	// apply max velocity
	p.velocity.X = core.Clamp(-maxRunSpeed, maxRunSpeed, p.velocity.X)

	// stop micro velocity
	if math.Abs(p.velocity.X) <= horizontalEpsilon {
		p.velocity.X = 0
	}

	// ----- Vertical -----
	p.isGrounded = p.position.Y >= 100
	if g.Input.GetAction(input.Jump).IsPressed && p.isGrounded {
		p.velocity.Y = -jumpSpeed
	} else {
		p.velocity.Y = 0
	}

	// apply gravity
	if !p.isGrounded && p.position.Y < 100 {
		p.acceleration.Y = gravityAcceleration
	}

	p.velocity.Y = core.Clamp(-maxJumpSpeed, maxFallSpeed, p.velocity.Y)

	// snap to floor (temp)
	if p.position.Y >= 100 {
		p.acceleration.Y = 0
		p.position.Y = 100
	}

	p.velocity.Y = p.acceleration.Y + p.velocity.Y
	p.position.Y = p.velocity.Y + p.position.Y

	// fmt.Println("acc: ", p.acceleration, "vel: ", p.velocity, "pos: ", p.position, "grounded: ", p.isGrounded)
}
