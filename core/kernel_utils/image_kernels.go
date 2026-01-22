package kernel_utils

import (
	"image"
	"image/color"
	imgHelpers "photo-man/core/image_io"
	"sync"
)

func ApplyKernel(srcImg image.Image, kernel [][]float32) image.Image {

	width := srcImg.Bounds().Dx()
	height := srcImg.Bounds().Dy()
	dx := len(kernel[0]) / 2
	dy := len(kernel) / 2
	kw := len(kernel)
	kh := len(kernel[0])
	newImg := image.NewRGBA64(srcImg.Bounds())

	return applyRGBAImageKernel(width, height, dx, dy, kw, kh, srcImg, newImg, kernel)

}

func ApplyKernelForBrush(x, y int, srcImg image.Image, kernel [][]float32) color.Color {

	dx, dy := len(kernel[0])/2, len(kernel)/2
	kw, kh := len(kernel), len(kernel[0])

	return applyRGBAImageKernelForBrush(x, y, dx, dy, kw, kh, srcImg, kernel)

}

// func applyGrayImageKernel(width, height, dx, dy, kw, kh int, oldImg image.Image, newImg *image.Gray, kernel [][]float32) *image.Gray {
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			normalizedValue := CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetGrayValue)
// 			newImg.Set(x, y, color.Gray{Y: normalizedValue})
// 		}
// 	}
// 	return newImg
// }

func applyRGBAImageKernel(width, height, dx, dy, kw, kh int, oldImg image.Image, newImg *image.RGBA64, kernel [][]float32) *image.RGBA64 {
	wg := sync.WaitGroup{}
	wg.Add(height)
	for y := range height {
		go func(y int) {
			defer wg.Done()
			for x := range width {
				var red, green, blue, alpha uint16
				red = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetRedValue)
				green = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetGreenValue)
				blue = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetBlueValue)
				alpha = imgHelpers.GetAlphaValue(x, y, oldImg)

				newImg.SetRGBA64(x, y, color.RGBA64{R: red, G: green, B: blue, A: alpha})
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func applyRGBAImageKernelForBrush(x, y int, dx, dy, kw, kh int, oldImg image.Image, kernel [][]float32) color.Color {

	var red, green, blue, alpha uint16
	red = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetRedValue)
	green = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetGreenValue)
	blue = CalculateNormalizedValue(x, y, dx, dy, kw, kh, oldImg, kernel, imgHelpers.GetBlueValue)
	alpha = imgHelpers.GetAlphaValue(x, y, oldImg)

	return color.RGBA64{R: red, G: green, B: blue, A: alpha}
}
