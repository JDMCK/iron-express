package gfx

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Atlas struct {
	FrameCount int

	rows int
	cols int
	img  *ebiten.Image
}

func (a *Atlas) GetFrame(frame int) image.Image {
	boundedFrame := frame % a.FrameCount
	frameWidth, frameHeight := a.GetFrameDims()
	frameX := boundedFrame % a.cols * frameWidth
	frameY := boundedFrame / a.cols * frameHeight
	frameRect := image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeight)
	return a.img.SubImage(frameRect)
}

func (a *Atlas) GetFrameDims() (width, height int) {
	pt := a.img.Bounds().Size()
	return pt.X / a.cols, pt.Y / a.rows
}

func NewAtlas(img *ebiten.Image, frameCount int, cols int) Atlas {
	return Atlas{
		img:        img,
		FrameCount: frameCount,
		cols:       cols,
		rows:       frameCount / cols,
	}
}
