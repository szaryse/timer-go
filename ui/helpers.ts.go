package ui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

const (
	uiRadius = 8
	py       = 4 // padding y
)

var Subtitles = map[string]string{
	"StartingIn":    "Zaczynamy",
	"SessionNumber": "Liczba sesji",
	"FocusTime":     "Skupienie",
	"BreakTime":     "Przerwa",
	"StreamTime":    "Transmisja",
	"OneLineView":   "Jedna linia",
	"Exit":          "Wyjście",
	"Start":         "Start",
	"Timer":         "Minutnik",
	"LastSession":   "Koniec",
	"Timeout":       "Nadgodziny", //
}

func calcRowY(idx int) int {
	return py + (idx*fontSize + (idx * my * 2))
}

func formatTime(time int) string {
	if time < 0 {
		time *= -1
	}
	minutes := time / 60
	seconds := time - minutes*60
	minutesStr := fmt.Sprintf("%2d", minutes)
	secondsStr := fmt.Sprintf("%02d", seconds)
	return fmt.Sprintf("%s:%s", minutesStr, secondsStr)
}

func formatFullTime(time int) string {
	if time < 0 {
		time *= -1
	}
	hours := time / 3600
	minutes := (time - hours*3600) / 60
	seconds := time - hours*3600 - minutes*60
	hoursStr := fmt.Sprintf("%d", hours)
	minutesStr := fmt.Sprintf("%02d", minutes)
	secondsStr := fmt.Sprintf("%02d", seconds)
	return fmt.Sprintf("%s:%s:%s", hoursStr, minutesStr, secondsStr)
}

type uiRect struct {
	x, y, w, h int
}

func createRoundedPath(w, h float32) (path vector.Path) {
	r := float32(uiRadius)
	bm := float32(btnMy)
	ds := float32(strokeWidth / 2)
	a90 := float32((90 * math.Pi) / 180)
	a180 := float32((180 * math.Pi) / 180)
	a270 := float32((270 * math.Pi) / 180)
	a360 := float32((360 * math.Pi) / 180)

	// top left corner + top line
	path.Arc(r+ds, r+bm+ds, r, a180, a270, 0)
	path.LineTo(w-r-ds, bm+ds)
	// top right corner + right line
	path.Arc(w-r-ds, bm+r+ds, r, a270, a360, 0)
	path.LineTo(w-ds, bm+h-r-ds)
	// bottom right corner + bottom line
	path.Arc(w-r-ds, bm+h-r-ds, r, 0, a90, 0)
	path.LineTo(r+ds, bm+h-ds)
	// bottom left corner + left line
	path.Arc(r+ds, bm+h-r-ds, r, a90, a180, 0)
	path.LineTo(ds, r+bm+ds)
	path.Close()
	return
}

func createStrokeOptions() *vector.StrokeOptions {
	ops := &vector.StrokeOptions{}
	ops.Width = strokeWidth
	ops.LineJoin = vector.LineJoinMiter
	return ops
}

func createTrianglesOptions() *ebiten.DrawTrianglesOptions {
	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.FillRuleNonZero
	return op
}

func checkElementIsSelected(cursorX, cursorY int, box *uiRect) bool {
	return cursorX > box.x &&
		cursorX < box.x+box.w &&
		cursorY > box.y+btnMy &&
		cursorY < box.y+box.h+btnMy
}
