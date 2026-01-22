package event_actions

import (
	"image"
	"photo-man/core/image_paint"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
)

func InitBlurBoardCanvas(st *state.AppState) *customUI.BlurBoard {

	w, h := float32(st.CanvasState.GetCurrentImage().Bounds().Dx()), float32(st.CanvasState.GetCurrentImage().Bounds().Dy())
	blurBoard := customUI.NewBlurBoard(fyne.NewSize(w, h), st.CanvasState.GetCurrentImage())

	st.CanvasState.AddBlurBoardLayer(blurBoard)

	return blurBoard
}

func BlurBoardAction(img image.Image, cs *state.BlurBoardState) image.Image {
	return image_paint.BrushAction(img, cs.GetBlurBoardCanvas().GetBoard())
}

func RemoveBlurBoardanvas(st *state.AppState) {
	st.CanvasState.RemoveBlurBoardLayer()
}
