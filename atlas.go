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

func (a *Atlas) GetFrame(row, col int) image.Image {
	if row >= a.rows || col >= a.cols {
		return nil
	}
	index := row*a.cols + col
	return a.Frames[index]
}

func loadFrame(img *eb.Image, row, col int, frameWidth, frameHeight int) *eb.Image {
	frameX := col * frameWidth
	frameY := row * frameHeight
	frameRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)
	return img.SubImage(frameRect).(*eb.Image)
}

func (a *Atlas) getDims() (rows, cols int) {
	return a.rows, a.cols
}

func (a *Atlas) getPixelDims() (width, height int) {
	return a.cols * a.FrameWidth, a.rows * a.FrameHeight
}

func NewAtlas(img *eb.Image, rows, cols int, frameWidth, frameHeight int) *Atlas {
	frames := make([]*eb.Image, 0, rows*cols)
	for r := range rows {
		for c := range cols {
			frames = append(frames, loadFrame(img, r, c, frameWidth, frameHeight))
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
