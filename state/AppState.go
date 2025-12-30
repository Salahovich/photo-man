package state

import (
	"image"
	"photo-man/core/image_adjustments"
	"sync"

	"fyne.io/fyne/v2/data/binding"
)

type AppState struct {
	CanvasState       *CanvasState
	AdjustmentState   *AdjustmentState
	AdjustmentFactors *AdjustmentFactors
	transformations   []func(image.Image) image.Image
}

func NewAppState() *AppState {
	newState := AppState{
		CanvasState: &CanvasState{
			communication: make(chan image.Image),
			canvasMutex:   &sync.RWMutex{},
		},
		AdjustmentState: &AdjustmentState{
			Brightness: binding.NewFloat(),
			Contrast:   binding.NewFloat(),
			Saturation: binding.NewFloat(),
		},
		AdjustmentFactors: &AdjustmentFactors{
			BaseBrightnessFactor: 1000.0,
			BaseContrastFactor:   5.1,
			BaseSaturationFactor: 2.0 / 50,
			BrightnessFactor:     0.0,
			ContrastFactor:       1.0,
			SaturationFactor:     0.0,
		},
		transformations: make([]func(image.Image) image.Image, 0),
	}
	newState.AdjustmentState.InitAdjustmentsState()

	return &newState
}

func (s *AppState) RegisterTransformation(callback func(image.Image) image.Image) {
	s.transformations = append(s.transformations, callback)
}

func (s *AppState) ResetTransformation() {
	s.transformations = make([]func(image2 image.Image) image.Image, 0)
}

func (s *AppState) ApplyAllModification() {
	s.CanvasState.GetCanvasMutex().Lock()
	defer s.CanvasState.GetCanvasMutex().Unlock()

	s.CanvasState.originalImage = image_adjustments.UpdateBrightness(s.CanvasState.originalImage, s.AdjustmentFactors.BrightnessFactor)
	s.CanvasState.originalImage = image_adjustments.UpdateContrast(s.CanvasState.originalImage, s.AdjustmentFactors.ContrastFactor)
	s.CanvasState.originalImage = image_adjustments.UpdateSaturation(s.CanvasState.originalImage, s.AdjustmentFactors.SaturationFactor)

	for _, callback := range s.transformations {
		s.CanvasState.originalImage = callback(s.CanvasState.originalImage)
	}
}

func (s *AppState) Reset() {
	s.AdjustmentState.InitAdjustmentsState()
	s.AdjustmentFactors.InitAdjustmentsFactors()
}
