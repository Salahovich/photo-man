package main

import (
	"fmt"
	"image"
	"image/draw"
	filters2 "learn-repo/core/filters"
	"learn-repo/core/gray_filters"
	imageio "learn-repo/core/image_io"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
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
