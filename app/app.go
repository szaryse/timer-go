package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szaryse/timer-go255/timer"
	"github.com/szaryse/timer-go255/ui"
)

type App struct {
	timer timer.Timer
	ui    ui.UI
}

func NewApp() App {
	return App{
		timer: timer.NewTimer(),
		ui:    ui.CreateUI(),
	}
}

func (app *App) Update() error {
	if err := app.ui.Update(); err != nil {
		return err
	}
	// todo
	// action := checkUiAction
	// changeTimerState(action)
	if err := app.timer.Update(); err != nil {
		return err
	}
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	app.ui.Render(screen, &app.timer)
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
