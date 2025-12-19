package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"learn-repo/filters"
	"learn-repo/gray_filters"
	imageio "learn-repo/image_io"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	output := canvas.NewText(time.Now().Format(time.TimeOnly), color.NRGBA{G: 0xff, A: 0xff})
	output.TextStyle.Monospace = true
	output.TextSize = 32
	w.SetContent(output)

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fyne.Do(func() {
				output.Text = time.Now().Format(time.TimeOnly)
				output.Refresh()
			})
		}
	}()
	w.ShowAndRun()
}

func writeGrayImg(imageData image.Image, format, path string) *image.Gray {
	// filter image
	newImage := gray_filters.BasicGreyScale(imageData)
	err := imageio.WriteImage(newImage, path, "grey", format)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
	return newImage
}
func writeBlurGrayImg(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters.SimpleBlur(blurGrayImg, filters.LOW_BLUR)
	err := imageio.WriteImage(newImage, path, "gray-blur", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("BLUR GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSharpGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters.SharpImage(blurGrayImg, filters.MEDIUM_SHARP)
	err := imageio.WriteImage(newImage, path, "sharp-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeEmbossGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters.EmbossImage(blurGrayImg, filters.STANDARD_EMBOSS)
	err := imageio.WriteImage(newImage, path, "emboss-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSobelGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters.SobelImage(blurGrayImg, filters.RIGHT_SOBEL)
	err := imageio.WriteImage(newImage, path, "sobel-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeOutlineGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters.OutlineImage(blurGrayImg, filters.STANDARD_OUTLINE)
	err := imageio.WriteImage(newImage, path, "outline-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeBlurRGBAImg(rgbImg image.Image, format, path string) {
	// write image
	newImage := filters.SimpleBlur(rgbImg, filters.HIGH_BLUR)
	err := imageio.WriteImage(newImage, path, "rgb-blur", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("BLUR RGB IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSharpRGBAImage(rgbImg image.Image, format, path string) {
	// write image
	newImage := filters.SharpImage(rgbImg, filters.LOW_SHARP)
	err := imageio.WriteImage(newImage, path, "sharp-rgb", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func ToRGBA(img image.Image) *image.RGBA {
	// 1. If it's already RGBA, just return it
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}

	// 2. Create a new empty RGBA image with the same bounds
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// 3. Draw the source image onto the new RGBA image
	// This handles all the conversion math (YCbCr -> RGBA, etc.) for you
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}

func startProfiling() {
	f1, _ := os.Create("start.prof")
	err1 := pprof.WriteHeapProfile(f1)
	if err1 != nil {
		return
	}
	err2 := f1.Close()
	if err2 != nil {
		return
	}

}

func endProfiling() {
	runtime.GC()

	// 3. Write the heap profile to the file
	f2, _ := os.Create("end.prof")
	err3 := pprof.WriteHeapProfile(f2)
	if err3 != nil {
		return
	}
	err4 := f2.Close()
	if err4 != nil {
		return
	}
}
