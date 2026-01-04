package image_io

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/kbinani/screenshot"
	"golang.org/x/image/draw"
)

func ReadImage(filePath string) (image.Image, string, error) {
	file, fileOpenErr := os.Open(filePath)
	if fileOpenErr != nil {
		return nil, "", fileOpenErr
	}

	imageData, format, decodeErr := image.Decode(file)
	if decodeErr != nil {
		return nil, "", decodeErr
	}

	if fileCloseErr := file.Close(); fileCloseErr != nil {
		return nil, "", fileCloseErr
	}

	return ToRGBA(imageData), format, nil
}

func WriteImage(imageData image.Image, path, format string) error {
	targetFile, createImageErr := os.Create(path + format)
	if createImageErr != nil {
		return createImageErr
	}

	if format == "jpeg" || format == "jpg" {
		err := jpeg.Encode(targetFile, imageData, nil)
		if err != nil {
			return err
		}
	} else if format == "png" {
		err := png.Encode(targetFile, imageData)
		if err != nil {
			return err
		}
	} else if format == "gif" {
		err := gif.Encode(targetFile, imageData, nil)
		if err != nil {
			return err
		}
	}

	fileCloseErr := targetFile.Close()
	if fileCloseErr != nil {
		return fileCloseErr
	}

	return nil
}

func ToRGBA(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA64(bounds)

	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}

func GetScreenSize() (int, int) {
	bounds := screenshot.GetDisplayBounds(0)
	return bounds.Dx(), bounds.Dy()
}

func Resize(src image.Image, newWidth, newHeight int) image.Image {
	dst := image.NewRGBA64(image.Rect(0, 0, newWidth, newHeight))
	scaler := draw.CatmullRom
	scaler.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	return dst
}

func Rescale(imageData image.Image) (bool, image.Image) {
	screenWidth, screenHeight := GetScreenSize()
	imgWidth, imgHeight := float64(imageData.Bounds().Dx()), float64(imageData.Bounds().Dy())
	newImgWidth, newImgHeight := imgWidth, imgHeight
	ratio := math.Min(imgWidth/imgHeight, imgHeight/imgWidth)
	doScale := false

	// rescale to get the suitable size for the screen
	for newImgWidth >= float64(screenWidth) || newImgHeight >= float64(screenHeight) {
		doScale = true
		newImgWidth, newImgHeight = ratio*newImgWidth, ratio*newImgHeight
	}

	if doScale {
		return true, Resize(imageData, int(newImgWidth), int(newImgHeight))
	}
	return false, nil
}
