package ui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/szaryse/timer-go255/timer"
)

const (
	fontSize     = 24 // line height (31.5) = 24 + (2*4)
	btnFontSize  = 18
	btnLx        = 232 // button left x
	vx           = 270 // value x
	btnRx        = 400 // button right x
	py           = 4   // padding y
	my           = 4   // margin y
	btnMy        = 4   // button margin y
	actionBtnY   = 175
	valueWidth   = 120
	buttonHeight = 22 // fontSize + my*2 - 2*btnM
	buttonWidth  = 24 // default button width
	textBtnWidth = 100
	strokeWidth  = 2
)

func (ui *UI) RenderSettingView(screen *ebiten.Image, timer *timer.Timer) {
	ui.renderRow(screen, 0, "Starting in", timer.StartInTime)
	ui.renderRow(screen, 1, "Session number", timer.SessionNumber)
	ui.renderRow(screen, 2, "Focus time", timer.FocusTime)
	ui.renderRow(screen, 3, "Break time", timer.BreakTime)
	ui.renderRowWithoutButtons(screen, 4, "Stream time", timer.StreamTime)
	ui.drawButton(screen, 8)
	if ui.beforeStart {
		ui.drawButton(screen, 9)
	} else {
		ui.drawButton(screen, 10)
	}
}

func (ui *UI) renderRow(screen *ebiten.Image, rowIndex int, label string, value int) {
	rowY := float64(calcRowY(rowIndex))
	renderText(screen, leftX, rowY, label)
	ui.drawButton(screen, 2*rowIndex)
	valueStr := formatValue(value, label)
	renderCenteredText(screen, vx, rowY, valueWidth, valueStr)
	ui.drawButton(screen, 2*rowIndex+1)
}

func (ui *UI) renderRowWithoutButtons(screen *ebiten.Image, rowIndex int, label string, value int) {
	rowY := float64(calcRowY(rowIndex))
	renderText(screen, leftX, rowY, label)
	valueStr := formatValue(value, label)
	renderCenteredText(screen, vx, rowY, valueWidth, valueStr)
}

func (ui *UI) handleClickOnSettings() {
	cursorX, cursorY := ebiten.CursorPosition()
	for _, button := range ui.SettingsButtons {
		if cursorX > button.x &&
			cursorX < button.x+button.w &&
			cursorY > button.y &&
			cursorY < button.y+button.h {

			switch button.action {
			case "start":
				if ui.beforeStart {
					fmt.Println("Click", button.action)
					ui.SelectedAction = button.action
					ui.changeView()
				}
			case "changeView":
				if ui.beforeStart == false {
					ui.changeView()
				}
			default:
				ui.SelectedAction = button.action
			}
		}
	}
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

func formatValue(value int, label string) (valueStr string) {
	if label == "Session number" {
		valueStr = fmt.Sprintf("%d", value)
	} else if label == "Stream time" {
		valueStr = formatFullTime(value)
	} else {
		valueStr = formatTime(value)
	}
	return
}
