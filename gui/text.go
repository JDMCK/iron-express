package gui

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed Inter_18pt-Regular.ttf
var interTTF []byte

const DefaultDPI = 72
const DefaultFontSize = 18

type Text struct {
	value string
	color color.Color
	size  float64
	face  *text.GoTextFace // font face
	rect  Rectangle
}

func NewText(value string, x, y int, size float64, col color.Color) Text {

	src, err := text.NewGoTextFaceSource(bytes.NewReader(interTTF))
	if err != nil {
		log.Fatal(err)
	}

	face := &text.GoTextFace{
		Source: src,
		Size:   size,
	}

	w, h := text.Measure(value, face, 0)
	return Text{
		value: value,
		color: col,
		size:  size,
		face:  face,
		rect:  Rectangle{x, y, int(w), int(h)},
	}

}

func (t *Text) CenterInRectangle(rect Rectangle) {
	dx := rect.width - t.rect.width
	dy := rect.height - t.rect.height
	newX := dx / 2
	newY := dy / 2
	t.rect.x = rect.x + newX
	t.rect.y = rect.y + newY
}

func (t *Text) SetValue(value string) {
	t.value = value
	w, h := text.Measure(value, t.face, 0)
	t.rect.width = int(w)
	t.rect.height = int(h)
}

func (t *Text) Update() {}

func (t *Text) Draw(screen *ebiten.Image) {
	txtOp := &text.DrawOptions{}
	txtOp.GeoM.Translate(float64(t.rect.x), float64(t.rect.y))
	text.Draw(screen, t.value, t.face, txtOp)
}
