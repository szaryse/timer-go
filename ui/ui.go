package ui

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/szaryse/timer-go255/timer"
	"image/color"
	"log"
)

//go:embed ..\assets\fonts\FiraCodeRegular.ttf
var FiraCodeRegular []byte

//go:embed ..\assets\fonts\DotoRegular.ttf
var DotoRegular []byte

var (
	firaCodeSource *text.GoTextFaceSource
	dotoSource     *text.GoTextFaceSource
)

func init() {
	fcs, err := text.NewGoTextFaceSource(bytes.NewReader(FiraCodeRegular))
	if err != nil {
		log.Fatal(err)
	}
	firaCodeSource = fcs

	ds, err := text.NewGoTextFaceSource(bytes.NewReader(DotoRegular))
	if err != nil {
		log.Fatal(err)
	}
	dotoSource = ds
}

type ViewState int

const (
	SettingsView = iota
	TimerView
)

const (
	SettingsWidth  = 640
	SettingsHeight = 480
	TimerWidth     = 960
	TimerHeight    = 240
)

type UI struct {
	CurrentView     ViewState
	WindowPositionX int
	WindowPositionY int
}

func CreateUI() UI {
	return UI{CurrentView: SettingsView}
}

func (ui *UI) Render(screen *ebiten.Image, t *timer.Timer) {
	// todo create the settings view
	if ui.CurrentView == SettingsView {
		op := &text.DrawOptions{}
		op.GeoM.Translate(0, 0)
		op.ColorScale.ScaleWithColor(color.White)
		status := "TEST"
		text.Draw(screen, status, &text.GoTextFace{
			Source: dotoSource,
			Size:   24,
		}, op)
	}
	// todo create the timer view
	if ui.CurrentView == TimerView {
		op := &text.DrawOptions{}
		op.GeoM.Translate(50, 60)
		op.ColorScale.ScaleWithColor(color.White)
		time := fmt.Sprintf("%d", t.Count/60)
		text.Draw(screen, time, &text.GoTextFace{
			Source: firaCodeSource,
			Size:   24,
		}, op)
	}
}

func (ui *UI) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		ui.changeView()
	}
	return nil
}

func (ui *UI) changeView() {
	ui.WindowPositionX, ui.WindowPositionY = ebiten.WindowPosition()

	switch ui.CurrentView {
	case SettingsView:
		ui.CurrentView = TimerView
		ebiten.SetWindowSize(TimerWidth, TimerHeight)
		ui.WindowPositionY += SettingsHeight - TimerHeight
	case TimerView:
		ui.CurrentView = SettingsView
		ebiten.SetWindowSize(SettingsWidth, SettingsHeight)
		ui.WindowPositionY -= SettingsHeight - TimerHeight
	default:
		panic(fmt.Errorf("unknown state: %q", ui.CurrentView))
	}
	ebiten.SetWindowPosition(ui.WindowPositionX, ui.WindowPositionY)
}
