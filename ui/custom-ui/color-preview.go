package customUI

import (
	"image/color"
	"photo-man/core/image_io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type CustomColorPicker struct {
	widget.BaseWidget
	systemColor *image_io.SystemColor
	OnClick     func(color.Color)
}

func NewCustomColorPicker(size fyne.Size, SystemColor *image_io.SystemColor, onClick func(color.Color)) *CustomColorPicker {
	var item *CustomColorPicker

	item = &CustomColorPicker{
		systemColor: SystemColor,
		OnClick:     onClick,
	}
	item.systemColor.Rect.SetMinSize(size)
	item.systemColor.Rect.CornerRadius = 4
	item.ExtendBaseWidget(item)

	return item
}
func (cp *CustomColorPicker) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (cp *CustomColorPicker) Tapped(*fyne.PointEvent) {
	picker := dialog.NewColorPicker("Select Color", "Choose a color", func(c color.Color) {
		cp.systemColor.UpdateColor(c)
		cp.Refresh()
	}, fyne.CurrentApp().Driver().AllWindows()[0])
	picker.Advanced = true
	picker.Show()

}

func (cp *CustomColorPicker) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(&cp.systemColor.Rect)
}

func (cp *CustomColorPicker) MinSize() fyne.Size {
	return cp.systemColor.Rect.MinSize()
}
