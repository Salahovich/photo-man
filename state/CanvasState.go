package state

import (
	"image"
	"sync"
)

type CanvasState struct {
	imageInCanvas bool
	currentImage  image.Image
	scaledImage   image.Image
	originalImage image.Image
	communication chan image.Image
	canvasMutex   *sync.RWMutex
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

func (c *CanvasState) GetCanvasMutex() *sync.RWMutex {
	return c.canvasMutex
}
func (c *CanvasState) GetFormat() string {
	return c.format
}

func (c *CanvasState) IsImageInCanvas() bool {
	return c.imageInCanvas
}

func (c *CanvasState) SetImageInCanvs(isIn bool) {
	c.imageInCanvas = isIn
}
