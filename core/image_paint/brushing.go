package image_paint

import (
	"image"
	"image/color"
	"sync"
)

func Brush(current image.Image, paintBoard image.Image) image.Image {
	newImg := image.NewRGBA64(current.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(current.Bounds().Dy())
	for y := 0; y < current.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < current.Bounds().Dx(); x++ {
				r, g, b, a := paintBoard.At(x, y).RGBA()
				r_, g_, b_, a_ := current.At(x, y).RGBA()
				paintColor := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
				pixelColor := color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}

				if a == 0 {
					newImg.SetRGBA64(x, y, pixelColor)
				} else {
					newImg.SetRGBA64(x, y, paintColor)
				}
			}
		}(y)
	}
	wg.Wait()
	return newImg
}
