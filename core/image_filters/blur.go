package image_filters

import (
	"image"
	"photo-man/core/kernel_utils"
)

var low_blur_kernel = [][]float32{
	{1.0 / 9, 1.0 / 9, 1.0 / 9},
	{1.0 / 9, 1.0 / 9, 1.0 / 9},
	{1.0 / 9, 1.0 / 9, 1.0 / 9}}
var medium_blur_kernel = [][]float32{
	{1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25},
	{1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25},
	{1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25},
	{1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25},
	{1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25, 1.0 / 25}}
var high_blur_kernel = [][]float32{
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49},
	{1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49, 1.0 / 49}}

var low_gaussian_blur_kernel = [][]float32{
	{1.0 / 16, 2.0 / 16, 1.0 / 16},
	{2.0 / 16, 4.0 / 16, 2.0 / 16},
	{1.0 / 16, 2.0 / 16, 1.0 / 16}}
var medium_gaussian_blur_kernel = [][]float32{
	{1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256},
	{4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256},
	{6.0 / 256, 24.0 / 256, 36.0 / 256, 24.0 / 256, 6.0 / 256},
	{4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256},
	{1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256}}
var high_gaussian_blur_kernel = [][]float32{
	{1.0 / 4096, 6.0 / 4096, 15.0 / 4096, 20.0 / 4096, 15.0 / 4096, 6.0 / 4096, 1.0 / 4096},
	{6.0 / 4096, 36.0 / 4096, 90.0 / 4096, 120.0 / 4096, 90.0 / 4096, 36.0 / 4096, 6.0 / 4096},
	{15.0 / 4096, 90.0 / 4096, 225.0 / 4096, 300.0 / 4096, 225.0 / 4096, 90.0 / 4096, 15.0 / 4096},
	{20.0 / 4096, 120.0 / 4096, 300.0 / 4096, 400.0 / 4096, 300.0 / 4096, 120.0 / 4096, 20.0 / 4096},
	{15.0 / 4096, 90.0 / 4096, 225.0 / 4096, 300.0 / 4096, 225.0 / 4096, 90.0 / 4096, 15.0 / 4096},
	{6.0 / 4096, 36.0 / 4096, 90.0 / 4096, 120.0 / 4096, 90.0 / 4096, 36.0 / 4096, 6.0 / 4096},
	{1.0 / 4096, 6.0 / 4096, 15.0 / 4096, 20.0 / 4096, 15.0 / 4096, 6.0 / 4096, 1.0 / 4096}}

type BLUR_QUALITY int

const (
	LOW_BLUR    BLUR_QUALITY = 1
	MEDIUM_BLUR BLUR_QUALITY = 2
	HIGH_BLUR   BLUR_QUALITY = 3
)

func SimpleBlur(old image.Image, quality BLUR_QUALITY) image.Image {
	switch quality {
	case LOW_BLUR:
		return kernel_utils.ApplyKernel(old, low_blur_kernel)
	case MEDIUM_BLUR:
		return kernel_utils.ApplyKernel(old, medium_blur_kernel)
	case HIGH_BLUR:
		return kernel_utils.ApplyKernel(old, high_blur_kernel)
	default:
		return kernel_utils.ApplyKernel(old, low_blur_kernel)
	}
}

func GaussianBlur(old image.Image, quality BLUR_QUALITY) image.Image {
	switch quality {
	case LOW_BLUR:
		return kernel_utils.ApplyKernel(old, low_gaussian_blur_kernel)
	case MEDIUM_BLUR:
		return kernel_utils.ApplyKernel(old, medium_gaussian_blur_kernel)
	case HIGH_BLUR:
		return kernel_utils.ApplyKernel(old, high_gaussian_blur_kernel)
	default:
		return kernel_utils.ApplyKernel(old, low_gaussian_blur_kernel)
	}
}
