package config

import (
	"fmt"
	"iron-express/input"
	"os"
	"strconv"
	"strings"
)

const inputConfigPath = "config/inputs.config"

func LoadInput() (*input.System, error) {
	data, err := os.ReadFile(inputConfigPath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	mapping := make(input.Mapping)

	for _, line := range lines {
		action, bindings, err := parseInput(line)
		if err != nil {
			continue // likely a comment or blank line
		}
		mapping[input.Action(action)] = bindings
	}

	return input.NewSystem(mapping, input.GetConnectedGamepadID()), nil
}

func parseInput(line string) (input.Action, []input.Binding, error) {
	var (
		action input.Action
		key    int
		mouse  int
		button int
	)

	parts := strings.Split(line, " ")
	if len(parts) != 4 {
		return 0, nil, fmt.Errorf("Failed to parse input: %s", line)
	}
	// action
	actionPart := parts[0]
	switch strings.ToLower(actionPart) {
	case "up":
		action = input.Up
	case "down":
		action = input.Down
	case "left":
		action = input.Left
	case "right":
		action = input.Right
	case "jump":
		action = input.Jump
	case "primary":
		action = input.Primary
	case "secondary":
		action = input.Secondary
	case "interact":
		action = input.Interact
	case "escape":
		action = input.Escape
	default:
		return 0, nil, fmt.Errorf("Input %s not recognized", actionPart)
	}
	bindings := make([]input.Binding, 0)

	// keyboard
	key, err := strconv.Atoi(parts[1])
	if err == nil {
		keyBind, err := input.NewBinding(input.KeyboardKey, key)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, keyBind)
	}
	// mouse
	mouse, err = strconv.Atoi(parts[2])
	if err == nil {
		mouseBind, err := input.NewBinding(input.MouseButton, mouse)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, mouseBind)
	}
	// gamepad
	button, err = strconv.Atoi(parts[3])
	if err == nil {
		gamepadBind, err := input.NewBinding(input.GamepadButton, button)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, gamepadBind)
	}

	return action, bindings, nil
}
