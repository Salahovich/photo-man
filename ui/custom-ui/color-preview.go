package customUI

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type CustomColorPicker struct {
	widget.BaseWidget
	Rect    *canvas.Rectangle
	OnClick func(color.Color)
}

func NewCustomColorPicker(size fyne.Size, onClick func(color.Color)) *CustomColorPicker {
	var item *CustomColorPicker

	item = &CustomColorPicker{
		Rect:         canvas.NewRectangle(color.Black),
		OnClick:      onClick,
	}
	item.Rect.SetMinSize(size)
	item.Rect.CornerRadius = 4
	item.ExtendBaseWidget(item)

	return item
}

func (cp *CustomColorPicker) Tapped(*fyne.PointEvent) {
	picker := dialog.NewColorPicker("Select Color", "Choose a color", func(c color.Color) {
		cp.Rect.FillColor = c
		cp.OnClick(c)
		cp.Refresh()
	}, fyne.CurrentApp().Driver().AllWindows()[0])
	picker.Advanced = true
	picker.Show()

}

func (cp *CustomColorPicker) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(cp.Rect)
}

func (cp *CustomColorPicker) MinSize() fyne.Size {
	return cp.Rect.MinSize()
}
