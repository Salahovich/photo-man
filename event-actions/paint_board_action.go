package event_actions

import (
	"image"
	"image/color"
	"photo-man/core/image_paint"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
)

func InitPaintBoardCanvas(st *state.AppState) *customUI.PaintBoard {

	w, h := float32(st.CanvasState.GetCurrentImage().Bounds().Dx()), float32(st.CanvasState.GetCurrentImage().Bounds().Dy())
	paintBoard := customUI.NewPaintBoard(fyne.NewSize(w, h), color.Black)

	st.CanvasState.AddPaintBoardLayer(paintBoard)

	return paintBoard
}

func PaintBoardAction(img image.Image, cs *state.PaintBoardState) image.Image {
	return image_paint.Brush(img, cs.GetPaintBoardCanvas().GetBoard())
}

func RemovePaintBoardCanvas(st *state.AppState) {
	st.CanvasState.RemovePaintBoardLayer()
}
