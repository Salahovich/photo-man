package state

import "math"

type AdjustmentState struct {
	Brightness float64
	Contrast   float64
	Saturation float64
}

func (as *AdjustmentState) SetBrightness(brightness float64) {
	as.Brightness = math.Max(0, math.Min(100, brightness))
}
func (as *AdjustmentState) SetSaturation(saturation float64) {
	as.Saturation = math.Max(0, math.Min(100, saturation))
}
func (as *AdjustmentState) SetContrast(contrast float64) {
	as.Contrast = math.Max(0, math.Min(100, contrast))
}

func (as *AdjustmentState) ResetAdjustments() {
	as.Brightness = 50
	as.Contrast = 50
	as.Saturation = 50
}
