package filters

import (
	"image"
	"learn-repo/kernel_utils"
)

type EMBOSS_QUALITY int

const (
	STANDARD_EMBOSS EMBOSS_QUALITY = 1
)

var standardEmbossKernel = [][]float32{
	{-2.0, -1.0, 0.0},
	{-1.0, 1.0, 1.0},
	{0.0, 1.0, 2.0}}

func EmbossImage(old image.Image, quality EMBOSS_QUALITY) image.Image {
	switch quality {
	case STANDARD_EMBOSS:
		return kernel_utils.ApplyKernel(old, standardEmbossKernel)
	default:
		return kernel_utils.ApplyKernel(old, standardEmbossKernel)
	}
}
