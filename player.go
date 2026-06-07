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
	accel        core.Vector2
	state        string
	animations   gfx.AnimationMap
	facingRight  bool
	isGrounded   bool
	shotCooldown int
	Collider     *core.Collider
}

const playerWidth = 32
const playerHeight = 32

// Horizontal movement
const runAcceleration float64 = 3
const runDeceleration float64 = 1.3
const maxRunSpeed float64 = 7
const horizontalEpsilon float64 = 0.01

// Vertical movement
const jumpAccel float64 = 5
const maxJumpSpeed float64 = 7
const maxFallSpeed float64 = 10
const gravityAccel float64 = 0.9

const TEMPGround float64 = 300

// Gun parameters
const gunPower float64 = 20
const gunDelay int = 30

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
	case g.Input.GetAction(input.Primary).IsPressed || p.shotCooldown > 0:
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

	setPlayerAccel(p, g)

	// Shooting affects acceleration
	if p.state == Shooting {
		p.Shoot()
	}

	setPlayerVelocity(p, g)
	movePlayer(p, g)

	// TODO: don't change facing param while shooting
	if p.shotCooldown == 0 {
		switch {
		case p.velocity.X > 0:
			p.facingRight = true
		case p.velocity.X < 0:
			p.facingRight = false
		}
	}

	p.animations[p.state].Update()
}

func (p *Player) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	p.Collider.Draw(screen)
	p.animations[p.state].Draw(screen, int(p.position.X), int(p.position.Y), p.facingRight)
}

func setPlayerAccel(p *Player, g *Game) {
	if g.Input.GetAction(input.Left).IsPressed {
		p.accel.X = -runAcceleration
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.accel.X = runAcceleration
	} else {
		p.accel.X = 0
	}

	// TODO: should set this somewhere else
	p.isGrounded = p.position.Y >= TEMPGround

	if g.Input.GetAction(input.Jump).IsPressed && p.isGrounded {
		p.accel.Y = -jumpAccel - gravityAccel
	}

	// Gravity always
	p.accel.Y += gravityAccel

	p.accel.Y = core.Clamp(-jumpAccel, gravityAccel, p.accel.Y)
}

// Calculate and set the horizontal and vertical velocities
// for the Player. This is based on input presses (which sets accel), current
// velocity, the max velocity, and gravity.
//
// Note that the player only attempts to "move" i.e. change position in movePlayer.
func setPlayerVelocity(p *Player, g *Game) {
	p.velocity.X += p.accel.X
	p.velocity.X /= runDeceleration

	p.velocity.X = core.Clamp(-maxRunSpeed, maxRunSpeed, p.velocity.X)

	// stop micro velocity
	if math.Abs(p.velocity.X) <= horizontalEpsilon {
		p.velocity.X = 0
	}

	p.velocity.Y += p.accel.Y
	p.velocity.Y = core.Clamp(-maxJumpSpeed, maxFallSpeed, p.velocity.Y)
}

func movePlayer(p *Player, g *Game) {
	p.Collider.Position = core.Vector2{
		X: p.velocity.X + p.Collider.Position.X,
		Y: p.velocity.Y + p.Collider.Position.Y,
	}

	levelLayers := g.GetCurrLevel().layers

	for _, layer := range levelLayers {
		dir, amt := core.IntersectAABB(p.Collider, layer.Collider)

		switch dir {
		case core.XCol:
			p.Collider.Position.X += amt
		case core.YCol:
			p.Collider.Position.Y += amt
		}
	}

	p.position = p.Collider.Position

	// TEMPORARY: ensure player doesn't fall out of the world
	p.Collider.Position.Y = core.Clamp(0, TEMPGround, p.Collider.Position.Y)
	p.position.Y = core.Clamp(0, TEMPGround, p.Collider.Position.Y)
}

// TODO: spawn a bullet in this function
func (p *Player) Shoot() {
	if p.shotCooldown > 0 {
		p.shotCooldown -= 1
		return
	}

	cursorX, cursorY := eb.CursorPosition()
	aimDir := core.VectorNormalize(core.Vector2{
		X: float64(cursorX) - p.position.X,
		Y: float64(cursorY) - p.position.Y,
	})

	p.accel.X -= aimDir.X * gunPower
	p.accel.Y -= aimDir.Y * gunPower

	p.shotCooldown = gunDelay

	p.facingRight = aimDir.X >= 0
}
