package input

type Action int    // What should happen from that input?
type InputMean int // Where did that input come from?
type InputMode int // what did that input just do?

const (
	Up Action = iota
	Down
	Left
	Right
	Jump
	Primary
	Secondary
	Interact
	Escape
)

const (
	KeyboardKey InputMean = iota
	MouseButton
	GamepadButton
)

const (
	isUp InputMode = iota
	isDown
	isPressed
	isReleased
)

type InputSource struct {
	mean      InputMean
	gamepadId int
}
