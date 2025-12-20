package image_io

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
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

	fileCloseErr := file.Close()
	if fileCloseErr != nil {
		return nil, "", fileCloseErr
	}

	return imageData, format, nil
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
