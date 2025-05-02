package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Button struct {
	x, y, w, h int
	label      string
	action     string
	color      color.Color
}

func createButton(x int, label, action string) Button {
	return Button{
		x,
		0,
		0,
		buttonHeight,
		label,
		action,
		PrimaryColor,
	}
}

func createSettingsButtons() [12]Button {
	var buttons [12]Button
	buttons[0] = createButton(btnLx, "+", "increaseStart")
	buttons[1] = createButton(btnRx, "-", "decreaseStart")
	buttons[2] = createButton(btnLx, "+", "increaseSession")
	buttons[3] = createButton(btnRx, "-", "decreaseSession")
	buttons[4] = createButton(btnLx, "+", "increaseFocus")
	buttons[5] = createButton(btnRx, "-", "decreaseFocus")
	buttons[6] = createButton(btnLx, "+", "increaseBreak")
	buttons[7] = createButton(btnRx, "-", "decreaseBreak")
	buttons[8] = createButton(70, "Exit", "Exit")
	buttons[9] = createButton(260, "Start", "Start")

	for idx := range buttons {
		buttons[idx].w = buttonWidth
		buttons[idx].y = calcRowY(idx / 2)
	}
	buttons[8].w = textBtnWidth
	buttons[8].y = actionBtnY
	buttons[9].w = textBtnWidth
	buttons[9].y = actionBtnY

	return buttons
}

func (ui *UI) drawButton(screen *ebiten.Image, idx int) {
	btn := ui.SettingsButtons[idx]
	vector.StrokeRect(screen, float32(btn.x), float32(btn.y+btnMy), float32(btn.w), float32(btn.h), strokeWidth, btn.color, true)
	renderButtonText(screen, float64(btn.x), float64(btn.y), float64(btn.w), btn.label)
}

func renderButtonText(screen *ebiten.Image, x, y, w float64, label string) {
	textFace := &text.GoTextFace{
		Source: firaCodeSource,
		Size:   btnFontSize,
	}
	labelWidth, labelHeight := text.Measure(label, textFace, 0)
	x += (w - labelWidth) / 2
	y += ((buttonHeight - labelHeight) / 2) + btnMy + (strokeWidth / 2)

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(TextColor)

	text.Draw(screen, label, textFace, op)
}
