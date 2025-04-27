package app

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/szaryse/timer-go255/timer"
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

type App struct {
	timer timer.Timer
}

func NewApp() App {
	timer.NewTimer()
	return App{}
}

func (app *App) Update() error {
	err := app.timer.Update()
	if err != nil {
		return err
	}
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	app.timer.Render(screen, firaCodeSource, dotoSource)
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
