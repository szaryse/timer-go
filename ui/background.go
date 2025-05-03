package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (ui *UI) DrawBackground(screen *ebiten.Image) {
	var width, height float32

	if ui.CurrentView == SettingsView {
		width = float32(SettingsWidth - (2 * BgPadding))
		height = float32(SettingsHeight - (2 * BgPadding))
	}
	if ui.CurrentView == TimerView {
		width = float32(TimerWidth - (2 * BgPadding))
		height = float32(TimerHeight - (2 * BgPadding))
	}

	vector.DrawFilledRect(screen, BgPadding, BgPadding, width, height, bgColor, true)
}
