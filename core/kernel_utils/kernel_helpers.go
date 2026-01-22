package kernel_utils

import "image"

func spotKernelWindow(
	x, y int,
	dx, dy int,
	kw, kh int,
	img image.Image,
	getColor func(x, y int, img image.Image) uint16,
	edgeHandler func(window [][]int32) ([][]int32, int)) ([][]int32, int) {

	// init result window
	result := make([][]int32, kw)
	for i := 0; i < kw; i++ {
		result[i] = make([]int32, kh)
		for j := 0; j < kh; j++ {
			result[i][j] = -1
		}
	}
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	stX, stY, enX, enY := x-dx, y-dy, x+dx, y+dy

	// check for edges overlap
	var overlap bool
	if stX < 0 || stX >= imgWidth ||
		stY < 0 || stY >= imgHeight ||
		enX < 0 || enX >= imgWidth ||
		enY < 0 || enY >= imgHeight {
		overlap = true
	}

	// start and end values for the result matrix and image
	resStX, resEnX, resStY, resEnY := 0, kw, 0, kh
	if stX < 0 {
		resStX = -stX
		stX = 0
	}
	if stY < 0 {
		resStY = -stY
		stY = 0
	}
	if enX >= imgWidth {
		resEnX = kw - (enX - x)
		enX = imgWidth - 1
	}
	if enY >= imgHeight {
		resEnY = kh - (enY - y)
		enY = imgHeight - 1
	}

	// init the inside pixels of the image into window
	for i, k := resStX, stY; i < resEnX && k <= enY; i, k = i+1, k+1 {
		for j, m := resStY, stX; j < resEnY && m <= enX; j, m = j+1, m+1 {
			result[i][j] = int32(getColor(m, k, img))
		}
	}

	if overlap {
		return edgeHandler(result)
	}
	return result, 0
}

func ignoreOverlappedEdges(window [][]int32) ([][]int32, int) {
	var count int
	for i := 0; i < len(window); i++ {
		for j := 0; j < len(window[i]); j++ {
			if window[i][j] == -1 {
				window[i][j] = 0
				count++
			}
		}
	}
	return window, count
}

func multiply(window [][]int32, kernel [][]float32) [][]float32 {
	rw, cl := len(window), len(window[0])
	result := make([][]float32, rw)
	for i := 0; i < rw; i++ {
		result[i] = make([]float32, cl)
		for j := 0; j < cl; j++ {
			result[i][j] = float32(window[i][j]) * kernel[i][j]
		}
	}
	return result
}

func normalizeWindow(window [][]float32) uint16 {
	var pixel float32
	for i := 0; i < len(window); i++ {
		for j := 0; j < len(window[i]); j++ {
			pixel += window[i][j]
		}
	}
	if pixel > 65535 {
		pixel = 65535
	} else if pixel < 0 {
		pixel = 0
	}
	return uint16(pixel)
}

func CalculateNormalizedValue(x, y, dx, dy, kw, kh int, oldImg image.Image, kernel [][]float32, getColor func(x, y int, img image.Image) uint16) uint16 {
	window, _ := spotKernelWindow(x, y, dx, dy, kw, kh, oldImg, getColor, ignoreOverlappedEdges)
	mulWindow := multiply(window, kernel)
	return normalizeWindow(mulWindow)
}

// func clamp_overlaped_edges(window [][]int32) [][]int32 {
//
// }

//func mirror_overlaped_edges(window [][]int32) [][]int32 {
//
//}
//
//func wrap_overlaped_edges(window [][]int32) [][]int32 {
//
//}
