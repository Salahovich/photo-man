package image_io

import (
	"image"
)

func GetRedValue(x, y int, img image.Image) uint8 {
	r, _, _, _ := img.At(x, y).RGBA()
	return uint8(r >> 8)
}
func GetGreenValue(x, y int, img image.Image) uint8 {
	_, g, _, _ := img.At(x, y).RGBA()
	return uint8(g >> 8)
}
func GetBlueValue(x, y int, img image.Image) uint8 {
	_, _, b, _ := img.At(x, y).RGBA()
	return uint8(b >> 8)
}
func GetAlphaValue(x, y int, img image.Image) uint8 {
	_, _, _, a := img.At(x, y).RGBA()
	return uint8(a >> 8)
}
func GetGrayValue(x, y int, img image.Image) uint8 {
	concrete, _ := img.(*image.Gray)
	return concrete.GrayAt(x, y).Y
}
