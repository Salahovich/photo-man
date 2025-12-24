package event_actions

import (
	"photo-man/core/image_adjustments"
	"photo-man/state"
)

func IncreaseBrightnessAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetBrightnessValue(value)
	if img == nil {
		img = image_adjustments.IncreaseBrightness(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddBrightnessValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.IncreaseBrightness)
	st.CanvasState.UpdateSceneImage(img)
}

func DecreaseBrightnessAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetBrightnessValue(value)
	if img == nil {
		img = image_adjustments.DecreaseBrightness(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddBrightnessValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.DecreaseBrightness)
	st.CanvasState.UpdateSceneImage(img)
}
