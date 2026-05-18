package config

import (
	"fmt"
	"iron-express/gfx"
	"os"
	"strconv"
	"strings"

	eb "github.com/hajimehoshi/ebiten/v2"
	ebutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const atlasConfigPath = "config/sprites/%s.atlas.config"

func LoadAnimationAtlas(name string) (gfx.AnimationMap, error) {
	data, err := os.ReadFile(fmt.Sprintf(atlasConfigPath, strings.ToLower(name)))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var (
		frameWidth  int
		frameHeight int
		img         *eb.Image
		atlas       *gfx.Atlas
		rawAnims    []string
	)
	animsMap := make(gfx.AnimationMap)

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

	atlas = gfx.NewAtlas(img, frameWidth, frameHeight)
	for _, rawAnim := range rawAnims {
		name, anim, err := parseAnimation(rawAnim, atlas)
		if err != nil {
			return nil, err
		}
		animsMap[name] = anim
	}
	return animsMap, nil
}

func parseAnimation(line string, atlas *gfx.Atlas) (string, *gfx.Animation, error) {
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

	return name, gfx.NewAnimation(atlas, row, duration, frames, loop), nil
}
