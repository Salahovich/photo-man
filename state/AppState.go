package state

import (
	"image"
	"image/color"
	colorBlending "photo-man/core/color_blending"
	"photo-man/core/image_adjustments"
	"photo-man/core/image_filters"
	"photo-man/core/image_io"
	"photo-man/core/image_paint"
	"photo-man/core/image_transform"
	customUI "photo-man/ui/custom-ui"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type AppState struct {
	CanvasState         *CanvasState
	AdjustmentState     *AdjustmentState
	AdjustmentFactors   *AdjustmentFactors
	BasicFilterState    *BasicFilterState
	ColorBlendState     *ColorBlendState
	Transformations     *TransformationState
	ToolDialogContainer *fyne.Container
	AppEdgeContainers   []*fyne.Container
	AppWindow           fyne.Window
	SystemColor         *image_io.SystemColor
}

func NewAppState(window fyne.Window) *AppState {
	newState := AppState{
		CanvasState: &CanvasState{
			communication: make(chan image.Image),
			canvasMutex:   &sync.RWMutex{},
			cropState: &CropState{
				isCropState:     false,
				cropImageCanvas: &customUI.ResizableRectangle{},
			},
			paintBoardState: &PaintBoardState{
				isPaintState:     false,
				paintBoardCanvas: &customUI.PaintBoard{},
			},
			blurBoradState: &BlurBoardState{
				isBlurState:     false,
				blurBoardCanvas: &customUI.BlurBoard{},
			},
			sharpenBoradState: &SharpenBoardState{
				isSharpenState:     false,
				sharpenBoardCanvas: &customUI.SharpenBoard{},
			},
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
			Mode:    nil,
			Opacity: binding.NewFloat(),
		},
		Transformations: &TransformationState{
			Rotate:         0,
			FlipVertical:   false,
			FlipHorizontal: false,
		},
		SystemColor: &image_io.SystemColor{
			Color: color.Black,
		},
		AppWindow: window,
	}
	newState.ColorBlendState.BlendColor = newState.SystemColor

	newState.AdjustmentState.InitAdjustmentsState()

	return &newState
}

func (s *AppState) SetAppEdgeContainers(containers []*fyne.Container) {
	s.AppEdgeContainers = containers
}

func (s *AppState) SetToolDialogContainers(containers *fyne.Container) {
	s.ToolDialogContainer = containers
}

func (s *AppState) ShowToolDialog(dialog *fyne.Container) {
	s.ToolDialogContainer.Add(dialog)
	s.ToolDialogContainer.Refresh()
}
func (s *AppState) RemoveToolDialog() {
	s.ToolDialogContainer.RemoveAll()
}

func (s *AppState) ApplyAllModificationOnOriginalImage() image.Image {
	img := s.CanvasState.GetOriginalImage()

	// paint the image
	if s.CanvasState.GetPaintBoardState().CanPaint() {
		img = image_paint.BrushAction(img, s.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().GetBoard())
	}

	// blur the image
	if s.CanvasState.GetBlurBoardState().CanBlur() {
		img = image_paint.BrushAction(img, s.CanvasState.GetBlurBoardState().blurBoardCanvas.GetBoard())
	}

	// sharpen the image
	if s.CanvasState.GetSharpenBoardState().CanSharpen() {
		img = image_paint.BrushAction(img, s.CanvasState.GetSharpenBoardState().sharpenBoardCanvas.GetBoard())
	}

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

	// crop image
	if s.CanvasState.cropState.CanCrop() {
		img = image_transform.Crop(img, s.CanvasState.cropState.GetStartPosition(), s.CanvasState.cropState.GetEndPosition())
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
	img = colorBlending.PerformBlending(img, s.SystemColor.Color, s.ColorBlendState.Mode)

	return img
}

func (s *AppState) Reset() {
	s.AdjustmentState.InitAdjustmentsState()
	s.AdjustmentFactors.InitAdjustmentsFactors()
	s.BasicFilterState.InitBasicFilterState()
	s.Transformations.InitTransformations()
	s.ColorBlendState.initColorBlendingState()
	s.CanvasState.GetCropState().InitCropState()
}
