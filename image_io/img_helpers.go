package image_io

import (
	"image"
)

type COLOR_TYPE uint8

const (
	GRAY    COLOR_TYPE = 0
	RGBA    COLOR_TYPE = 1
	UNKNOWN COLOR_TYPE = 3
)

func ImageColorType(img image.Image) COLOR_TYPE {
	switch img.(type) {
	case *image.Gray:
		return GRAY
	case *image.RGBA:
		return RGBA
	default:
		return UNKNOWN
	}
}

func InitImageInstance(img image.Image) (image.Image, COLOR_TYPE) {
	switch ImageColorType(img) {
	case GRAY:
		return image.NewGray(img.Bounds()), GRAY
	case RGBA:
		return image.NewRGBA(img.Bounds()), RGBA
	default:
		return nil, UNKNOWN
	}
}

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
