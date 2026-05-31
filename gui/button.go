package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	label     Text
	value     string
	img       *ebiten.Image
	rect      Rectangle
	onClick   func()
	isHovered bool
}

const DefaultButtonHeight = 40

type ButtonSize int
type ButtonColor int

const (
	Large ButtonSize = iota
	Medium
	Small
)

var buttonSizes = map[ButtonSize]int{
	Large:  200,
	Medium: 150,
	Small:  100,
}

const (
	Primary ButtonColor = iota
	Secondary
	Danger
)

var buttonColors = map[ButtonColor]color.Color{
	Primary:   color.RGBA{10, 70, 255, 255},
	Secondary: color.Gray{200},
	Danger:    color.RGBA{255, 20, 20, 255},
}

func NewButton(
	label string,
	width, height, x, y int,
	size float64,
	col color.Color,
	onClick func(),
) *Button {
	img := ebiten.NewImage(width, height)
	img.Fill(col)
	return NewButtonFromImage(label, img, width, height, x, y, size, col, onClick)
}

func NewButtonFromImage(
	label string,
	img *ebiten.Image,
	width, height, x, y int,
	size float64,
	col color.Color,
	onClick func(),
) *Button {
	rect := Rectangle{
		width:  width,
		height: height,
		x:      x,
		y:      y,
	}
	text := NewText(label, x, y, size, color.RGBA{255, 255, 255, 255})
	text.CenterInRectangle(rect)
	return &Button{
		label:   text,
		img:     img,
		rect:    rect,
		onClick: onClick,
	}
}

func NewBasicButton(label string, x, y int, size ButtonSize, col ButtonColor, onClick func()) *Button {
	width := buttonSizes[size]
	return NewButton(label, width, DefaultButtonHeight, x, y, DefaultFontSize, buttonColors[col], onClick)
}

func (b *Button) SetColor(col color.Color) {
	b.img.Fill(col)
}

func (b *Button) Update() {
	mx, my := ebiten.CursorPosition()
	b.isHovered = b.rect.InBounds(mx, my)

	if b.isHovered && inpututil.IsMouseButtonJustPressed(PrimaryButton) {
		b.onClick()
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.rect.x), float64(b.rect.y))
	if b.isHovered {
		op.ColorScale.ScaleWithColor(color.Gray{200}) // darken on hover
	}
	screen.DrawImage(b.img, op)
	b.label.Draw(screen)
}
