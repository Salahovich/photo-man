package gray_filters

import (
	"image"
	"image/color"
)

func BasicGreyScale(oldImage image.Image) *image.Gray {
	greyImage := image.NewGray(oldImage.Bounds())
	for y := 0; y < oldImage.Bounds().Dy(); y++ {
		for x := 0; x < oldImage.Bounds().Dx(); x++ {
			red, green, blue, _ := oldImage.At(x, y).RGBA()
			pixelValueAvg := uint8(((red + green + blue) / 3) >> 8)
			greyImage.Set(x, y, color.Gray{Y: pixelValueAvg})
		}
	}
	return greyImage
}
