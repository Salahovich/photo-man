package image_transform

import (
	"image"
	"image/color"
)

func FlipHorizontally(img image.Image) image.Image {
	rows, cols := img.Bounds().Dy(), img.Bounds().Dx()
	newImg := image.NewRGBA64(img.Bounds())
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			newImg.SetRGBA64(cols-(x+1), y, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}
	return newImg
}

func FlipVertically(img image.Image) image.Image {
	rows, cols := img.Bounds().Dy(), img.Bounds().Dx()
	newImg := image.NewRGBA64(img.Bounds())
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			newImg.SetRGBA64(x, rows-(y+1), color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
		}
	}
	return newImg
}
