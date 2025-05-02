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
	TimerWidth     = 960
	TimerHeight    = 240
)

var (
	TextColor    = color.Gray{Y: 192}
	PrimaryColor = color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
)

type UI struct {
	CurrentView     ViewState
	WindowPositionX int
	WindowPositionY int
	SettingsButtons [12]Button
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
		ui.RenderSettingView(screen)
	}
	// todo create the timer view
	if ui.CurrentView == TimerView {
		op := &text.DrawOptions{}
		op.GeoM.Translate(50, 60)
		op.ColorScale.ScaleWithColor(TextColor)
		time := fmt.Sprintf("%d", t.Count/60)
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
			ex, ey := ebiten.CursorPosition()

			for idx, button := range ui.SettingsButtons {
				if ex > button.x &&
					ex < button.x+button.w &&
					ey > button.y &&
					ey < button.y+button.h {
					fmt.Println(idx, button.action)
					// todo
					//switch idx {
					//case 10:
					//	fmt.Println("exit")
					//case 11:
					//	fmt.Println("start")
					//	ui.changeView()
					//}
				}
			}
		}
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
