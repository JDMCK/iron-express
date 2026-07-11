package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2"
	ebutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const inputConfigPath = "config/inputs.config"
const atlasConfigPath = "config/sprites/%s.atlas.config"
const mapConfigPath = "config/levels/%s.map.config"

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
		}
	}
	size := img.Bounds().Size()
	cols = size.X / frameWidth
	rows = size.Y / frameHeight
	return NewAtlas(img, cols, rows, frameWidth, frameHeight), nil
}

func LoadLevelLayers(name string) []*Layer {
	data, err := os.ReadFile(fmt.Sprintf(mapConfigPath, strings.ToLower(name)))
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitSeq(string(data), "\n")

	layers := make([]*Layer, 0, 3)
	atlasIndices := make([][]int, 0, 1024)

	var (
		atlasPath  string
		img        *ebiten.Image
		tileWidth  int
		tileHeight int
		mapWidth   int
		mapHeight  int
	)

	for line := range lines {
		k, v, err := ParseKV(line)
		if err != nil {
			continue // likely a comment or blank line
		}
		switch k {
		case "atlas_path":
			atlasPath = v
		case "frame_width":
			tileWidth, _ = strconv.Atoi(v)
		case "frame_height":
			tileHeight, _ = strconv.Atoi(v)
		case "map_width":
			mapWidth, _ = strconv.Atoi(v)
		case "map_height":
			mapHeight, _ = strconv.Atoi(v)
		}

		if strings.HasPrefix(k, "layer_") {
			var layer int
			fmt.Sscanf(k, "layer_%d", &layer)
			atlasIndices = append(atlasIndices, parseLayer(v))
		}
	}

	img, _, _ = ebutil.NewImageFromFile(atlasPath)
	size := img.Bounds().Size()
	cols := size.X / tileWidth
	rows := size.Y / tileHeight
	atlas := NewAtlas(img, cols, rows, tileWidth, tileHeight)

	if err != nil {
		log.Fatalf("Failed to load atlas: %s", atlasPath)
	}

	for _, i := range atlasIndices {
		layers = append(layers, NewLayer(mapWidth, mapHeight, tileWidth, tileHeight, atlas, i))
	}

	return layers
}

func parseLayer(data string) []int {
	parts := strings.Split(data, " ")
	indices := make([]int, 0, 526)
	for _, p := range parts {
		atlasIndex, count, _ := strings.Cut(p, "-")
		if count == "" {
			log.Fatalf("Failed to parse layer: %s", data)
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
	cols := size.X / frameWidth
	rows := size.Y / frameHeight
	atlas = NewAtlas(img, cols, rows, frameWidth, frameHeight)

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
