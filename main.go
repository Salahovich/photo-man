package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	filters2 "photo-man/core/filters"
	"photo-man/core/gray_filters"
	imageio "photo-man/core/image_io"
	"photo-man/ui"
	"runtime"
	"runtime/pprof"
)

func main() {
	ui.StartApp()
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
	newImage := filters2.SimpleBlur(blurGrayImg, filters2.LOW_BLUR)
	err := imageio.WriteImage(newImage, path, "gray-blur", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("BLUR GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSharpGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters2.SharpImage(blurGrayImg, filters2.MEDIUM_SHARP)
	err := imageio.WriteImage(newImage, path, "sharp-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeEmbossGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters2.EmbossImage(blurGrayImg, filters2.STANDARD_EMBOSS)
	err := imageio.WriteImage(newImage, path, "emboss-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSobelGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters2.SobelImage(blurGrayImg, filters2.RIGHT_SOBEL)
	err := imageio.WriteImage(newImage, path, "sobel-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeOutlineGrayImage(blurGrayImg *image.Gray, format, path string) {
	// write image
	newImage := filters2.OutlineImage(blurGrayImg, filters2.STANDARD_OUTLINE)
	err := imageio.WriteImage(newImage, path, "outline-gray", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeBlurRGBAImg(rgbImg image.Image, format, path string) {
	// write image
	newImage := filters2.SimpleBlur(rgbImg, filters2.HIGH_BLUR)
	err := imageio.WriteImage(newImage, path, "rgb-blur", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("BLUR RGB IMAGE has been created at the directory: %v \n", filepath.Dir(path))
}
func writeSharpRGBAImage(rgbImg image.Image, format, path string) {
	// write image
	newImage := filters2.SharpImage(rgbImg, filters2.LOW_SHARP)
	err := imageio.WriteImage(newImage, path, "sharp-rgb", format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SHARP GREY IMAGE has been created at the directory: %v \n", filepath.Dir(path))
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
