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
	newImg := image.NewRGBA(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustContrastPixel(img.At(x, y), contrastIncreaseFactor)
				newImg.Set(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func DecreaseContrast(img image.Image) image.Image {
	newImg := image.NewRGBA(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustContrastPixel(img.At(x, y), contrastDecreaseFactor)
				newImg.Set(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func adjustContrastPixel(rgba color.Color, contrastFactor float64) color.RGBA {
	r, g, b, a := rgba.RGBA()
	r_, g_, b_, a_ := float64(r>>8), float64(g>>8), float64(b>>8), float64(a>>8)
	r_ = math.Max(0, math.Min(255, contrastFactor*(r_-128)+128))
	g_ = math.Max(0, math.Min(255, contrastFactor*(g_-128)+128))
	b_ = math.Max(0, math.Min(255, contrastFactor*(b_-128)+128))

	return color.RGBA{R: uint8(r_), G: uint8(g_), B: uint8(b_), A: uint8(a_)}
}
