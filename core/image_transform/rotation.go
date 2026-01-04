package image_transform

import (
	"image"
	"image/color"
)

func RotateClockwise(img image.Image) image.Image {
	cols, rows := img.Bounds().Dx(), img.Bounds().Dy()
	newImg := image.NewRGBA64(image.Rect(0, 0, rows, cols))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			newImg.SetRGBA64(rows-(y+1), x, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}
	return newImg
}

func RotateAntiClockwise(img image.Image) image.Image {
	cols, rows := img.Bounds().Dx(), img.Bounds().Dy()
	newImg := image.NewRGBA64(image.Rect(0, 0, rows, cols))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			newImg.SetRGBA64(y, (cols-x)-1, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}

	return newImg
}
