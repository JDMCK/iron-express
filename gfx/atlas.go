package gfx

import (
	"image"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Atlas struct {
	// TODO: cache each subimage instead of creating one each frame
	img         *eb.Image
	frameWidth  int
	frameHeight int
}

func (a *Atlas) GetFrame(row, col int) image.Image {
	totalRows, totalCols := a.getDims()
	boundedRow := row % totalRows
	boundedCol := col % totalCols
	frameX := boundedCol * a.frameWidth
	frameY := boundedRow * a.frameHeight
	frameRect := image.Rect(frameX, frameY, frameX+a.frameWidth, frameY+a.frameHeight)
	return a.img.SubImage(frameRect)

}

// returns amount of rows and cols of frames in this atlas
func (a *Atlas) getDims() (rows, cols int) {
	width, height := a.getPixelDims()
	return height / a.frameHeight, width / a.frameWidth
}

func (a *Atlas) getPixelDims() (width, height int) {
	point := a.img.Bounds().Size()
	return point.X, point.Y
}

func NewAtlas(img *eb.Image, frameWidth, frameHeight int) *Atlas {
	return &Atlas{
		img:         img,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
	}
}
