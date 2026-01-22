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

type SharpenBoard struct {
	widget.BaseWidget
	sharpenedImage  *image.RGBA64
	originalImage   image.Image
	sharpenCanvas   *canvas.Image
	sharpenHardness image_filters.SHARPENING_QUALITY
	inEraserBrush   bool
	brushSize       int
	sharpened       bool
}

func NewSharpenBoard(size fyne.Size, original image.Image) *SharpenBoard {

	// init the board
	var sharpenBoard *SharpenBoard

	// init the main widget
	sharpenBoard = &SharpenBoard{
		sharpenedImage:  image.NewRGBA64(image.Rect(0, 0, int(size.Width), int(size.Height))),
		originalImage:   original,
		brushSize:       3.0,
		sharpenHardness: image_filters.MEDIUM_SHARP,
		sharpened:       false,
		inEraserBrush:   false,
	}
	sharpenBoard.sharpenCanvas = canvas.NewImageFromImage(sharpenBoard.sharpenedImage)

	sharpenBoard.ExtendBaseWidget(sharpenBoard)

	return sharpenBoard
}

func (sb *SharpenBoard) DragEnd() {}
func (sb *SharpenBoard) Dragged(ev *fyne.DragEvent) {
	if sb.brushSize == 0 {
		return
	}

	pixel := ev.Position
	radius := float64(sb.brushSize)

	// position a square of the brush size as its side!
	startPos := fyne.Position{X: float32(pixel.X - float32(sb.brushSize)), Y: float32(pixel.Y - float32(sb.brushSize))}
	endPos := fyne.Position{X: float32(pixel.X + float32(sb.brushSize)), Y: float32(pixel.Y + float32(sb.brushSize))}

	for y := int(startPos.Y); y < int(endPos.Y); y++ {
		for x := int(startPos.X); x < int(endPos.X); x++ {
			if x < 0 || x >= sb.sharpenedImage.Bounds().Dx() || y < 0 || y >= sb.sharpenedImage.Bounds().Dy() {
				continue
			}
			if sb.inEraserBrush {
				// erase blured pixels
				sb.sharpenedImage.SetRGBA64(x, y, color.RGBA64{})
			} else {
				// blur the target pixels inside the circle
				distance := math.Round(math.Hypot(float64(pixel.X)-float64(x), float64(pixel.Y)-float64(y)))
				if distance < radius {
					r, g, b, a := image_filters.SharpenForBrush(x, y, sb.originalImage, sb.sharpenHardness).RGBA()
					sb.sharpenedImage.SetRGBA64(x, y, color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
				}
			}
		}
	}

	defer func() {
		if !sb.sharpened {
			sb.sharpened = true
		}
		sb.Refresh()
	}()
}

func (sb *SharpenBoard) MouseMoved(ev *desktop.MouseEvent) {

}
func (sb *SharpenBoard) MouseIn(ev *desktop.MouseEvent) {

}
func (sb *SharpenBoard) MouseOut() {

}

func (sb *SharpenBoard) Tapped(ev *fyne.PointEvent) {
}

func (sb *SharpenBoard) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (sb *SharpenBoard) MinSize() fyne.Size {
	return fyne.NewSize(float32(sb.sharpenedImage.Bounds().Dx()), float32(sb.sharpenedImage.Bounds().Dy()))
}

func (sb *SharpenBoard) SetSharpenesHardness(sharpenHardness image_filters.SHARPENING_QUALITY) {
	sb.sharpenHardness = sharpenHardness
}

func (sb *SharpenBoard) GetSharpenesHardness() image_filters.SHARPENING_QUALITY {
	return sb.sharpenHardness
}

func (sb *SharpenBoard) SetBrushSize(size int) {
	sb.brushSize = size
}

func (sb *SharpenBoard) GetBrushSize() int {
	return sb.brushSize
}

func (sb *SharpenBoard) CanSharpen() bool {
	return sb.sharpened
}

func (sb *SharpenBoard) ToggleEraserBrush() {
	sb.inEraserBrush = !sb.inEraserBrush
}

func (sb *SharpenBoard) IsEraserBrush() bool {
	return sb.inEraserBrush
}

func (sb *SharpenBoard) GetBoard() image.Image {
	return sb.sharpenedImage
}

func (sb *SharpenBoard) ClearBoard() {
	sb.sharpenedImage = image.NewRGBA64(image.Rect(0, 0, sb.sharpenedImage.Bounds().Dx(), sb.sharpenedImage.Bounds().Dy()))
	sb.sharpenCanvas.Image = sb.sharpenedImage
	sb.sharpenCanvas.Refresh()
}

func (sb *SharpenBoard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(sb.sharpenCanvas)
}
