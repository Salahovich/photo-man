package image_filters

import (
	"fmt"
	"image"
	"photo-man/core/kernel_utils"
)

type SOBEL_DIRECTION int

const (
	TOP_SOBEL    SOBEL_DIRECTION = 1
	BOTTOM_SOBEL SOBEL_DIRECTION = 2
	LEFT_SOBEL   SOBEL_DIRECTION = 3
	RIGHT_SOBEL  SOBEL_DIRECTION = 4
)

var topSobelKernel = [][]float32{
	{1.0, 2.0, 1.0},
	{0.0, 0.0, 0.0},
	{-1.0, -2.0, -1.0}}
var bottomSobelKernel = [][]float32{
	{-1.0, -2.0, -1.0},
	{0.0, 0.0, 0.0},
	{1.0, 2.0, 1.0}}
var leftSobelKernel = [][]float32{
	{1.0, 0.0, -1.0},
	{2.0, 0.0, -2.0},
	{1.0, 0.0, -1.0}}
var rightSobelKernel = [][]float32{
	{-1.0, 0.0, 1.0},
	{-2.0, 0.0, 2.0},
	{-1.0, 0.0, 1.0}}

func SobelImage(old image.Image, direction SOBEL_DIRECTION) image.Image {
	fmt.Println(old.Bounds())
	switch direction {
	case TOP_SOBEL:
		return kernel_utils.ApplyKernel(old, topSobelKernel)
	case BOTTOM_SOBEL:
		return kernel_utils.ApplyKernel(old, bottomSobelKernel)
	case LEFT_SOBEL:
		return kernel_utils.ApplyKernel(old, leftSobelKernel)
	case RIGHT_SOBEL:
		return kernel_utils.ApplyKernel(old, rightSobelKernel)
	default:
		return old
	}
}
