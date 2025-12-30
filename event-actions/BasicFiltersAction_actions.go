package event_actions

import (
	"photo-man/core/image_adjustments"
	"photo-man/state"
)

func UpdateBasicFilters(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}

	adj := st.AdjustmentState
	newImg := st.CanvasState.GetScaledImage()

	if brightness, _ := adj.Brightness.Get(); brightness != 50.0 {
		scaler := (brightness - 50.0) * st.AdjustmentFactors.BaseBrightnessFactor
		st.AdjustmentFactors.BrightnessFactor = scaler
		newImg = image_adjustments.UpdateBrightness(newImg, scaler)
	}

	if contrast, _ := adj.Contrast.Get(); contrast != 50.0 {
		correctionFactor := (contrast - 50.0) * st.AdjustmentFactors.BaseContrastFactor
		scaler := (259.0 * (correctionFactor + 255.0)) / (255.0 * (259.0 - correctionFactor))
		st.AdjustmentFactors.ContrastFactor = scaler
		newImg = image_adjustments.UpdateContrast(newImg, scaler)
	}

	if saturation, _ := adj.Saturation.Get(); saturation != 50.0 {
		scaler := (saturation - 50.0) * st.AdjustmentFactors.BaseSaturationFactor
		st.AdjustmentFactors.SaturationFactor = scaler
		newImg = image_adjustments.UpdateSaturation(newImg, scaler)
	}

	st.CanvasState.UpdateSceneImage(newImg)
}
