package main

import (
	"iron-express/config"
	"iron-express/core"
	"iron-express/gfx"
	"iron-express/input"
	"math"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	Idling   = "idle"
	Running  = "run"
	Jumping  = "jump" // moving up
	Falling  = "fall" // moving down
	Shooting = "shoot"
)

type Player struct {
	position     core.Vector2
	velocity     core.Vector2
	acceleration core.Vector2
	state        string
	animations   gfx.AnimationMap
	facingRight  bool
	isGrounded   bool
	Collider     core.Collider
}

const playerWidth = 32
const playerHeight = 32

func NewPlayer() (*Player, error) {
	anims, err := config.LoadAnimationAtlas("player")
	if err != nil {
		return nil, err
	}

	pos := core.Vector2{X: 0.0, Y: 0.0}
	collider := core.NewCollider(pos, playerWidth, playerHeight)

	return &Player{
		state:      Idling,
		animations: anims,
		Collider:   collider,
	}, nil
}

func (p *Player) Update(g *Game) {
	move(p, g)
	p.animations[p.state].Update()

	switch {
	case g.Input.GetAction(input.Primary).IsPressed:
		p.state = Shooting
	case p.velocity.Y < 0 && p.isGrounded == false:
		p.state = Jumping
	case p.velocity.Y > 0 && p.isGrounded == false:
		p.state = Falling
	case math.Abs(p.velocity.X) > 0:
		p.state = Running
	default:
		p.state = Idling
	}
}

func (p *Player) Draw(screen *eb.Image) {
	p.Collider.Draw(screen)
	p.animations[p.state].Draw(screen, int(p.position.X), int(p.position.Y), p.facingRight)
}

// Horizontal
const runAcceleration float64 = 2
const runDeceleration float64 = 2 // friction
const maxRunSpeed float64 = 2
const horizontalEpsilon float64 = 0.01

// Vertical
const jumpSpeed float64 = 7
const maxJumpSpeed float64 = 7
const maxFallSpeed float64 = 10
const gravityAcceleration float64 = 0.5

const TEMPGround float64 = 100

func move(p *Player, g *Game) {
	// ----- Horizontal -----
	if g.Input.GetAction(input.Left).IsPressed {
		p.acceleration.X = -runAcceleration
		p.facingRight = false
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.acceleration.X = runAcceleration
		p.facingRight = true
	} else {
		p.acceleration.X = 0

		// apply friction when moving (but no key presses)
		if p.velocity.X != 0 {
			p.velocity.X /= runDeceleration
		}
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
		p.acceleration.Y = 0
		p.velocity.Y = -jumpSpeed
	}

	// apply gravity
	if p.isGrounded == false && p.position.Y < 100 {
		p.acceleration.Y = gravityAcceleration
	}

	p.velocity.Y = core.Clamp(-maxJumpSpeed, maxFallSpeed, p.velocity.Y)

	p.velocity.Y = p.acceleration.Y + p.velocity.Y
	p.position.Y = p.velocity.Y + p.position.Y

	// snap to floor (temp)
	p.position.Y = core.Clamp(0, 100, p.position.Y)

	// Update collision box
	p.Collider.Position = p.position
}
