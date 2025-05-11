package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Checkbox struct {
	box       uiRect
	isChecked bool
}

type CheckboxArray [1]Checkbox

func createCheckbox() (checkboxes CheckboxArray) {
	checkboxes[0] = Checkbox{
		box: uiRect{vx + (valueWidth / 2) - (buttonWidth / 2),
			calcRowY(5),
			buttonWidth,
			buttonHeight,
		},
		isChecked: false,
	}
	return
}

func (ui *UI) drawCheckbox(screen *ebiten.Image) {
	cb := ui.Checkboxes[0]
	if cb.isChecked {
		vector.DrawFilledRect(screen, float32(cb.box.x), float32(cb.box.y+btnMy),
			float32(cb.box.w), float32(cb.box.h), PrimaryColor, true)
		vector.StrokeLine(screen, float32(cb.box.x+4), float32(cb.box.y+(cb.box.h/2)+btnMy),
			float32(cb.box.x+(cb.box.w/2)), float32(cb.box.y+cb.box.h),
			3, TextColor, true)
		vector.StrokeLine(screen, float32(cb.box.x+(cb.box.w/2)), float32(cb.box.y+cb.box.h),
			float32(cb.box.x+cb.box.w-4), float32(cb.box.y+btnMy+4),
			3, TextColor, true)
	} else {
		vector.StrokeRect(screen, float32(cb.box.x+(strokeWidth/2)), float32(cb.box.y+btnMy+(strokeWidth/2)),
			float32(cb.box.w-(strokeWidth/2)), float32(cb.box.h-(strokeWidth/2)), strokeWidth, PrimaryColor, true)
	}
}
