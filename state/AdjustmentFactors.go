package state

type AdjustmentFactors struct {
	BrightnessFactor     float64
	ContrastFactor       float64
	SaturationFactor     float64
	BaseBrightnessFactor float64
	BaseContrastFactor   float64
	BaseSaturationFactor float64
}

func (adjf *AdjustmentFactors) InitAdjustmentsFactors() {
	adjf.BrightnessFactor = 0
	adjf.ContrastFactor = 1
	adjf.SaturationFactor = 0
}
