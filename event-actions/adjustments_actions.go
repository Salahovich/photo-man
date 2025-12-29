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

func IncreaseContrastAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetContrastValue(value)
	if img == nil {
		img = image_adjustments.IncreaseContrast(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddContrastValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.IncreaseContrast)
	st.CanvasState.UpdateSceneImage(img)
}

func DecreaseContrastAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetContrastValue(value)
	if img == nil {
		img = image_adjustments.DecreaseContrast(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddContrastValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.DecreaseContrast)
	st.CanvasState.UpdateSceneImage(img)
}

func IncreaseSaturationAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetSaturationValue(value)
	if img == nil {
		img = image_adjustments.IncreaseSaturation(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddSaturationValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.IncreaseSaturation)
	st.CanvasState.UpdateSceneImage(img)
}

func DecreaseSaturationAction(st *state.AppState, value int) {
	img := st.AdjustmentState.GetSaturationValue(value)
	if img == nil {
		img = image_adjustments.DecreaseSaturation(st.CanvasState.GetCurrentImage())
		st.AdjustmentState.AddSaturationValue(value, img)
	}
	st.CanvasState.RegisterModification(image_adjustments.DecreaseSaturation)
	st.CanvasState.UpdateSceneImage(img)
}
