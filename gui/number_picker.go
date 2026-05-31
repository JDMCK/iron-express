package gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type NumberPicker struct {
	x, y     int
	value    int
	buttons  [4]*Button
	valueImg *ebiten.Image
	rect     Rectangle
	text     Text
}

const defaultNPButtonSize = 18

var defaultNPButtonColor = color.Gray{100}

func newNumberPickerButton(jump int, x, y int, onClick func()) *Button {
	label := strconv.Itoa(jump)
	return NewButton(label, defaultNPButtonSize, defaultNPButtonSize, x, y, 16, defaultNPButtonColor, onClick)
}

func NewNumberPicker(smallJump, bigJump int, initialValue int, x, y int) *NumberPicker {
	np := NumberPicker{
		value: initialValue,
		x:     x,
		y:     y,
	}
	btns := [4]*Button{
		newNumberPickerButton(-bigJump, x, y, func() { np.value -= bigJump }),
		newNumberPickerButton(-smallJump, x+defaultNPButtonSize, y, func() { np.value -= smallJump }),
		newNumberPickerButton(smallJump, x+(4*defaultNPButtonSize), y, func() { np.value += smallJump }),
		newNumberPickerButton(bigJump, x+(5*defaultNPButtonSize), y, func() { np.value += bigJump }),
	}
	np.buttons = btns
	np.rect = Rectangle{x, y, defaultNPButtonSize * 6, defaultNPButtonSize}
	text := NewText(strconv.Itoa(np.value), np.x, np.y, 16, defaultNPButtonColor)
	text.CenterInRectangle(np.rect)
	np.text = text
	return &np
}

func (n *NumberPicker) Update() {
	for _, el := range n.buttons {
		el.Update()
	}
	n.text.SetValue(strconv.Itoa(n.value))
	n.text.CenterInRectangle(n.rect)
}

func (n *NumberPicker) Draw(screen *ebiten.Image) {
	for _, btn := range n.buttons {
		btn.Draw(screen)
	}
	n.text.Draw(screen)
}
