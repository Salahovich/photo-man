package state

import (
	"image"
	"sync"
)

type CanvasState struct {
	currentImage  image.Image
	scaledImage   image.Image
	originalImage image.Image
	modifications []func(image.Image) image.Image
	communication chan image.Image
	canvasMutex   sync.RWMutex
	format        string
}

func (c *CanvasState) UpdateSceneImage(img image.Image) {
	c.canvasMutex.Lock()
	defer c.canvasMutex.Unlock()
	c.currentImage = img
	c.communication <- c.currentImage
}

func (c *CanvasState) SetImage(original image.Image, scaled image.Image) {
	c.canvasMutex.Lock()
	defer c.canvasMutex.Unlock()
	c.originalImage = original
	c.scaledImage = scaled
	c.currentImage = scaled
	c.communication <- c.currentImage
}

func (c *CanvasState) GetChannel() chan image.Image {
	return c.communication
}

func (c *CanvasState) RegisterModification(callback func(image.Image) image.Image) {
	c.modifications = append(c.modifications, callback)
}
func (c *CanvasState) ResetModifications() {
	c.modifications = make([]func(image2 image.Image) image.Image, 0)
}

func (c *CanvasState) SetFormat(format string) {
	c.format = format
}

func (c *CanvasState) GetScaledImage() image.Image {
	c.canvasMutex.RLock()
	defer c.canvasMutex.RUnlock()
	return c.scaledImage
}

func (c *CanvasState) GetOriginalImage() image.Image {
	c.canvasMutex.RLock()
	defer c.canvasMutex.RUnlock()
	return c.originalImage
}

func (c *CanvasState) GetCurrentImage() image.Image {
	c.canvasMutex.RLock()
	defer c.canvasMutex.RUnlock()
	return c.currentImage
}

func (c *CanvasState) GetFormat() string {
	return c.format
}

func (c *CanvasState) ApplyAllModification() {
	c.canvasMutex.Lock()
	defer c.canvasMutex.Unlock()

	for _, callback := range c.modifications {
		c.originalImage = callback(c.originalImage)
	}
}
