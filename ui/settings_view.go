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

type ButtonsArray [11]Button
type CheckboxArray [1]CheckBox

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

func (ui *UI) RenderSettingView(screen *ebiten.Image, t *timer.Timer) {
	ui.renderRow(screen, 0, "Starting in", t.StartInTime)
	ui.renderRow(screen, 1, "Session number", t.SessionNumber)
	ui.renderRow(screen, 2, "Focus time", t.FocusTime)
	ui.renderRow(screen, 3, "Break time", t.BreakTime)
	ui.renderRowWithoutButtons(screen, 4, "Stream time", t.StreamTime)
	ui.renderRowWithCheckbox(screen, 5, "Stream time only")
	ui.renderButton(screen, 8)
	if t.IsRunning == false {
		ui.renderButton(screen, 9)
	} else {
		ui.renderButton(screen, 10)
	}
}

func (ui *UI) renderRow(screen *ebiten.Image, rowIndex int, label string, value int) {
	rowY := float64(calcRowY(rowIndex))
	renderText(screen, leftX, rowY, label)
	ui.renderButton(screen, 2*rowIndex)
	valueStr := formatValue(value, label)
	renderCenteredText(screen, vx, rowY, valueWidth, valueStr)
	ui.renderButton(screen, 2*rowIndex+1)
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

func CreateButtons() ButtonsArray {
	var buttons ButtonsArray
	buttons[0] = NewButton(btnLx, "+")    // increase start
	buttons[1] = NewButton(btnRx, "-")    // decrease start
	buttons[2] = NewButton(btnLx, "+")    // increase session
	buttons[3] = NewButton(btnRx, "-")    // decrease session
	buttons[4] = NewButton(btnLx, "+")    // increase focus
	buttons[5] = NewButton(btnRx, "-")    // decrease focus
	buttons[6] = NewButton(btnLx, "+")    // increase break
	buttons[7] = NewButton(btnRx, "-")    // decrease break
	buttons[8] = NewButton(70, "Exit")    // exit
	buttons[9] = NewButton(260, "Start")  // start
	buttons[10] = NewButton(260, "Timer") // change view

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

func CreateCheckboxes() (checkboxes CheckboxArray) {
	checkboxes[0] = NewCheckbox()
	return
}
