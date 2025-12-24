package state

import "image"

type AdjustmentState struct {
	Brightness [100]image.Image
	Contrast   [100]image.Image
	Saturation [100]image.Image
}

func (as *AdjustmentState) AddBrightnessValue(value int, img image.Image) {
	value--
	if value < 0 {
		value = 0
	}
	as.Brightness[value] = img
}
func (as *AdjustmentState) GetBrightnessValue(value int) image.Image {
	value--
	if value < 0 {
		value = 0
	}
	return as.Brightness[value]
}

func (as *AdjustmentState) AddContrastValue(value int, img image.Image) {
	value--
	if value < 0 {
		value = 0
	}
	as.Contrast[value] = img
}
func (as *AdjustmentState) GetContrastValue(value int) image.Image {
	value--
	if value < 0 {
		value = 0
	}
	return as.Contrast[value]
}

func (as *AdjustmentState) AddSaturationValue(value int, img image.Image) {
	value--
	if value < 0 {
		value = 0
	}
	as.Saturation[value] = img
}
func (as *AdjustmentState) GetSaturationValue(value int) image.Image {
	value--
	if value < 0 {
		value = 0
	}
	return as.Saturation[value]
}

func (as *AdjustmentState) ResetAdjustments() {
	as.Brightness = [100]image.Image{}
	as.Contrast = [100]image.Image{}
	as.Saturation = [100]image.Image{}
}
