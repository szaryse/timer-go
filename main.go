package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/szaryse/timer-go255/app"
	"github.com/szaryse/timer-go255/ui"
	"log"
)

func main() {
	ebiten.SetWindowSize(ui.SettingsWidth, ui.SettingsHeight)
	ebiten.SetWindowTitle("Timer Go v0.1")
	ebiten.SetWindowFloating(true)
	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true

	newApp := app.NewApp()

	if err := ebiten.RunGameWithOptions(&newApp, op); err != nil {
		log.Fatal(err)
	}
}
