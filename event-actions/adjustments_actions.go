package event_actions

import (
	"photo-man/core/image_adjustments"
	"photo-man/state"
)

func UpdateAdjustments(st *state.AppState) {
	adj := st.AdjustmentState
	newImg := st.CanvasState.GetScaledImage()

	if adj.Brightness != 50.0 {
		scaler := (adj.Brightness - 50.0) * st.AdjustmentFactors.BaseBrightnessFactor
		st.AdjustmentFactors.BrightnessFactor = scaler
		newImg = image_adjustments.UpdateBrightness(newImg, scaler)
	}

	if adj.Contrast != 50.0 {
		correctionFactor := (adj.Contrast - 50.0) * st.AdjustmentFactors.BaseContrastFactor
		scaler := (259.0 * (correctionFactor + 255.0)) / (255.0 * (259.0 - correctionFactor))
		st.AdjustmentFactors.ContrastFactor = scaler
		newImg = image_adjustments.UpdateContrast(newImg, scaler)
	}

	if adj.Saturation != 50.0 {
		scaler := (adj.Saturation - 50.0) * st.AdjustmentFactors.BaseSaturationFactor
		st.AdjustmentFactors.SaturationFactor = scaler
		newImg = image_adjustments.UpdateSaturation(newImg, scaler)
	}

	st.CanvasState.UpdateSceneImage(newImg)
}
