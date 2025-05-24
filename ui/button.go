package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
)

const (
	textBtnWidth = 100
	actionBtnY   = 205
	btnFontSize  = 18
	btnRx        = 397 // button right x
	btnLx        = 229 // button left x
)

type Button struct {
	box    uiRect
	label  string
	action string
	color  color.Color
}

type ButtonsArray [11]Button

func createButton(x int, label, action string) Button {
	return Button{
		uiRect{x,
			0,
			0,
			buttonHeight},
		label,
		action,
		PrimaryColor,
	}
}

func createSettingsButtons() ButtonsArray {
	var buttons ButtonsArray
	buttons[0] = createButton(btnLx, "+", "increaseStart")
	buttons[1] = createButton(btnRx, "-", "decreaseStart")
	buttons[2] = createButton(btnLx, "+", "increaseSession")
	buttons[3] = createButton(btnRx, "-", "decreaseSession")
	buttons[4] = createButton(btnLx, "+", "increaseFocus")
	buttons[5] = createButton(btnRx, "-", "decreaseFocus")
	buttons[6] = createButton(btnLx, "+", "increaseBreak")
	buttons[7] = createButton(btnRx, "-", "decreaseBreak")
	buttons[8] = createButton(70, "Exit", "exit")
	buttons[9] = createButton(260, "Start", "start")
	buttons[10] = createButton(260, "Timer", "changeView")

	for idx := range buttons {
		buttons[idx].box.w = buttonWidth
		buttons[idx].box.y = calcRowY(idx / 2)
	}

	actionButtons := buttons[8:]
	for idx := range actionButtons {
		actionButtons[idx].box.w = textBtnWidth
		actionButtons[idx].box.y = actionBtnY
	}

	return buttons
}

func (ui *UI) drawButton(screen *ebiten.Image, idx int) {
	btn := ui.SettingsButtons[idx]
	path := createRoundedPath(float32(btn.box.w), float32(btn.box.h))
	ops := createStrokeOptions()
	ui.vertices, ui.indices = path.AppendVerticesAndIndicesForStroke(ui.vertices[:0], ui.indices[:0], ops)
	ui.setVertices(btn.color, btn.box.x, btn.box.y)
	op := createTrianglesOptions()
	screen.DrawTriangles(ui.vertices, ui.indices, whiteSubImage, op)

	renderButtonText(screen, float64(btn.box.x), float64(btn.box.y),
		float64(btn.box.w), btn.label)
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
