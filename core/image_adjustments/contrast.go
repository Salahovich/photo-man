package image_adjustments

import (
	"image"
	"image/color"
	"math"
	"sync"
)

// contrast level ranges from -255.0 to 255.0
var (
	contrastIncreaseLevel  = 4.0
	contrastIncreaseFactor = (259.0 * (contrastIncreaseLevel + 255)) / (255.0 * (259 - contrastIncreaseLevel))
	contrastDecreaseLevel  = -4.0
	contrastDecreaseFactor = (259.0 * (contrastDecreaseLevel + 255)) / (255.0 * (259 - contrastDecreaseLevel))
)

func IncreaseContrast(img image.Image) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustContrastPixel(img.At(x, y), contrastIncreaseFactor)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func DecreaseContrast(img image.Image) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustContrastPixel(img.At(x, y), contrastDecreaseFactor)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func adjustContrastPixel(rgba color.Color, contrastFactor float64) color.RGBA64 {
	r, g, b, a := rgba.RGBA()
	r_, g_, b_, a_ := float64(r), float64(g), float64(b), float64(a)
	r_ = math.Max(0, math.Min(65535, contrastFactor*(r_-32768)+32768))
	g_ = math.Max(0, math.Min(65535, contrastFactor*(g_-32768)+32768))
	b_ = math.Max(0, math.Min(65535, contrastFactor*(b_-32768)+32768))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}
