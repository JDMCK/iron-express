package input

import (
	"fmt"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action int // What should happen from that input?
type Source int // Where did that input come from?

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
	KeyboardKey Source = iota
	MouseButton
	GamepadButton

	sourceCount
)

type Mode struct {
	IsPressed  bool
	IsReleased bool
}

type Binding struct {
	source Source
	value  int // keys, gamepad buttons, and mouse buttons all use enums (ints)
}

type Mapping map[Action][]Binding

type System struct {
	gamepadId eb.GamepadID
	mapping   Mapping
}

func (s *System) GetAction(a Action) Mode {
	bindings := s.mapping[a] // we can assume all actions are in map
	isPressed := false
	isReleased := false
	// checks if any bindings are being pressed/released
	for _, b := range bindings {
		switch b.source {
		case KeyboardKey:
			isPressed = eb.IsKeyPressed(eb.Key(b.value)) || isPressed
			isReleased = inpututil.IsKeyJustReleased(eb.Key(b.value)) || isReleased
		case MouseButton:
			isPressed = eb.IsMouseButtonPressed(eb.MouseButton(b.value)) || isPressed
			isReleased = inpututil.IsMouseButtonJustReleased(eb.MouseButton(b.value)) || isReleased
		case GamepadButton:
			isPressed = eb.IsGamepadButtonPressed(s.gamepadId, eb.GamepadButton(b.value)) || isPressed
			isReleased = inpututil.IsGamepadButtonJustReleased(s.gamepadId, eb.GamepadButton(b.value)) || isReleased
		}
	}
	return Mode{
		IsPressed:  isPressed,
		IsReleased: isReleased,
	}
}

func NewBinding(s Source, v int) (Binding, error) {
	if s >= sourceCount {
		return Binding{}, fmt.Errorf("Source must be [0, 2], not %d", s)
	}
	return Binding{source: s, value: v}, nil
}

func NewSystem(m Mapping, gp eb.GamepadID) *System {
	return &System{
		mapping:   m,
		gamepadId: gp,
	}
}

func GetConnectedGamepadID() eb.GamepadID {
	// TODO find connected gamepads
	return eb.GamepadID(0)
}
