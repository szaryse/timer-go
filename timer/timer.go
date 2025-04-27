package timer

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
)

type Timer struct {
	count int
}

func NewTimer() Timer {
	return Timer{}
}

func (t *Timer) Render(screen *ebiten.Image, firaCodeSource *text.GoTextFaceSource, dotoSource *text.GoTextFaceSource) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(50, 60)
	op.ColorScale.ScaleWithColor(color.White)
	time := fmt.Sprintf("%d", t.count/60)
	text.Draw(screen, time, &text.GoTextFace{
		Source: firaCodeSource,
		Size:   24,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(0, 0)
	op.ColorScale.ScaleWithColor(color.White)
	status := "TEST"
	text.Draw(screen, status, &text.GoTextFace{
		Source: dotoSource,
		Size:   24,
	}, op)
}

func (t *Timer) Update() error {
	t.count += 1
	return nil
}
