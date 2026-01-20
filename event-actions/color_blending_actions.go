package event_actions

import (
	"image"
	"image/color"
	colorBlending "photo-man/core/color_blending"
	"photo-man/state"
)

func UpdateColorBlendingAction(img image.Image, cbs *state.ColorBlendState) image.Image {

	if opacity, _ := cbs.Opacity.Get(); opacity != 0 && cbs.Mode != nil {
		opacityNorm := (opacity / 100.0) * 65535.0
		r, g, b, _ := cbs.BlendColor.Color.RGBA()
		blendColor := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(opacityNorm)}

		img = colorBlending.PerformBlending(img, blendColor, cbs.Mode)
	}

	return img
}
