package image_transform

import (
	"image"
)

func RotateClockwise(img image.Image) image.Image {
	cols, rows := img.Bounds().Dx(), img.Bounds().Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, rows, cols))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			newImg.Set(rows-(y+1), x, img.At(x, y))
		}
	}
	return newImg
}

func RotateAntiClockwise(img image.Image) image.Image {
	cols, rows := img.Bounds().Dx(), img.Bounds().Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, rows, cols))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			newImg.Set(y, x, img.At(x, y))
		}
	}
	return newImg
}
