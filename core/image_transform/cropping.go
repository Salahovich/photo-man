package image_transform

import (
	"image"
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
)

func Crop(img image.Image, start fyne.Position, end fyne.Position) image.Image {
	width := int(end.X - start.X)
	height := int(end.Y - start.Y)
	newImg := image.NewRGBA64(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}
	wg.Add(newImg.Bounds().Dy())
	for y, y_ := int(start.Y), 0; y_ < height; y, y_ = y+1, y_+1 {
		go func(y, y_ int) {
			defer wg.Done()
			for x, x_ := int(start.X), 0; x_ < width; x, x_ = x+1, x_+1 {
				r, g, b, a := img.At(x, y).RGBA()
				newImg.SetRGBA64(x_, y_, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
			}
		}(y, y_)
	}
	wg.Wait()
	return newImg
}
