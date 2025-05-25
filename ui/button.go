package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
)

const (
	textBtnWidth = 100
	actionBtnY   = 205
	btnFontSize  = 18
	btnRx        = 397 // button right x
	btnLx        = 229 // button left x
)

type Button struct {
	box       uiRect
	label     string
	color     color.Color
	mouseDown bool
	onPressed func(b *Button)
}

func (b *Button) SetOnPressed(f func(b *Button)) {
	b.onPressed = f
}

func (b *Button) Update() error {
	cursorX, cursorY := ebiten.CursorPosition()
	isSelected := checkElementIsSelected(cursorX, cursorY, &b.box)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.mouseDown = isSelected
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}

	if isSelected {
		b.color = SelectedColor
	} else {
		b.color = PrimaryColor
	}

	return nil
}

func NewButton(x int, label string) Button {
	return Button{
		box: uiRect{
			x,
			0,
			0,
			buttonHeight,
		},
		label: label,
		color: PrimaryColor,
	}
}

func (ui *UI) renderButton(screen *ebiten.Image, idx int) {
	btn := ui.SettingsButtons[idx]
	path := createRoundedPath(float32(btn.box.w), float32(btn.box.h))
	ops := createStrokeOptions()
	ui.vertices, ui.indices = path.AppendVerticesAndIndicesForStroke(ui.vertices[:0], ui.indices[:0], ops)
	ui.setVertices(btn.color, btn.box.x, btn.box.y)
	op := createTrianglesOptions()
	screen.DrawTriangles(ui.vertices, ui.indices, whiteSubImage, op)

	renderButtonText(screen, float64(btn.box.x), float64(btn.box.y),
		float64(btn.box.w), btn.label)
}

func renderButtonText(screen *ebiten.Image, x, y, w float64, label string) {
	textFace := &text.GoTextFace{
		Source: firaCodeSource,
		Size:   btnFontSize,
	}
	labelWidth, labelHeight := text.Measure(label, textFace, 0)
	x += (w - labelWidth) / 2
	y += ((buttonHeight - labelHeight) / 2) + btnMy + (strokeWidth / 2)
	
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(TextColor)

	text.Draw(screen, label, textFace, op)
}
