package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontSize   = 24 // 18 + 2*3px
	lx         = 10
	pY         = 4
	vx         = 320
	mb         = 8
	valueWidth = 120
)

func (ui *UI) RenderSettingView(screen *ebiten.Image) {
	renderRow(screen, 0, "Starting in ", "5:00")
	renderRow(screen, 1, "Session number", "6")
	renderRow(screen, 2, "Focus time", "25:00")
	renderRow(screen, 3, "Break time", "5:00")
	renderRow(screen, 4, "Stream time", "3:05:00")
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

func renderCenteredText(screen *ebiten.Image, x, y float64, label string) {
	textFace := &text.GoTextFace{
		Source: firaCodeSource,
		Size:   fontSize,
	}
	labelWidth, _ := text.Measure(label, textFace, 0)
	x += (valueWidth - labelWidth) / 2

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(TextColor)

	text.Draw(screen, label, textFace, op)
}

func renderRow(screen *ebiten.Image, idx int, label, value string) {
	rowY := float64(pY + (idx*fontSize + idx*mb))
	renderText(screen, lx, rowY, label)
	renderCenteredText(screen, vx, rowY, value)
}
