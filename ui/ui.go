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
var firaCodeSource *text.GoTextFaceSource

func init() {
	fcs, err := text.NewGoTextFaceSource(bytes.NewReader(FiraCodeRegular))
	if err != nil {
		log.Fatal(err)
	}
	firaCodeSource = fcs
}

//go:embed ..\assets\fonts\DotoRegular.ttf
var DotoRegular []byte
var dotoSource *text.GoTextFaceSource

func init() {
	ds, err := text.NewGoTextFaceSource(bytes.NewReader(DotoRegular))
	if err != nil {
		log.Fatal(err)
	}
	dotoSource = ds
}

type ViewState int

const (
	SettingsView ViewState = iota
	TimerView
)

const (
	SettingsWidth  = 440
	SettingsHeight = 220
	TimerWidth     = 640
	BgPadding      = 4
	TimerHeight    = 54 + 2*BgPadding
	leftX          = 10 // left padding x (inside background)
)

var (
	TextColor    = color.Gray{Y: 192}
	PrimaryColor = color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
	CyanColor    = color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff}
	RedColor     = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	bgColor      = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xc0}
)

type UI struct {
	CurrentView     ViewState
	WindowPositionX int
	WindowPositionY int
	SettingsButtons [12]Button
	SelectedAction  string
}

func CreateUI() UI {
	return UI{
		CurrentView:     SettingsView,
		SettingsButtons: createSettingsButtons(),
	}
}

func (ui *UI) Render(screen *ebiten.Image, t *timer.Timer) {
	ui.DrawBackground(screen)
	if ui.CurrentView == SettingsView {
		ui.RenderSettingView(screen, t)
	}
	if ui.CurrentView == TimerView {
		ui.RenderTimerView(screen, t)
	}
}

func (ui *UI) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if ui.CurrentView == TimerView {
			ui.changeView()
		}
		if ui.CurrentView == SettingsView {
			ui.handleClickOnSettings()
		}
	}
	return nil
}

func (ui *UI) ActionUpdate(t *timer.Timer) error {
	action := ui.SelectedAction
	if len(action) > 0 {
		if action == "exit" {
			return ebiten.Termination
		}
		t.HandleAction(action)
		ui.SelectedAction = ""
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
