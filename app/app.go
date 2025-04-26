package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szaryse/timer-go255/timer"
)

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
	app.timer.Render(screen)
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
