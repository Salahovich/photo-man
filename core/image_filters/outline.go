package image_filters

import (
	"image"
	"photo-man/core/kernel_utils"
)

type OUTLINE_QUALITY int

const (
	STANDARD_OUTLINE OUTLINE_QUALITY = 1
)

var standardOutlineKernel = [][]float32{
	{-1.0, -1.0, -1.0},
	{-1.0, 9, -1.0},
	{-1.0, -1.0, -1.0}}

func OutlineImage(old image.Image, quality OUTLINE_QUALITY) image.Image {
	switch quality {
	case STANDARD_OUTLINE:
		return kernel_utils.ApplyKernel(old, standardOutlineKernel)
	default:
		return old
	}
}
