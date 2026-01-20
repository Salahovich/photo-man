package image_io

import (
	"image"
	"image/color"
)

type SystemColor struct {
	Color color.Color
}

func GetRedValue(x, y int, img image.Image) uint16 {
	r, _, _, _ := img.At(x, y).RGBA()
	return uint16(r)
}
func GetGreenValue(x, y int, img image.Image) uint16 {
	_, g, _, _ := img.At(x, y).RGBA()
	return uint16(g)
}
func GetBlueValue(x, y int, img image.Image) uint16 {
	_, _, b, _ := img.At(x, y).RGBA()
	return uint16(b)
}
func GetAlphaValue(x, y int, img image.Image) uint16 {
	_, _, _, a := img.At(x, y).RGBA()
	return uint16(a)
}
func GetGrayValue(x, y int, img image.Image) uint16 {
	concrete, _ := img.(*image.Gray)
	r, _, _, _ := concrete.GrayAt(x, y).RGBA()
	return uint16(r)
}
