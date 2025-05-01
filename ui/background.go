package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const padding = 4

var bgColor = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xc0}

func (ui *UI) DrawBackground(screen *ebiten.Image) {
	var width, height float32

	if ui.CurrentView == SettingsView {
		width = float32(SettingsWidth - (2 * padding))
		height = float32(SettingsHeight - (2 * padding))
	}
	if ui.CurrentView == TimerView {
		width = float32(TimerWidth - (2 * padding))
		height = float32(TimerHeight - (2 * padding))
	}

	vector.DrawFilledRect(screen, padding, padding, width, height, bgColor, true)
}
