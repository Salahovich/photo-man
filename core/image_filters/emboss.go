package image_filters

import (
	"image"
	"photo-man/core/kernel_utils"
)

type EMBOSS_QUALITY int

const (
	LIGHT_EMBOSS EMBOSS_QUALITY = 1
	DARK_EMBOSS  EMBOSS_QUALITY = 2
	HEAVY_EMBOSS EMBOSS_QUALITY = 3
)

var lightEmbossKernel = [][]float32{
	{1.0, 0.0, 0.0},
	{0.0, 1.0, 0.0},
	{0.0, 0.0, -1.0}}
var heavyEmbossKernel = [][]float32{
	{-2.0, -1.0, 0.0},
	{-1.0, 1.0, 1.0},
	{0.0, 1.0, 2.0}}
var darkEmbossKernel = [][]float32{
	{-1.0, -0.5, 0.0},
	{-0.5, 0.5, 0.5},
	{0.0, 0.5, 1.0}}

func EmbossImage(old image.Image, quality EMBOSS_QUALITY) image.Image {
	switch quality {
	case LIGHT_EMBOSS:
		return kernel_utils.ApplyKernel(old, lightEmbossKernel)
	case DARK_EMBOSS:
		return kernel_utils.ApplyKernel(old, darkEmbossKernel)
	case HEAVY_EMBOSS:
		return kernel_utils.ApplyKernel(old, heavyEmbossKernel)
	default:
		return old
	}
}
