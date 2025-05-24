package ui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/szaryse/timer-go255/timer"
	"image"
	"image/color"
)

var (
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

const (
	fontSize     = 24  // line height (31.5) = 24 + (2*4)
	vx           = 270 // value x
	my           = 4   // margin y
	btnMy        = 3   // button margin y
	valueWidth   = 120
	buttonHeight = 24
	buttonWidth  = 30 // default button width
	strokeWidth  = 2
)

func (ui *UI) RenderSettingView(screen *ebiten.Image, timer *timer.Timer) {
	ui.renderRow(screen, 0, "Starting in", timer.StartInTime)
	ui.renderRow(screen, 1, "Session number", timer.SessionNumber)
	ui.renderRow(screen, 2, "Focus time", timer.FocusTime)
	ui.renderRow(screen, 3, "Break time", timer.BreakTime)
	ui.renderRowWithoutButtons(screen, 4, "Stream time", timer.StreamTime)
	ui.renderRowWithCheckbox(screen, 5, "Stream time only")
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

func (ui *UI) renderRowWithCheckbox(screen *ebiten.Image, rowIndex int, label string) {
	rowY := float64(calcRowY(rowIndex))
	renderText(screen, leftX, rowY, label)
	ui.drawCheckbox(screen)
}

func (ui *UI) handleClickOnSettings() {
	cursorX, cursorY := ebiten.CursorPosition()
	for idx := range ui.SettingsButtons {
		if checkIsElementSelected(cursorX, cursorY, &ui.SettingsButtons[idx].box) {
			ui.selectAction(&ui.SettingsButtons[idx])
		}
	}
	for idx := range ui.Checkboxes {
		if checkIsElementSelected(cursorX, cursorY, &ui.Checkboxes[idx].box) {
			ui.isStreamOnly = !ui.isStreamOnly
			ui.Checkboxes[idx].isChecked = ui.isStreamOnly
		}
	}
}

func (ui *UI) selectAction(button *Button) {
	switch button.action {
	case "start":
		if ui.beforeStart {
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

func (ui *UI) hoverElement() {
	cursorX, cursorY := ebiten.CursorPosition()
	buttonIdx := -1
	checkboxIdx := -1
	for idx := range ui.SettingsButtons {
		if checkIsElementSelected(cursorX, cursorY, &ui.SettingsButtons[idx].box) {
			if ui.SettingsButtons[idx].color == PrimaryColor {
				ui.SettingsButtons[idx].color = SelectedColor
			}
			if ui.beforeStart &&
				(ui.SettingsButtons[idx].label == "Start" ||
					ui.SettingsButtons[idx].label == "Timer") {
				buttonIdx = 9
			} else if !ui.beforeStart &&
				(ui.SettingsButtons[idx].label == "Start" ||
					ui.SettingsButtons[idx].label == "Timer") {
				buttonIdx = 10
			} else {
				buttonIdx = idx
			}
		}
	}
	ui.clearHoverOnButtons(buttonIdx)
	for idx := range ui.Checkboxes {
		if checkIsElementSelected(cursorX, cursorY, &ui.Checkboxes[idx].box) {
			if ui.Checkboxes[idx].color == PrimaryColor {
				ui.Checkboxes[idx].color = SelectedColor
			}
			checkboxIdx = idx
		}
	}
	ui.clearHoverOnCheckbox(checkboxIdx)
}

func (ui *UI) clearHoverOnButtons(selectedIdx int) {
	for idx := range ui.SettingsButtons {
		if selectedIdx != idx && ui.SettingsButtons[idx].color == SelectedColor {
			ui.SettingsButtons[idx].color = PrimaryColor
		}
	}
}

func (ui *UI) clearHoverOnCheckbox(selectedIdx int) {
	for idx := range ui.Checkboxes {
		if selectedIdx != idx && ui.Checkboxes[idx].color == SelectedColor {
			ui.Checkboxes[idx].color = PrimaryColor
		}
	}
}

func (ui *UI) setVertices(color color.Color, x, y int) {
	red, g, b, a := color.RGBA()
	for i := range ui.vertices {
		ui.vertices[i].DstX = ui.vertices[i].DstX + float32(x)
		ui.vertices[i].DstY = ui.vertices[i].DstY + float32(y)
		ui.vertices[i].SrcX = 1
		ui.vertices[i].SrcY = 1
		ui.vertices[i].ColorR = float32(red) / float32(0xffff)
		ui.vertices[i].ColorG = float32(g) / float32(0xffff)
		ui.vertices[i].ColorB = float32(b) / float32(0xffff)
		ui.vertices[i].ColorA = float32(a) / float32(0xffff)
	}
}

func checkIsElementSelected(cursorX, cursorY int, box *uiRect) bool {
	return cursorX > box.x &&
		cursorX < box.x+box.w &&
		cursorY > box.y+btnMy &&
		cursorY < box.y+box.h+btnMy
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
