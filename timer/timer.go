package timer

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Timer struct {
	count int
}

func NewTimer() Timer {
	return Timer{}
}

func (t *Timer) Render(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", t.count/60))
}

func (t *Timer) Update() error {
	t.count += 1
	return nil
}
