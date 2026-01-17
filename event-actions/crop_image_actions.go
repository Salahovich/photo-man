package event_actions

import (
	"image"
	"image/color"
	"photo-man/core/image_transform"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
)

func InitCropImageCanvas(st *state.AppState) *customUI.ResizableRectangle {

	w, h := float32(st.CanvasState.GetCurrentImage().Bounds().Dx()), float32(st.CanvasState.GetCurrentImage().Bounds().Dy())
	cropRectangle := customUI.NewResizableRectangle(fyne.NewSize(w, h), color.Transparent)


	st.CanvasState.AddCropLayer(cropRectangle)

	return cropRectangle
}

func CropImageAction(img image.Image, cs *state.CropState) image.Image {
	return image_transform.Crop(img, cs.GetStartPosition(), cs.GetEndPosition())
}

func RemoveCropImageCanvas(st *state.AppState) {
	st.CanvasState.RemoveCropLayer()
}
