package state

import (
	"math"

	"fyne.io/fyne/v2/data/binding"
)

type AdjustmentState struct {
	Brightness binding.Float
	Contrast   binding.Float
	Saturation binding.Float
}

func (as *AdjustmentState) SetBrightness(brightness float64) {
	as.Brightness.Set(math.Max(0, math.Min(100, brightness)))
}
func (as *AdjustmentState) SetSaturation(saturation float64) {
	as.Saturation.Set(math.Max(0, math.Min(100, saturation)))
}
func (as *AdjustmentState) SetContrast(contrast float64) {
	as.Contrast.Set(math.Max(0, math.Min(100, contrast)))
}

func (as *AdjustmentState) InitAdjustmentsState() {
	as.Brightness.Set(50)
	as.Contrast.Set(50)
	as.Saturation.Set(50)
}
