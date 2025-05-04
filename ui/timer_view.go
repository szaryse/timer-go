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
	dotoPy       = 6
	dotoLx       = 16
	sessionX     = 175
	streamLabelX = 300
	streamTimeX  = 480
)

func (ui *UI) RenderTimerView(screen *ebiten.Image, t *timer.Timer) {
	textFace := &text.GoTextFace{
		Source: dotoSource,
		Size:   35,
	}
	renderSessionTime(screen, textFace, t)
	renderStreamTime(screen, textFace, t)
}

func renderSessionTime(screen *ebiten.Image, textFace *text.GoTextFace, t *timer.Timer) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(dotoLx, dotoPy)
	op.ColorScale.ScaleWithColor(color.White)
	label := setSessionLabel(t.Activity)
	text.Draw(screen, label, textFace, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(sessionX, dotoPy)
	timeInt := t.Count / timer.Tick
	timeColor := setTimeColor(timeInt)
	op.ColorScale.ScaleWithColor(timeColor)
	time := formatTime(timeInt)
	text.Draw(screen, time, textFace, op)
}

func renderStreamTime(screen *ebiten.Image, textFace *text.GoTextFace, t *timer.Timer) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(streamLabelX, dotoPy)
	op.ColorScale.ScaleWithColor(color.White)
	label := "| Stream"
	text.Draw(screen, label, textFace, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(streamTimeX, dotoPy)
	timeInt := t.TotalCount / timer.Tick
	timeColor := setTimeColor(timeInt)
	op.ColorScale.ScaleWithColor(timeColor)
	timeString := getTimeString(timeInt)
	text.Draw(screen, timeString, textFace, op)
}

func getTimeString(timeInt int) (timeString string) {
	if timeInt >= 3600 {
		return formatFullTime(timeInt)
	} else {
		return formatTime(timeInt)
	}
}

func setSessionLabel(activity timer.ActivityState) string {
	switch activity {
	case timer.StartingInState:
		return "Start"
	case timer.FocusState:
		return "Focus"
	case timer.BreakState:
		return "Break"
	case timer.TimeoutState:
		return "Timeout"
	}
	return ""
}

func setTimeColor(time int) color.RGBA {
	if time <= 30 {
		return RedColor
	}
	if time <= 210 {
		r, g, b, err := HSLToRGB(float64(time-30), 1, 0.5)
		if err != nil {
			fmt.Println(err)
			return CyanColor
		}
		return color.RGBA{R: r, G: g, B: b, A: 0xff}
	}
	return CyanColor
}

// HSLToRGB function from package colorconv.
// It provides conversion of color to HSL, HSV and hex value.
// https://github.com/Crazy3lf/colorconv/blob/master/colorconv.go
func HSLToRGB(h, s, l float64) (r, g, b uint8, err error) {
	if h < 0 || h >= 360 ||
		s < 0 || s > 1 ||
		l < 0 || l > 1 {
		return 0, 0, 0, errors.New("HSLToRGB: inputs out of range")
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
