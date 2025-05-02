package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontSize     = 24 // line height (31.5) = 24 + (2*4)
	btnFontSize  = 18
	lx           = 10  // label x
	btnLx        = 232 // button left x
	vx           = 270 // value x
	btnRx        = 400 // button right x
	py           = 4   // padding y
	my           = 4   // margin y
	btnMy        = 4   // button margin y
	actionBtnY   = 170
	valueWidth   = 120
	buttonHeight = 22 // fontSize + my*2 - 2*btnM
	buttonWidth  = 24 // default button width
	textBtnWidth = 100
	strokeWidth  = 2
)

func (ui *UI) RenderSettingView(screen *ebiten.Image) {
	ui.renderRow(screen, 0, "Starting in ", "5:00")
	ui.renderRow(screen, 1, "Session number", "6")
	ui.renderRow(screen, 2, "Focus time", "25:00")
	ui.renderRow(screen, 3, "Break time", "5:00")
	ui.renderRow(screen, 4, "Stream time", "3:05:00")
	ui.drawButton(screen, 8)
	ui.drawButton(screen, 9)
}

func renderText(screen *ebiten.Image, x, y float64, label string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(TextColor)
	textFace := &text.GoTextFace{
		Source: firaCodeSource,
		Size:   fontSize,
	}
	text.Draw(screen, label, textFace, op)
}

func renderCenteredText(screen *ebiten.Image, x, y, w float64, label string) {
	textFace := &text.GoTextFace{
		Source: firaCodeSource,
		Size:   fontSize,
	}
	labelWidth, _ := text.Measure(label, textFace, 0)
	x += (w - labelWidth) / 2

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(TextColor)

	text.Draw(screen, label, textFace, op)
}

func (ui *UI) renderRow(screen *ebiten.Image, rowIndex int, label, value string) {
	rowY := float64(calcRowY(rowIndex))
	renderText(screen, lx, rowY, label)
	ui.drawButton(screen, 2*rowIndex)
	renderCenteredText(screen, vx, rowY, valueWidth, value)
	ui.drawButton(screen, 2*rowIndex+1)
}

func calcRowY(idx int) int {
	return py + (idx*fontSize + (idx * my * 2))
}
