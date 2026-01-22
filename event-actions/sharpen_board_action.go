package event_actions

import (
	"image"
	"photo-man/core/image_paint"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
)

func InitSharpenBoardCanvas(st *state.AppState) *customUI.SharpenBoard {

	w, h := float32(st.CanvasState.GetCurrentImage().Bounds().Dx()), float32(st.CanvasState.GetCurrentImage().Bounds().Dy())
	sharpenBoard := customUI.NewSharpenBoard(fyne.NewSize(w, h), st.CanvasState.GetCurrentImage())

	st.CanvasState.AddSharpenBoardLayer(sharpenBoard)

	return sharpenBoard
}

func SharpenBoardAction(img image.Image, cs *state.SharpenBoardState) image.Image {
	return image_paint.BrushAction(img, cs.GetSharpenBoardCanvas().GetBoard())
}

func RemoveSharpenBoardanvas(st *state.AppState) {
	st.CanvasState.RemoveBlurBoardLayer()
}
