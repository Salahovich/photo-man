package state

import (
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
)

type CropState struct {
	isCropState     bool
	cropImageCanvas *customUI.ResizableRectangle
}

func (cs *CropState) EnableCropState() {
	cs.isCropState = true
}

func (cs *CropState) DisableCropState() {
	cs.isCropState = false
}

func (cs *CropState) IsInCropState() bool {
	return cs.isCropState
}

func (cs *CropState) SetCropImageCanvas(rectangle *customUI.ResizableRectangle) {
	cs.cropImageCanvas = rectangle
}

func (cs *CropState) GetStartPosition() fyne.Position {
	return cs.cropImageCanvas.GetCurrStartPosition()
}
func (cs *CropState) GetEndPosition() fyne.Position {
	return cs.cropImageCanvas.GetCurrEndPosition()
}
func (cs *CropState) CanCrop() bool {
	return cs.cropImageCanvas.CanCrop()
}
func (cs *CropState) InitCropState() {
	cs.isCropState = false
}
