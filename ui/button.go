package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Button struct {
	x, y, w, h int
	label      string
	color      color.Color
}

func createButton(x int, label string) Button {
	return Button{
		x,
		0,
		0,
		buttonHeight,
		label,
		PrimaryColor,
	}
}

func createSettingsButtons() [12]Button {
	var buttons [12]Button

	buttons[0] = createButton(btnLx, "+")
	buttons[1] = createButton(btnRx, "-")
	buttons[2] = createButton(btnLx, "+")
	buttons[3] = createButton(btnRx, "-")
	buttons[4] = createButton(btnLx, "+")
	buttons[5] = createButton(btnRx, "-")
	buttons[6] = createButton(btnLx, "+")
	buttons[7] = createButton(btnRx, "-")
	buttons[8] = createButton(btnLx, "+")
	buttons[9] = createButton(btnRx, "-")

	buttons[10] = createButton(160, "Exit")
	buttons[11] = createButton(320, "Start")

	for idx := range buttons {
		buttons[idx].w = buttonWidth
	}

	buttons[10].w = textBtnWidth
	buttons[11].w = textBtnWidth

	return buttons
}

func (ui *UI) drawButton(screen *ebiten.Image, idx int, y float64) {
	btn := ui.SettingsButtons[idx]
	btn.y = int(y) + btnMy

	vector.StrokeRect(screen, float32(btn.x), float32(btn.y), float32(btn.w), float32(btn.h), strokeWidth, btn.color, true)
	renderCenteredText(screen, float64(btn.x), y, float64(btn.w), btn.label)
}
