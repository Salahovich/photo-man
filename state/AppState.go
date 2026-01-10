package state

import (
	"image"
	"image/color"
	colorBlending "photo-man/core/color_blending"
	"photo-man/core/image_adjustments"
	"photo-man/core/image_filters"
	"photo-man/core/image_transform"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type AppState struct {
	CanvasState       *CanvasState
	AdjustmentState   *AdjustmentState
	AdjustmentFactors *AdjustmentFactors
	BasicFilterState  *BasicFilterState
	ColorBlendState   *ColorBlendState
	Transformations   *TransformationState
	AppEdgeContainers []*fyne.Container
	AppWindow         fyne.Window
}

func NewAppState(window fyne.Window) *AppState {
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
		BasicFilterState: &BasicFilterState{
			Blur:       0,
			Emboss:     0,
			Outline:    0,
			Sharpening: 0,
		},
		AdjustmentFactors: &AdjustmentFactors{
			BaseBrightnessFactor: 1000.0,
			BaseContrastFactor:   5.1,
			BaseSaturationFactor: 2.0 / 50,
			BrightnessFactor:     0.0,
			ContrastFactor:       1.0,
			SaturationFactor:     0.0,
		},
		ColorBlendState: &ColorBlendState{
			Color:   color.Black,
			Mode:    nil,
			Opacity: binding.NewFloat(),
		},
		Transformations: &TransformationState{
			Rotate:         0,
			FlipVertical:   false,
			FlipHorizontal: false,
		},
		AppWindow: window,
	}
	newState.AdjustmentState.InitAdjustmentsState()

	return &newState
}

func (s *AppState) SetAppEdgeContainers(containers []*fyne.Container) {
	s.AppEdgeContainers = containers
}

func (s *AppState) ApplyAllModificationOnOriginalImage() image.Image {
	img := s.CanvasState.GetOriginalImage()

	// rotate the image
	if s.Transformations.Rotate < 0 {
		img = image_transform.RotateClockwise(img)
	} else if s.Transformations.Rotate > 0 {
		for range s.Transformations.Rotate {
			img = image_transform.RotateAntiClockwise(img)
		}
	}

	// flip the image
	if s.Transformations.FlipHorizontal {
		img = image_transform.FlipHorizontally(img)
	}
	if s.Transformations.FlipVertical {
		img = image_transform.FlipVertically(img)
	}

	// filter the image
	img = image_filters.GaussianBlur(img, s.BasicFilterState.Blur)
	img = image_filters.EmbossImage(img, s.BasicFilterState.Emboss)
	img = image_filters.OutlineImage(img, s.BasicFilterState.Outline)
	img = image_filters.SharpImage(img, s.BasicFilterState.Sharpening)
	img = image_filters.SobelImage(img, s.BasicFilterState.Sobel)

	// adjust the image
	img = image_adjustments.UpdateBrightness(img, s.AdjustmentFactors.BrightnessFactor)
	img = image_adjustments.UpdateContrast(img, s.AdjustmentFactors.ContrastFactor)
	img = image_adjustments.UpdateSaturation(img, s.AdjustmentFactors.SaturationFactor)

	// color blending
	img = colorBlending.PerformBlending(img, s.ColorBlendState.Color, s.ColorBlendState.Mode)

	return img
}

func (s *AppState) Reset() {
	s.AdjustmentState.InitAdjustmentsState()
	s.AdjustmentFactors.InitAdjustmentsFactors()
	s.BasicFilterState.InitBasicFilterState()
	s.Transformations.InitTransformations()
	s.ColorBlendState.initColorBlendingState()
}
