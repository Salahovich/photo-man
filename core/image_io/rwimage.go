package image_io

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

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

func WriteImage(imageData image.Image, path, filterName, format string) error {
	targetDir := filepath.Dir(path)
	newName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + "-" + filterName + filepath.Ext(path)
	newPath := filepath.Join(targetDir, newName)

	targetFile, createImageErr := os.Create(newPath)
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
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}

func GetScreenSize() (int, int) {
	bounds := screenshot.GetDisplayBounds(0)
	return bounds.Dx(), bounds.Dy()
}

func Resize(src image.Image, newWidth, newHeight int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	scaler := draw.CatmullRom
	scaler.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	return dst
}

func Rescale(imageData image.Image) (bool, image.Image) {
	screenWidth, screenHeight := GetScreenSize()
	imgWidth, imgHeight := imageData.Bounds().Dx(), imageData.Bounds().Dy()

	if imgWidth >= screenWidth || imgHeight >= screenHeight {
		newImgWidth := screenWidth / 2
		k := float32(newImgWidth) / float32(imgWidth)
		newImgHeight := int(float32(imgHeight) * k)
		return true, Resize(imageData, newImgWidth, newImgHeight)
	}
	return false, nil
}
