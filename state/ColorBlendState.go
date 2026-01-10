package state

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2/data/binding"
)

type ColorBlendState struct {
	Color   color.Color
	Mode    func(color.Color, color.Color) color.RGBA64
	Opacity binding.Float
}

func (cb *ColorBlendState) SetColor(blendColor color.Color) {
	cb.Color = blendColor
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
