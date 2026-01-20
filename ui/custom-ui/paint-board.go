package customUI

import (
	"image"
	"image/color"

	"math"

	"photo-man/core/image_io"

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

type BRUSH_TYPE int

const (
	BRUSH_TYPE_RECTANGLE = 1
	BRUSH_TYPE_CIRCLE    = 2
	BRUSH_TYPE_ERASER    = 3
)

type PaintBoard struct {
	widget.BaseWidget
	board     *image.RGBA64
	canvas    *canvas.Image
	color     *image_io.SystemColor
	painted   bool
	brushType BRUSH_TYPE
	brushSize int
}

func NewPaintBoard(size fyne.Size, col *image_io.SystemColor) *PaintBoard {

	// init the board
	var paintBoard *PaintBoard

	// init the main widget
	paintBoard = &PaintBoard{
		board:     image.NewRGBA64(image.Rect(0, 0, int(size.Width), int(size.Height))),
		color:     col,
		brushSize: 1.0,
		painted:   false,
	}
	paintBoard.canvas = canvas.NewImageFromImage(paintBoard.board)

	paintBoard.ExtendBaseWidget(paintBoard)

	return paintBoard
}

func (pb *PaintBoard) DragEnd() {}
func (pb *PaintBoard) Dragged(ev *fyne.DragEvent) {
	if pb.brushType == 0 || pb.brushSize == 0 {
		return
	}

	pixel := ev.Position
	r, g, b, a := pb.color.Color.RGBA()
	targetColor := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
	radius := float64(pb.brushSize)

	// position a square of the brush size as its side!
	startPos := fyne.Position{X: float32(pixel.X - float32(pb.brushSize)), Y: float32(pixel.Y - float32(pb.brushSize))}
	endPos := fyne.Position{X: float32(pixel.X + float32(pb.brushSize)), Y: float32(pixel.Y + float32(pb.brushSize))}

	for y := int(startPos.Y); y < int(endPos.Y); y++ {
		for x := int(startPos.X); x < int(endPos.X); x++ {
			if x < 0 || x >= pb.board.Bounds().Dx() || y < 0 || y >= pb.board.Bounds().Dy() {
				continue
			}
			switch pb.brushType {
			case BRUSH_TYPE_RECTANGLE:
				pb.board.SetRGBA64(x, y, targetColor)
			case BRUSH_TYPE_CIRCLE:
				distance := math.Round(math.Hypot(float64(pixel.X)-float64(x), float64(pixel.Y)-float64(y)))
				if distance < radius {
					pb.board.SetRGBA64(x, y, targetColor)
				}
			case BRUSH_TYPE_ERASER:
				pb.board.SetRGBA64(x, y, color.RGBA64{})
			}
		}
	}
	pb.Refresh()

	defer func() {
		if !pb.painted {
			pb.painted = true
		}
	}()
}

func (pb *PaintBoard) MouseMoved(ev *desktop.MouseEvent) {

}
func (pb *PaintBoard) MouseIn(ev *desktop.MouseEvent) {

}
func (pb *PaintBoard) MouseOut() {

}

func (pb *PaintBoard) Tapped(ev *fyne.PointEvent) {
}

func (pb *PaintBoard) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (pb *PaintBoard) MinSize() fyne.Size {
	return fyne.NewSize(float32(pb.board.Bounds().Dx()), float32(pb.board.Bounds().Dy()))
}

func (pb *PaintBoard) GetBoard() image.Image {
	return pb.board
}

func (pb *PaintBoard) SetBrushSize(size int) {
	pb.brushSize = size
}
func (pb *PaintBoard) GetBrushSize() int {
	return pb.brushSize
}
func (pb *PaintBoard) IsPainted() bool {
	return pb.painted
}

func (pb *PaintBoard) SetBrushType(brushType BRUSH_TYPE) {
	pb.brushType = brushType
}

func (pb *PaintBoard) GetBrushType() BRUSH_TYPE {
	return pb.brushType
}

func (pb *PaintBoard) SetNoBrush() {
	pb.brushType = 0
}

func (pb *PaintBoard) ClearBoard() {
	pb.painted = false
	pb.board = image.NewRGBA64(image.Rect(0, 0, pb.board.Bounds().Dx(), pb.board.Bounds().Dy()))
	pb.canvas.Image = pb.board
	pb.canvas.Refresh()
}

func (pb *PaintBoard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(pb.canvas)
}
