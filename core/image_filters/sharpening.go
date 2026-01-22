package image_filters

import (
	"image"
	"image/color"
	"photo-man/core/kernel_utils"
)

type SHARPENING_QUALITY int

const (
	LOW_SHARP    SHARPENING_QUALITY = 1
	MEDIUM_SHARP SHARPENING_QUALITY = 2
	HIGH_SHARP   SHARPENING_QUALITY = 3
)

var lowSharpeningKernel = [][]float32{
	{0.0, -1.0, 0.0},
	{-1.0, 5.0, -1.0},
	{0.0, -1.0, 0.0}}
var mediumSharpeningKernel = [][]float32{
	{0.0, -1.0, 0.0},
	{-1.0, 5.5, -1.0},
	{0.0, -1.0, 0.0}}
var highSharpeningKernel = [][]float32{
	{0.0, -1.0, 0.0},
	{-1.0, 6.0, -1.0},
	{0.0, -1.0, 0.0}}

func SharpenForBrush(x, y int, old image.Image, quality SHARPENING_QUALITY) color.Color {
	switch quality {
	case LOW_SHARP:
		return kernel_utils.ApplyKernelForBrush(x, y, old, lowSharpeningKernel)
	case MEDIUM_SHARP:
		return kernel_utils.ApplyKernelForBrush(x, y, old, mediumSharpeningKernel)
	case HIGH_SHARP:
		return kernel_utils.ApplyKernelForBrush(x, y, old, highSharpeningKernel)
	default:
		return color.RGBA64{}
	}
}

func SharpImage(old image.Image, quality SHARPENING_QUALITY) image.Image {
	switch quality {
	case LOW_SHARP:
		return kernel_utils.ApplyKernel(old, lowSharpeningKernel)
	case MEDIUM_SHARP:
		return kernel_utils.ApplyKernel(old, mediumSharpeningKernel)
	case HIGH_SHARP:
		return kernel_utils.ApplyKernel(old, highSharpeningKernel)
	default:
		return old
	}
}
