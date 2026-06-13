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
	moveState    string
	animState    string
	animations   gfx.AnimationMap
	facingRight  bool
	isGrounded   bool // TODO: isGrounded should get replaced by moveState (not jumping or falling)
	shotCooldown int
	Collider     *core.Collider
}

const playerWidth = 32
const playerHeight = 32

// Horizontal movement
const runAcceleration float64 = 0.5
const runDeceleration float64 = 0.5
const maxRunSpeed float64 = 7

// Vertical movement
const jumpAccel float64 = 4
const maxJumpSpeed float64 = 10
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
		moveState:  Idling,
		animations: anims,
		Collider:   collider,
	}, nil
}

func setStates(p *Player) {
	if p.velocity.Y < 0 && p.isGrounded == false {
		p.moveState = Jumping
		p.animState = Jumping
	} else if p.velocity.Y > 0 && p.isGrounded == false {
		p.moveState = Falling
		p.animState = Falling
	} else if math.Abs(p.velocity.X) > 0 {
		p.moveState = Running
		p.animState = Running
	} else {
		p.moveState = Idling
		p.animState = Idling
	}
}

func (p *Player) Update(g *Game) {
	// Determine what state the player is in, e.g. falling, jumping, etc.
	setStates(p)

	// Accelerate the player based on their state
	setPlayerAccel(p, g)

	// Handle a shot from the player's gun
	if g.Input.GetAction(input.Primary).IsPressed == true || p.shotCooldown > 0 {
		p.animState = Shooting
		p.Shoot()
	}

	// Fix player's velocity based on their acceleration. Handles clamping.
	setPlayerVelocity(p, g)

	// Attempt to move player to their new position.
	movePlayer(p, g)

	if p.shotCooldown == 0 {
		switch {
		case p.velocity.X > 0:
			p.facingRight = true
		case p.velocity.X < 0:
			p.facingRight = false
		}
	}

	p.animations[p.animState].Update()
}

func (p *Player) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	p.Collider.Draw(screen)
	p.animations[p.moveState].Draw(screen, int(p.position.X), int(p.position.Y), p.facingRight)
}

// Handles player acceleration from player input.
// Rule of thumb: we SET acceleration, and ADD to velocity.
func setPlayerAccel(p *Player, g *Game) {
	if g.Input.GetAction(input.Left).IsPressed {
		p.accel.X = -runAcceleration
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.accel.X = runAcceleration
	} else {
		p.accel.X = 0
	}

	if g.Input.GetAction(input.Jump).IsPressed && p.moveState != Falling && p.moveState != Jumping {
		// Perform a jump
		p.accel.Y = -jumpAccel - gravityAccel
	}
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

	// No player horizontal input
	if !g.Input.GetAction(input.Left).IsPressed && !g.Input.GetAction(input.Right).IsPressed {
		switch p.velocity.X > 0 {
		case true:
			p.velocity.X = math.Max(p.velocity.X-runDeceleration, 0)
		case false:
			p.velocity.X = math.Min(p.velocity.X+runDeceleration, 0)
		}
	}

	p.velocity.X = core.Clamp(-maxRunSpeed, maxRunSpeed, p.velocity.X)

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
	p.isGrounded = p.position.Y >= TEMPGround
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
