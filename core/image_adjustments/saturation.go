package image_adjustments

import (
	"image"
	"image/color"
	"math"
	"sync"
)

func UpdateSaturation(img image.Image, scaler float64) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := adjustSaturationPixel(img.At(x, y), scaler)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func adjustSaturationPixel(rgba color.Color, saturationFactor float64) color.RGBA64 {
	r, g, b, a := rgba.RGBA()
	r_, g_, b_, a_ := float64(r), float64(g), float64(b), float64(a)
	grey := r_*0.2126 + g_*0.7152 + b_*0.0722
	r_ = math.Max(0, math.Min(65535, r_+(r_-grey)*(saturationFactor)))
	g_ = math.Max(0, math.Min(65535, g_+(g_-grey)*(saturationFactor)))
	b_ = math.Max(0, math.Min(65535, b_+(b_-grey)*(saturationFactor)))
	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}
