package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szaryse/timer-go255/app"
	"github.com/szaryse/timer-go255/ui"
	"log"
)

func main() {
	ebiten.SetWindowSize(ui.SettingsWidth, ui.SettingsHeight)
	ebiten.SetWindowTitle("Timer 2 v1.0")
	ebiten.SetWindowFloating(true)
	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	if err := ebiten.RunGameWithOptions(app.NewApp(), op); err != nil {
		log.Fatal(err)
	}
}
