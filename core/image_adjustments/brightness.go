package image_adjustments

import (
	"image"
	"image/color"
	"math"
	"sync"
)

var (
	brightnessValue = 1000.0
)

func IncreaseBrightness(img image.Image) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustBrightnessPixel(img.At(x, y), brightnessValue)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func DecreaseBrightness(img image.Image) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustBrightnessPixel(img.At(x, y), -brightnessValue)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func adjustBrightnessPixel(rgba color.Color, brightnessFactor float64) color.RGBA64 {
	r, g, b, a := rgba.RGBA()
	r_, g_, b_, a_ := float64(r), float64(g), float64(b), float64(a)
	r_ = math.Max(0, math.Min(65535, brightnessFactor+r_))
	g_ = math.Max(0, math.Min(65535, brightnessFactor+g_))
	b_ = math.Max(0, math.Min(65535, brightnessFactor+b_))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}
