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
	Collider     *core.Collider
}

const playerWidth = 32
const playerHeight = 32

// Horizontal movement
const runAcceleration float64 = 2
const runDeceleration float64 = 1.5
const maxRunSpeed float64 = 7
const horizontalEpsilon float64 = 0.01

// Vertical movement
const jumpForce float64 = 5
const maxJumpSpeed float64 = 7
const maxFallSpeed float64 = 10
const gravityAcceleration float64 = 0.6

const TEMPGround float64 = 300

// Gun parameters
const gunPower float64 = 1

func NewPlayer() (*Player, error) {
	anims, err := config.LoadAnimationAtlas("player")
	if err != nil {
		return nil, err
	}

	pos := core.Vector2{X: 100.0, Y: 0.0}
	collider := core.NewCollider(pos, playerWidth, playerHeight)

	return &Player{
		position:   pos,
		state:      Idling,
		animations: anims,
		Collider:   collider,
	}, nil
}

func (p *Player) Update(g *Game) {
	switch {
	case g.Input.GetAction(input.Primary).IsPressed:
		p.state = Shooting
		p.Shoot()
	case p.velocity.Y < 0 && p.isGrounded == false:
		p.state = Jumping
	case p.velocity.Y > 0 && p.isGrounded == false:
		p.state = Falling
	case math.Abs(p.velocity.X) > 0:
		p.state = Running
	default:
		p.state = Idling
	}

	setPlayerAccel(p, g)
	setPlayerVelocity(p, g)
	movePlayer(p, g)

	p.animations[p.state].Update()
}

func (p *Player) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	p.Collider.Draw(screen)
	p.animations[p.state].Draw(screen, int(p.position.X), int(p.position.Y), p.facingRight)
}

func setPlayerAccel(p *Player, g *Game) {
	if g.Input.GetAction(input.Left).IsPressed {
		p.acceleration.X = -runAcceleration
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.acceleration.X = runAcceleration
	} else {
		p.acceleration.X = 0
	}

	// TODO: should set this somewhere else
	p.isGrounded = p.position.Y >= TEMPGround

	// Jump instantly
	if g.Input.GetAction(input.Jump).IsPressed && p.isGrounded {
		p.acceleration.Y = -jumpForce - gravityAcceleration
	}

	// Gravity always
	p.acceleration.Y += gravityAcceleration

	p.acceleration.Y = core.Clamp(-jumpForce, gravityAcceleration, p.acceleration.Y)

	fmt.Printf("Accel: %v %v\n", p.acceleration.X, p.acceleration.Y)
}

// Calculate and set the horizontal and vertical velocities
// for the Player. This is based on input presses (which sets accel), current
// velocity, the max velocity, and gravity.
//
// Note that the player only attempts to "move" i.e. change position in movePlayer.
func setPlayerVelocity(p *Player, g *Game) {
	// TODO: shouldn't set acceleration here, only influence it

	p.velocity.X += p.acceleration.X
	p.velocity.X /= runDeceleration

	p.velocity.X = core.Clamp(-maxRunSpeed, maxRunSpeed, p.velocity.X)

	// stop micro velocity
	if math.Abs(p.velocity.X) <= horizontalEpsilon {
		p.velocity.X = 0
	}

	p.velocity.Y += p.acceleration.Y
	p.velocity.Y = core.Clamp(-maxJumpSpeed, maxFallSpeed, p.velocity.Y)
	fmt.Printf("Velo: %v %v\n", p.velocity.X, p.velocity.Y)
}

func movePlayer(p *Player, g *Game) {
	newPosition := core.Vector2{X: p.velocity.X + p.Collider.Position.X, Y: p.velocity.Y + p.Collider.Position.Y}
	levelLayers := g.GetCurrLevel().layers

	newPosCol := core.NewCollider(newPosition, playerWidth, playerHeight)

	for _, layer := range levelLayers {
		dir, amt := core.IntersectAABB(newPosCol, layer.Collider)

		switch dir {
		case core.XCol:
			newPosition.X += amt
		case core.YCol:
			newPosition.Y += amt
		}
	}

	p.position = newPosition
	p.Collider.Position = newPosition

	// TEMPORARY: ensure player doesn't fall out of the world
	p.Collider.Position.Y = core.Clamp(0, TEMPGround, p.Collider.Position.Y)
	p.position.Y = core.Clamp(0, TEMPGround, p.Collider.Position.Y)
}

func (p *Player) Shoot() {
	// Curr: Influence accel
	// Set a shooting cooldown
	// Eventually: spawn a Bullet

	cursorX, cursorY := eb.CursorPosition()
	aimDir := core.VectorNormalize(core.Vector2{
		X: float64(cursorX) - p.position.X,
		Y: float64(cursorY) - p.position.Y,
	})

	fmt.Println(aimDir)
	p.acceleration.X -= aimDir.X * gunPower
	p.acceleration.Y += aimDir.Y * gunPower

	p.facingRight = aimDir.X >= 0
}
