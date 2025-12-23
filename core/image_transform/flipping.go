package image_transform

import (
	"image"
)

func FlipHorizontally(img image.Image) image.Image {
	rows, cols := img.Bounds().Dy(), img.Bounds().Dx()
	newImg := image.NewRGBA(img.Bounds())
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			newImg.Set(cols-(x+1), y, img.At(x, y))
		}
	}
	return newImg
}

func FlipVertically(img image.Image) image.Image {
	rows, cols := img.Bounds().Dy(), img.Bounds().Dx()
	newImg := image.NewRGBA(img.Bounds())
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			newImg.Set(x, rows-(y+1), img.At(x, y))
		}
	}
	return newImg
}
