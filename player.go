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
	isGrounded   bool
	shotCooldown int
	Collider     *core.Collider
}

const playerWidth = 32
const playerHeight = 32

// Horizontal movement
const runAcceleration float64 = 1
const runDeceleration float64 = 1
const maxRunSpeed float64 = 4

// Vertical movement
const jumpAccel float64 = 4
const maxJumpSpeed float64 = 8
const maxFallSpeed float64 = 10
const gravity float64 = 1

const TEMPGround float64 = 300

// Gun parameters
const gunPowerX float64 = 25
const gunPowerY float64 = 4
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
		animState:  Idling,
		animations: anims,
		Collider:   collider,
	}, nil
}

func (p *Player) Update(g *Game) {
	setMoveState(p)

	setPlayerAccel(p, g)

	// Movement changes from shooting should come after player movement (?)
	if g.Input.GetAction(input.Primary).IsPressed || p.shotCooldown > 0 {
		p.Shoot()
	}

	setPlayerVelocity(p, g)

	movePlayer(p, g)

	if g.Input.GetAction(input.Right).IsPressed {
		p.facingRight = true
	} else if g.Input.GetAction(input.Left).IsPressed {
		p.facingRight = false
	}

	setAnimState(p)
	p.animations[p.animState].Update()
}

func (p *Player) Draw(screen *eb.Image, op *eb.DrawImageOptions) {
	p.Collider.Draw(screen)
	p.animations[p.animState].Draw(screen, int(p.position.X), int(p.position.Y), p.facingRight)
}

// Determine what movement state the player is in, e.g. falling, jumping, etc.
func setMoveState(p *Player) {
	if p.velocity.Y < 0 && !p.isGrounded { // BUG: at the top of a jump, Y velocity == 0, but we are still jumping
		p.moveState = Jumping
	} else if p.velocity.Y > 0 && !p.isGrounded {
		p.moveState = Falling
	} else if math.Abs(p.velocity.X) > 0 {
		p.moveState = Running
	} else {
		p.moveState = Idling
	}
}

func setAnimState(p *Player) {
	if p.shotCooldown > 0 {
		p.animState = Shooting
	} else if p.velocity.Y < 0 && p.isGrounded == false {
		p.animState = Jumping
	} else if p.velocity.Y > 0 && p.isGrounded == false {
		p.animState = Falling
	} else if math.Abs(p.velocity.X) > 0 {
		p.animState = Running
	} else {
		p.animState = Idling
	}
}

// Handles player acceleration from player input.
func setPlayerAccel(p *Player, g *Game) {
	if g.Input.GetAction(input.Left).IsPressed {
		p.accel.X = -runAcceleration
	} else if g.Input.GetAction(input.Right).IsPressed {
		p.accel.X = runAcceleration
	} else {
		p.accel.X = 0
	}

	if g.Input.GetAction(input.Jump).IsPressed && p.isGrounded {
		// Perform a jump
		p.isGrounded = false
		p.accel.Y = -jumpAccel - gravity
	}
	p.accel.Y += gravity

	p.accel.Y = core.Clamp(-jumpAccel, gravity, p.accel.Y)
}

// Slow the player down by 'runDeceleration' amount
func applyPlayerDecel(p *Player) {
	switch p.velocity.X > 0 {
	case true:
		p.velocity.X = math.Max(p.velocity.X-runDeceleration, 0)
	case false:
		p.velocity.X = math.Min(p.velocity.X+runDeceleration, 0)
	}
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
		applyPlayerDecel(p)
	}

	p.velocity.X = core.Clamp(-maxRunSpeed, maxRunSpeed, p.velocity.X)

	p.velocity.Y += p.accel.Y
	p.velocity.Y = core.Clamp(-maxJumpSpeed, maxFallSpeed, p.velocity.Y)
}

// Attempt to move player to their new position. Handles collision.
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
	if p.isGrounded {
		p.velocity.Y = 0
	}
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

	p.velocity.X -= aimDir.X * gunPowerX
	p.velocity.Y -= aimDir.Y * gunPowerY

	p.shotCooldown = gunDelay

	p.facingRight = aimDir.X >= 0
}
