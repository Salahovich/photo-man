package customUI

import (
	"image"
	"image/color"
	"photo-man/core/image_io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type CustomImageCanvas struct {
	widget.BaseWidget
	canvas       *canvas.Image
	spottedColor *image_io.SystemColor
	eyeDropState bool
}

func NewCustomImageCanvas(spottedColor *image_io.SystemColor) *CustomImageCanvas {
	var item *CustomImageCanvas

	item = &CustomImageCanvas{
		canvas:       canvas.NewImageFromResource(nil),
		spottedColor: spottedColor,
		eyeDropState: false,
	}

	item.canvas.FillMode = canvas.ImageFillContain

	item.ExtendBaseWidget(item)

	return item
}

func (cim *CustomImageCanvas) SetImageInCanvas(img image.Image) {
	cim.canvas.Image = img
	cim.canvas.SetMinSize(fyne.Size{Width: float32(img.Bounds().Dx()), Height: float32(img.Bounds().Dy())})
	cim.canvas.Refresh()
}

func (cim *CustomImageCanvas) GettImageInCanvas() image.Image {
	return cim.canvas.Image
}

func (cim *CustomImageCanvas) SetSpottedColor(color *image_io.SystemColor) {
	cim.spottedColor = color
}

func (cim *CustomImageCanvas) GetSpottedColor() color.Color {
	return cim.spottedColor.Color
}

func (cim *CustomImageCanvas) changeSpottedColor(color color.Color) {
	cim.spottedColor.UpdateColor(color)
}

func (cim *CustomImageCanvas) Tapped(ev *fyne.PointEvent) {
	if cim.eyeDropState {
		pos := ev.Position
		r, g, b, a := cim.canvas.Image.At(int(pos.X), int(pos.Y)).RGBA()
		cim.changeSpottedColor(color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
	}
}

func (cim *CustomImageCanvas) ToggleEyeDropState() {
	cim.eyeDropState = !cim.eyeDropState
}

func (cim *CustomImageCanvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(cim.canvas)
}

func (cim *CustomImageCanvas) ClearCanvas() {
	cim.canvas.Image = nil
	cim.canvas.Refresh()
}
func (cim *CustomImageCanvas) MinSize() fyne.Size {
	return cim.canvas.MinSize()
}
