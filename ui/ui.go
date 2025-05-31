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

//go:embed ..\assets\fonts\SilkscreenRegular.ttf
var SilkscreenRegular []byte
var silkscreenSource *text.GoTextFaceSource

func init() {
	ds, err := text.NewGoTextFaceSource(bytes.NewReader(SilkscreenRegular))
	if err != nil {
		log.Fatal(err)
	}
	silkscreenSource = ds
}

type ViewState int

const (
	SettingsView ViewState = iota
	TimerView
)

const (
	SettingsWidth  = 440
	SettingsHeight = 250
	TimerWidth     = 360
	BgPadding      = 4
	TimerHeight    = 54 + 2*BgPadding
	leftX          = 10 // left padding x (inside background)
)

var (
	TextColor     = color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	PrimaryColor  = color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
	SelectedColor = color.RGBA{R: 0x00, G: 0xe0, B: 0x00, A: 0xff}
	bgColor       = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xcc}
)

type UI struct {
	CurrentView     ViewState
	WindowPositionX int
	WindowPositionY int
	SettingsButtons ButtonsArray
	SelectedAction  string
	Checkboxes      CheckboxArray
	isOneLineView   bool
	vertices        []ebiten.Vertex
	indices         []uint16
}

func CreateUI() (ui UI) {
	ui = UI{
		CurrentView: SettingsView,
	}
	return
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
	if ui.CurrentView == SettingsView {
		for idx := range ui.SettingsButtons {
			err := ui.SettingsButtons[idx].Update()
			if err != nil {
				return err
			}
		}
		for idx := range ui.Checkboxes {
			err := ui.Checkboxes[idx].Update()
			if err != nil {
				return err
			}
		}
	}

	if ui.CurrentView == TimerView && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		ui.ChangeView()
	}

	return nil
}

func (ui *UI) ChangeView() {
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

func (ui *UI) ChangeTimerMode() {
	ui.isOneLineView = !ui.isOneLineView
}
