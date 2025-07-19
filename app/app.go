package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szaryse/timer-go255/timer"
	"github.com/szaryse/timer-go255/ui"
)

type App struct {
	timer timer.Timer
	ui    ui.UI
	exit  bool
}

func (app *App) Update() error {
	if app.exit {
		return ebiten.Termination
	}
	if err := app.ui.Update(); err != nil {
		return err
	}
	if app.timer.Activity == timer.Init && app.timer.IsRunning {
		app.timer.Activity = timer.StartingInState
		app.ui.ChangeView()
	}
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

func NewApp() (app *App) {
	app = &App{}
	app.timer = timer.NewTimer()
	app.ui = ui.CreateUI()
	app.SetButtons()
	app.SetCheckboxes()
	return
}

func (app *App) SetButtons() {
	app.ui.SettingsButtons = ui.CreateButtons()
	app.ui.SettingsButtons[0].SetOnPressed(func(b *ui.Button) {
		app.timer.IncreaseStart()
	})
	app.ui.SettingsButtons[1].SetOnPressed(func(b *ui.Button) {
		app.timer.DecreaseStart()
	})
	app.ui.SettingsButtons[2].SetOnPressed(func(b *ui.Button) {
		app.timer.IncreaseSession()
	})
	app.ui.SettingsButtons[3].SetOnPressed(func(b *ui.Button) {
		app.timer.DecreaseSession()
	})
	app.ui.SettingsButtons[4].SetOnPressed(func(b *ui.Button) {
		app.timer.IncreaseFocus()
	})
	app.ui.SettingsButtons[5].SetOnPressed(func(b *ui.Button) {
		app.timer.DecreaseFocus()
	})
	app.ui.SettingsButtons[6].SetOnPressed(func(b *ui.Button) {
		app.timer.IncreaseBreak()
	})
	app.ui.SettingsButtons[7].SetOnPressed(func(b *ui.Button) {
		app.timer.DecreaseBreak()
	})
	app.ui.SettingsButtons[8].SetOnPressed(func(b *ui.Button) {
		app.exit = true
	})
	app.ui.SettingsButtons[9].SetOnPressed(func(b *ui.Button) {
		if app.timer.IsRunning == false {
			app.timer.HandleStart()
		}
	})
	app.ui.SettingsButtons[10].SetOnPressed(func(b *ui.Button) {
		if app.timer.Activity != timer.Init {
			app.ui.ChangeView()
		}
	})

}

func (app *App) SetCheckboxes() {
	app.ui.Checkboxes = ui.CreateCheckboxes()
	app.ui.Checkboxes[0].SetOnCheckChanged(func(cb *ui.CheckBox) {
		app.ui.ChangeTimerMode()
	})
}
