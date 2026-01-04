package event_actions

import (
	"image"
	"photo-man/core/image_adjustments"
	"photo-man/state"
)

func UpdateAdjustmentsAction(newImg image.Image, adj *state.AdjustmentState, adjFactor *state.AdjustmentFactors) image.Image {

	if brightness, _ := adj.Brightness.Get(); brightness != 50.0 {
		scaler := (brightness - 50.0) * adjFactor.BaseBrightnessFactor
		adjFactor.BrightnessFactor = scaler
		newImg = image_adjustments.UpdateBrightness(newImg, scaler)
	}

	if contrast, _ := adj.Contrast.Get(); contrast != 50.0 {
		correctionFactor := (contrast - 50.0) * adjFactor.BaseContrastFactor
		scaler := (259.0 * (correctionFactor + 255.0)) / (255.0 * (259.0 - correctionFactor))
		adjFactor.ContrastFactor = scaler
		newImg = image_adjustments.UpdateContrast(newImg, scaler)
	}

	if saturation, _ := adj.Saturation.Get(); saturation != 50.0 {
		scaler := (saturation - 50.0) * adjFactor.BaseSaturationFactor
		adjFactor.SaturationFactor = scaler
		newImg = image_adjustments.UpdateSaturation(newImg, scaler)
	}

	return newImg
}
