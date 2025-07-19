package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type CheckBox struct {
	box            uiRect
	isChecked      bool
	color          color.RGBA
	mouseDown      bool
	onCheckChanged func(c *CheckBox)
}

func (c *CheckBox) SetOnCheckChanged(f func(c *CheckBox)) {
	c.onCheckChanged = f
}

func NewCheckbox() CheckBox {
	return CheckBox{
		box: uiRect{vx + (valueWidth / 2) - (buttonWidth / 2),
			calcRowY(5),
			buttonWidth,
			buttonHeight,
		},
		isChecked: false,
		color:     PrimaryColor,
	}
}

func (c *CheckBox) Update() error {
	cursorX, cursorY := ebiten.CursorPosition()
	isSelected := checkElementIsSelected(cursorX, cursorY, &c.box)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		c.mouseDown = isSelected
	} else {
		if c.mouseDown {
			c.isChecked = !c.isChecked
			if c.onCheckChanged != nil {
				c.onCheckChanged(c)
			}
		}
		c.mouseDown = false
	}

	if isSelected {
		c.color = SelectedColor
	} else {
		c.color = PrimaryColor
	}

	return nil
}

func (ui *UI) drawCheckbox(screen *ebiten.Image) {
	cb := ui.Checkboxes[0]
	if cb.isChecked {
		ui.drawFilledBackground(screen, &cb)
		ui.drawCheckboxLines(screen, &cb)
	} else {
		ui.drawUncheckedCheckbox(screen, &cb)
	}
}

func (ui *UI) drawFilledBackground(screen *ebiten.Image, cb *CheckBox) {
	path := createRoundedPath(float32(cb.box.w), float32(cb.box.h))
	ui.vertices, ui.indices = path.AppendVerticesAndIndicesForFilling(ui.vertices[:0], ui.indices[:0])
	ui.setVertices(cb.color, cb.box.x, cb.box.y)
	op := createTrianglesOptions()
	screen.DrawTriangles(ui.vertices, ui.indices, whiteSubImage, op)
}

func (ui *UI) drawCheckboxLines(screen *ebiten.Image, cb *CheckBox) {
	path := createPath()
	ops := createCheckboxStrokeOptions()
	ui.vertices, ui.indices = path.AppendVerticesAndIndicesForStroke(ui.vertices[:0], ui.indices[:0], ops)
	ui.setVertices(color.Black, cb.box.x, cb.box.y)
	op := createTrianglesOptions()
	screen.DrawTriangles(ui.vertices, ui.indices, whiteSubImage, op)
}

func (ui *UI) drawUncheckedCheckbox(screen *ebiten.Image, cb *CheckBox) {
	path := createRoundedPath(float32(cb.box.w), float32(cb.box.h))
	ops := createStrokeOptions()
	ui.vertices, ui.indices = path.AppendVerticesAndIndicesForStroke(ui.vertices[:0], ui.indices[:0], ops)
	ui.setVertices(cb.color, cb.box.x, cb.box.y)
	op := createTrianglesOptions()
	screen.DrawTriangles(ui.vertices, ui.indices, whiteSubImage, op)
}

func createPath() (path vector.Path) {
	path.MoveTo(6, 13+btnMy)
	path.LineTo(14, 19+btnMy)
	path.LineTo(23, 7+btnMy)
	return
}

func createCheckboxStrokeOptions() (ops *vector.StrokeOptions) {
	ops = &vector.StrokeOptions{}
	ops.Width = 4
	ops.LineJoin = vector.LineJoinMiter
	ops.LineCap = vector.LineCapRound
	return
}
