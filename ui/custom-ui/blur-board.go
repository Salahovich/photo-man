package customUI

import (
	"image"
	"image/color"
	"math"

	"photo-man/core/image_filters"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// generate compile time error if the interface methods did not implemented.
var _ desktop.Hoverable = (*PaintBoard)(nil)
var _ fyne.Draggable = (*PaintBoard)(nil)
var _ desktop.Cursorable = (*PaintBoard)(nil)
var _ fyne.Tappable = (*PaintBoard)(nil)

type BlurBoard struct {
	widget.BaseWidget
	bluredImage   *image.RGBA64
	originalImage image.Image
	blurCanvas    *canvas.Image
	blurHardness  image_filters.BLUR_QUALITY
	blurType      image_filters.BLUR_TYPE
	inEraserBrush bool
	brushSize     int
	blured        bool
}

func NewBlurBoard(size fyne.Size, original image.Image) *BlurBoard {

	// init the board
	var blurBoard *BlurBoard

	// init the main widget
	blurBoard = &BlurBoard{
		bluredImage:   image.NewRGBA64(image.Rect(0, 0, int(size.Width), int(size.Height))),
		originalImage: original,
		brushSize:     3.0,
		blurHardness:  image_filters.MEDIUM_BLUR,
		blurType:      image_filters.SHAPE_BLUR,
		blured:        false,
		inEraserBrush: false,
	}
	blurBoard.blurCanvas = canvas.NewImageFromImage(blurBoard.bluredImage)

	blurBoard.ExtendBaseWidget(blurBoard)

	return blurBoard
}

func (bb *BlurBoard) DragEnd() {}
func (bb *BlurBoard) Dragged(ev *fyne.DragEvent) {
	if bb.brushSize == 0 {
		return
	}

	pixel := ev.Position
	radius := float64(bb.brushSize)

	// position a square of the brush size as its side!
	startPos := fyne.Position{X: float32(pixel.X - float32(bb.brushSize)), Y: float32(pixel.Y - float32(bb.brushSize))}
	endPos := fyne.Position{X: float32(pixel.X + float32(bb.brushSize)), Y: float32(pixel.Y + float32(bb.brushSize))}

	for y := int(startPos.Y); y < int(endPos.Y); y++ {
		for x := int(startPos.X); x < int(endPos.X); x++ {
			if x < 0 || x >= bb.bluredImage.Bounds().Dx() || y < 0 || y >= bb.bluredImage.Bounds().Dy() {
				continue
			}
			if bb.inEraserBrush {
				bb.bluredImage.SetRGBA64(x, y, color.RGBA64{})
			} else {
				distance := math.Round(math.Hypot(float64(pixel.X)-float64(x), float64(pixel.Y)-float64(y)))
				if distance < radius {
					r, g, b, a := image_filters.BlurForBrush(x, y, bb.originalImage, bb.blurType, bb.blurHardness).RGBA()
					bb.bluredImage.SetRGBA64(x, y, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
				}
			}
		}
	}

	defer func() {
		if !bb.blured {
			bb.blured = true
		}
		bb.Refresh()
	}()
}

func (bb *BlurBoard) MouseMoved(ev *desktop.MouseEvent) {

}
func (bb *BlurBoard) MouseIn(ev *desktop.MouseEvent) {

}
func (bb *BlurBoard) MouseOut() {

}

func (bb *BlurBoard) Tapped(ev *fyne.PointEvent) {
}

func (bb *BlurBoard) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (bb *BlurBoard) MinSize() fyne.Size {
	return fyne.NewSize(float32(bb.bluredImage.Bounds().Dx()), float32(bb.bluredImage.Bounds().Dy()))
}

func (bb *BlurBoard) SetBlurType(blurType image_filters.BLUR_TYPE) {
	bb.blurType = blurType
}

func (bb *BlurBoard) GetBlurType() image_filters.BLUR_TYPE {
	return bb.blurType
}

func (bb *BlurBoard) SetBlurHardness(blurHardness image_filters.BLUR_QUALITY) {
	bb.blurHardness = blurHardness
}

func (bb *BlurBoard) GetBlurHardness() image_filters.BLUR_QUALITY {
	return bb.blurHardness
}

func (bb *BlurBoard) SetBrushSize(size int) {
	bb.brushSize = size
}

func (bb *BlurBoard) GetBrushSize() int {
	return bb.brushSize
}

func (bb *BlurBoard) CanBlur() bool {
	return bb.blured
}

func (bb *BlurBoard) ToggleEraserBrush() {
	bb.inEraserBrush = !bb.inEraserBrush
}

func (bb *BlurBoard) IsEraserBrush() bool {
	return bb.inEraserBrush
}

func (bb *BlurBoard) GetBoard() image.Image {
	return bb.bluredImage
}

func (bb *BlurBoard) ClearBoard() {
	bb.bluredImage = image.NewRGBA64(image.Rect(0, 0, bb.bluredImage.Bounds().Dx(), bb.bluredImage.Bounds().Dy()))
	bb.blurCanvas.Image = bb.bluredImage
	bb.blurCanvas.Refresh()
}

func (bb *BlurBoard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(bb.blurCanvas)
}
