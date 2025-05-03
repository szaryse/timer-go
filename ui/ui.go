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
	BgPadding      = 4
	SettingsWidth  = 440
	SettingsHeight = 220
	TimerWidth     = 440
	TimerHeight    = 54 + 2*BgPadding
)

var (
	TextColor    = color.Gray{Y: 192}
	PrimaryColor = color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
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
	// todo #1 create the timer view
	if ui.CurrentView == TimerView {
		op := &text.DrawOptions{}
		op.GeoM.Translate(0, 0)
		op.ColorScale.ScaleWithColor(TextColor)
		time := formatFullTime(t.Count / timer.Tick)
		text.Draw(screen, time, &text.GoTextFace{
			Source: firaCodeSource,
			Size:   24,
		}, op)
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

func (ui *UI) handleClickOnSettings() {
	cursorX, cursorY := ebiten.CursorPosition()
	for _, button := range ui.SettingsButtons {
		if cursorX > button.x &&
			cursorX < button.x+button.w &&
			cursorY > button.y &&
			cursorY < button.y+button.h {
			ui.SelectedAction = button.action
			if button.action == "start" {
				ui.changeView()
			}
		}
	}
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
