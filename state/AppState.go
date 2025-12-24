package state

import (
	"image"
)

type AppState struct {
	CanvasState     *CanvasState
	AdjustmentState *AdjustmentState
}

func NewAppState() *AppState {
	newState := AppState{
		CanvasState: &CanvasState{
			modifications: make([]func(image.Image) image.Image, 0),
			communication: make(chan image.Image),
		},
		AdjustmentState: &AdjustmentState{
			Brightness: [100]image.Image{},
			Contrast:   [100]image.Image{},
			Saturation: [100]image.Image{},
		},
	}
	newState.Init()

	return &newState
}

func (s *AppState) Reset() {
	s.CanvasState.ResetModifications()
	s.AdjustmentState.ResetAdjustments()
	s.Init()
}

func (s *AppState) Init() {
	s.AdjustmentState.Brightness[49] = s.CanvasState.currentImage
	s.AdjustmentState.Contrast[49] = s.CanvasState.currentImage
	s.AdjustmentState.Saturation[49] = s.CanvasState.currentImage
}
