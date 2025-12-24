package image_adjustments

import (
	"image"
	"image/color"
)

var (
	brightnessValue int32 = 3
)

func IncreaseBrightness(img image.Image) image.Image {
	newImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			newColor := adjustPixel(img.At(x, y), brightnessValue)
			newImg.Set(x, y, newColor)
		}
	}
	return newImg
}

func DecreaseBrightness(img image.Image) image.Image {
	newImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			newColor := adjustPixel(img.At(x, y), -brightnessValue)
			newImg.Set(x, y, newColor)
		}
	}
	return newImg
}

func adjustPixel(rgba color.Color, brightnessFactor int32) color.RGBA {
	r, g, b, a := rgba.RGBA()
	r_, g_, b_, a_ := int32(r>>8), int32(g>>8), int32(b>>8), int32(a>>8)
	r_, g_, b_ = r_+brightnessFactor, g_+brightnessFactor, b_+brightnessFactor
	if r_ > 255 {
		r_ = 255
	}
	if r_ < 0 {
		r_ = 0
	}
	if g_ > 255 {
		g_ = 255
	}
	if g_ < 0 {
		g_ = 0
	}
	if b_ > 255 {
		b_ = 255
	}
	if b_ < 0 {
		b_ = 0
	}

	return color.RGBA{R: uint8(r_), G: uint8(g_), B: uint8(b_), A: uint8(a_)}
}
