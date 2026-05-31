package gui

import "github.com/hajimehoshi/ebiten/v2"

type Element interface {
	Update()
	Draw(dst *ebiten.Image)
}

type Input interface {
	GetValue() string
	OnClick()
	OnHover()
}
