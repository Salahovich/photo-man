package state

import (
	"image"
	customUI "photo-man/ui/custom-ui"
	"sync"

	"fyne.io/fyne/v2"
)

type CanvasState struct {
	imageInCanvas     bool
	currentImage      image.Image
	scaledImage       image.Image
	originalImage     image.Image
	canvasStack       *fyne.Container
	cropState         *CropState
	paintBoardState   *PaintBoardState
	blurBoradState    *BlurBoardState
	sharpenBoradState *SharpenBoardState
	communication     chan image.Image
	canvasMutex       *sync.RWMutex
	format            string
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

func (c *CanvasState) SetScaledImage(img image.Image) {
	c.scaledImage = img
}

func (c *CanvasState) SetCanvasStack(canvasStack *fyne.Container) {
	c.canvasStack = canvasStack
}

func (c *CanvasState) GetCanvasStack() *fyne.Container {
	return c.canvasStack
}

func (c *CanvasState) GetCropState() *CropState {
	return c.cropState
}

func (c *CanvasState) GetPaintBoardState() *PaintBoardState {
	return c.paintBoardState
}
func (c *CanvasState) GetBlurBoardState() *BlurBoardState {
	return c.blurBoradState
}
func (c *CanvasState) GetSharpenBoardState() *SharpenBoardState {
	return c.sharpenBoradState
}

func (c *CanvasState) AddPaintBoardLayer(layer *customUI.PaintBoard) {
	c.paintBoardState.paintBoardCanvas = layer
	c.canvasStack.Add(layer)
	c.paintBoardState.EnablePaintBoard()
}

func (c *CanvasState) RemovePaintBoardLayer() {
	c.canvasStack.Remove(c.paintBoardState.paintBoardCanvas)
	c.paintBoardState.DisablePaintBoard()
}

func (c *CanvasState) AddBlurBoardLayer(layer *customUI.BlurBoard) {
	c.blurBoradState.blurBoardCanvas = layer
	c.blurBoradState.EnableBlurBoard()
	c.canvasStack.Add(layer)
}

func (c *CanvasState) RemoveBlurBoardLayer() {
	c.canvasStack.Remove(c.blurBoradState.blurBoardCanvas)
	c.blurBoradState.DisableBlurBoard()
}

func (c *CanvasState) AddSharpenBoardLayer(layer *customUI.SharpenBoard) {
	c.sharpenBoradState.sharpenBoardCanvas = layer
	c.sharpenBoradState.EnableSharpenBoard()
	c.canvasStack.Add(layer)
}

func (c *CanvasState) RemoveSharpenBoardLayer() {
	c.canvasStack.Remove(c.sharpenBoradState.GetSharpenBoardCanvas())
	c.blurBoradState.DisableBlurBoard()
}

func (c *CanvasState) AddCropLayer(layer *customUI.ResizableRectangle) {
	c.cropState.cropImageCanvas = layer
	c.cropState.EnableCropState()
	c.canvasStack.Add(layer)
}

func (c *CanvasState) RemoveCropLayer() {
	c.canvasStack.Remove(c.cropState.cropImageCanvas)
	c.cropState.DisableCropState()
}
