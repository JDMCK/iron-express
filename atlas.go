package main

import (
	"image"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Atlas struct {
	Frames      []*eb.Image
	rows        int
	cols        int
	FrameWidth  int
	FrameHeight int
}

func (a *Atlas) GetFrameFromCoords(col, row int) image.Image {
	if col >= a.cols || row >= a.rows {
		return nil
	}
	index := row*a.cols + col
	return a.Frames[index]
}

func loadFrame(img *eb.Image, col, row int, frameWidth, frameHeight int) *eb.Image {
	frameX := col * frameWidth
	frameY := row * frameHeight
	frameRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)
	return img.SubImage(frameRect).(*eb.Image)
}

func NewAtlas(img *eb.Image, cols, rows int, frameWidth, frameHeight int) *Atlas {
	frames := make([]*eb.Image, 0, rows*cols)
	for r := range rows {
		for c := range cols {
			frames = append(frames, loadFrame(img, c, r, frameWidth, frameHeight))
		}
	}
	return &Atlas{
		Frames:      frames,
		rows:        rows,
		cols:        cols,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
	}
}
