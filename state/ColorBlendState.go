package state

import (
	"image/color"
	"math"
	"photo-man/core/image_io"

	"fyne.io/fyne/v2/data/binding"
)

type ColorBlendState struct {
	BlendColor *image_io.SystemColor
	Mode       func(color.Color, color.Color) color.RGBA64
	Opacity    binding.Float
}

func (cb *ColorBlendState) SetMode(mode func(color.Color, color.Color) color.RGBA64) {
	cb.Mode = mode
}
func (cb *ColorBlendState) SetOpacity(contrast float64) {
	cb.Opacity.Set(math.Max(0, math.Min(100, contrast)))
}

func (cb *ColorBlendState) initColorBlendingState() {
	cb.Opacity.Set(0)
}
