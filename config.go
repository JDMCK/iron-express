package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	eb "github.com/hajimehoshi/ebiten/v2"
	ebutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const inputConfigPath = "config/inputs.config"
const mapConfigPath = "config/levels/%s.config.map"

func LoadInput() (*Input, error) {
	data, err := os.ReadFile(inputConfigPath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	mapping := make(Mapping)

	for _, line := range lines {
		action, bindings, err := parseInput(line)
		if err != nil {
			continue // likely a comment or blank line
		}
		mapping[Action(action)] = bindings
	}

	return NewSystem(mapping, GetConnectedGamepadID()), nil
}

func parseInput(line string) (Action, []Binding, error) {
	var (
		action Action
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
		action = Up
	case "down":
		action = Down
	case "left":
		action = Left
	case "right":
		action = Right
	case "jump":
		action = Jump
	case "primary":
		action = Primary
	case "secondary":
		action = Secondary
	case "interact":
		action = Interact
	case "escape":
		action = Escape
	default:
		return 0, nil, fmt.Errorf("Input %s not recognized", actionPart)
	}
	bindings := make([]Binding, 0)

	// keyboard
	key, err := strconv.Atoi(parts[1])
	if err == nil {
		keyBind, err := NewBinding(KeyboardKey, key)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, keyBind)
	}
	// mouse
	mouse, err = strconv.Atoi(parts[2])
	if err == nil {
		mouseBind, err := NewBinding(MouseButton, mouse)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, mouseBind)
	}
	// gamepad
	button, err = strconv.Atoi(parts[3])
	if err == nil {
		gamepadBind, err := NewBinding(GamepadButton, button)
		if err != nil {
			return 0, nil, err
		}
		bindings = append(bindings, gamepadBind)
	}

	return action, bindings, nil
}

const atlasConfigPath = "config/sprites/%s.atlas.config"

func LoadAtlas(name string) (*Atlas, error) {
	data, err := os.ReadFile(fmt.Sprintf(atlasConfigPath, strings.ToLower(name)))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var (
		frameWidth  int
		frameHeight int
		rows        int
		cols        int
		img         *eb.Image
	)

	for _, line := range lines {
		k, v, err := ParseKV(line)
		if err != nil {
			continue // likely a comment or blank line
		}
		switch k {
		case "atlas_path":
			img, _, err = ebutil.NewImageFromFile(v)
			if err != nil {
				return nil, err
			}
		case "frame_width":
			frameWidth, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		case "frame_height":
			frameHeight, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		case "rows":
			rows, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		case "cols":
			cols, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return NewAtlas(img, rows, cols, frameWidth, frameHeight), nil
}

// func LoadLevelAtlas(path string) (*Level, error) {
// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	lines := strings.SplitSeq(string(data), "\n")

// 	layers := make([]Layer, 0, 10)
// 	var (
// 		atlasPath  string
// 		tileWidth  int
// 		tileHeight int
// 		mapWidth   int
// 		mapHeight  int
// 	)

// 	for line := range lines {
// 		k, v, err := ParseKV(line)
// 		if err != nil {
// 			continue // likely a comment or blank line
// 		}
// 		switch k {
// 		case "atlas_path":
// 			atlasPath = v
// 		case "tile_width":
// 			tileWidth, _ = strconv.Atoi(v)
// 		case "tile_height":
// 			tileHeight, _ = strconv.Atoi(v)
// 		case "map_width":
// 			mapWidth, _ = strconv.Atoi(v)
// 		case "map_height":
// 			mapHeight, _ = strconv.Atoi(v)
// 		}

// 		if strings.HasPrefix(k, "layer_") {
// 			var layer int
// 			fmt.Sscanf(k, "layer_%d", &layer)
// 			ImportedLayerIndices = append(ImportedLayerIndices, parseLayer(v))
// 		}
// 	}
// }

func parseLayer(data string, tileCount int) []int {
	parts := strings.Split(data, " ")
	indices := make([]int, 0, tileCount)
	for _, p := range parts {
		atlasIndex, count, _ := strings.Cut(p, "-")
		if count == "" {
			log.Fatal("Failed to parse layer.")
		}
		countN, _ := strconv.Atoi(count)
		if atlasIndex == "" {
			for range countN {
				indices = append(indices, -1)
			}
			continue
		}
		for range countN {
			i, _ := strconv.Atoi(atlasIndex)
			indices = append(indices, i)
		}
	}
	if len(indices) != tileCount {
		log.Fatal("Failed to parse layer.")
	}
	return indices
}

func LoadAnimationAtlas(name string) (AnimationMap, error) {
	data, err := os.ReadFile(fmt.Sprintf(atlasConfigPath, strings.ToLower(name)))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var (
		frameWidth  int
		frameHeight int
		img         *eb.Image
		atlas       *Atlas
		rawAnims    []string
	)
	animsMap := make(AnimationMap)

	for _, line := range lines {
		k, v, err := ParseKV(line)
		if err != nil {
			continue // likely a comment or blank line
		}
		switch k {
		case "atlas_path":
			img, _, err = ebutil.NewImageFromFile(v)
			if err != nil {
				return nil, err
			}
		case "frame_width":
			frameWidth, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		case "frame_height":
			frameHeight, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		case "anim":
			rawAnims = append(rawAnims, v)
		}
	}

	size := img.Bounds().Size()
	rows := size.Y / frameHeight
	cols := size.X / frameWidth
	atlas = NewAtlas(img, rows, cols, frameWidth, frameHeight)

	for _, rawAnim := range rawAnims {
		name, anim, err := parseAnimation(rawAnim, atlas)
		if err != nil {
			return nil, err
		}
		animsMap[name] = anim
	}
	return animsMap, nil
}

func parseAnimation(line string, atlas *Atlas) (string, *Animation, error) {
	var (
		name     string
		row      int
		duration int
		frames   int
		loop     bool
	)
	num, err := fmt.Sscanf(line, "%s %d %d %d %t", &name, &row, &duration, &frames, &loop)

	if num != 5 {
		return "", nil, fmt.Errorf("Incorrect number of arguments for animation: %d", num)
	}
	if err != nil {
		return "", nil, err
	}

	return name, NewAnimation(atlas, row, duration, frames, loop), nil
}

func ParseKV(line string) (string, string, error) {
	key, value, found := strings.Cut(line, "=")
	if !found {
		return "", "", fmt.Errorf("Invalid line %s", line)
	}
	return strings.ToLower(strings.TrimSpace(key)), strings.TrimSpace(value), nil
}
