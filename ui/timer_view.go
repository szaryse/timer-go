package ui

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/szaryse/timer-go255/timer"
	"image/color"
	"math"
)

const (
	fontX float64 = 16
	row1y float64 = BgPadding
	row2y float64 = BgPadding + timerInnerHeight/2
)

func (ui *UI) RenderTimerView(screen *ebiten.Image, t *timer.Timer) {
	textFace := &text.GoTextFace{
		Source: firaCodeMSource,
		Size:   20.5793, // ratio 1,312
	}
	if ui.isOneLineView {
		ui.renderOneLine(screen, t, textFace)
	} else {
		ui.renderTwoLines(screen, t, textFace)
	}
}

func (ui *UI) renderTwoLines(screen *ebiten.Image, t *timer.Timer, textFace *text.GoTextFace) {
	renderLabel(screen, textFace, t, "", row1y, TextColor)
	renderTimeValue(screen, textFace, t.Count, row1y, formatTime)
	y := row2y
	renderLabel(screen, textFace, t, Subtitles["StreamTime"], y, TextColor)
	renderTimeValue(screen, textFace, t.TotalCount, y, getTimeString)
}

func (ui *UI) renderOneLine(screen *ebiten.Image, t *timer.Timer, textFace *text.GoTextFace) {
	time := t.TotalCount
	if t.Activity == timer.StartingInState {
		time = t.Count
	}
	timeInt := time / timer.Tick
	timeColor := setTimeColor(timeInt)

	label := Subtitles["StreamTime"]
	if t.Activity == timer.StartingInState {
		label = Subtitles["StartingInState"]
	}

	_, h := text.Measure("123456789012", textFace, 0)
	y := row2y - (0.5 * h)

	renderLabel(screen, textFace, t, label, y, timeColor)
	renderTimeValue(screen, textFace, time, y, getTimeString)
}

func renderLabel(screen *ebiten.Image, textFace *text.GoTextFace, t *timer.Timer,
	label string, y float64, color color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(fontX, y)
	op.ColorScale.ScaleWithColor(color)
	if len(label) == 0 {
		label = setSessionLabel(t.Activity, t.SessionNumber)
	}
	text.Draw(screen, label, textFace, op)
}

func renderTimeValue(screen *ebiten.Image, textFace *text.GoTextFace, value int, y float64, formatFn func(int) string) {
	op := &text.DrawOptions{}
	timeInt := value / timer.Tick
	timeColor := setTimeColor(timeInt)
	op.ColorScale.ScaleWithColor(timeColor)
	time := formatFn(timeInt)
	w, _ := text.Measure(time, textFace, 0)
	op.GeoM.Translate(TimerWidth-w-fontX, y)
	text.Draw(screen, time, textFace, op)
}

func getTimeString(timeInt int) (timeString string) {
	if timeInt >= 3600 {
		return formatFullTime(timeInt)
	}
	if timeInt >= 0 {
		return formatTime(timeInt)
	}
	if timeInt > -3600 {
		return "-" + formatTime(timeInt)
	}
	return "-" + formatFullTime(timeInt)
}

func setSessionLabel(activity timer.ActivityState, session int) string {
	switch activity {
	case timer.StartingInState:
		return Subtitles["StartingIn"]
	case timer.FocusState:
		if session > 1 {
			return Subtitles["FocusTime"] + fmt.Sprintf(" %d", session)
		} else {
			return Subtitles["LastSession"]
		}
	case timer.BreakState:
		return Subtitles["BreakTime"]
	case timer.TimeoutState:
		return Subtitles["Timeout"]
	default:
		return ""
	}
}

func setTimeColor(time int) color.RGBA {
	var r, g, b uint8
	var err error

	switch true {
	case time <= 20:
		r, g, b, err = HSLToRGB(0, 1, 0.6275) // red
	case time <= 200:
		r, g, b, err = HSLToRGB(float64(time-20), 1, 0.6275) // from cyan to red
	case time <= 250:
		s := (250 - float64(time)) / 50
		r, g, b, err = HSLToRGB(180, s, 0.6275) // from grey to cyan
	default:
		r, g, b, err = HSLToRGB(180, 0, 0.6275) // grey
	}

	if err != nil {
		fmt.Println(err)
		r, g, b, err = HSLToRGB(180, 0, 0.6275)
	}

	return color.RGBA{R: r, G: g, B: b, A: 0xff}
}

// HSLToRGB function from package colorconv.
// It provides conversion of color to HSL, HSV and hex value.
// https://github.com/Crazy3lf/colorconv/blob/master/colorconv.go
func HSLToRGB(h, s, l float64) (r, g, b uint8, err error) {
	if h < 0 || h >= 360 ||
		s < 0 || s > 1 ||
		l < 0 || l > 1 {
		errMessage := fmt.Sprintf("HSLToRGB: Invalid H/S/L values %f / %f / %f", h, s, l)
		return 0, 0, 0, errors.New(errMessage)
	}
	// When 0 ≤ h < 360, 0 ≤ s ≤ 1 and 0 ≤ l ≤ 1:
	C := (1 - math.Abs((2*l)-1)) * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - (C / 2)
	var Rnot, Gnot, Bnot float64

	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}
	r = uint8(math.Round((Rnot + m) * 255))
	g = uint8(math.Round((Gnot + m) * 255))
	b = uint8(math.Round((Bnot + m) * 255))
	return r, g, b, nil
}
