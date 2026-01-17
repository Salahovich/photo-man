package customUI

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// generate compile time error if the interface methods did not implemented.
var _ desktop.Hoverable = (*ResizableRectangle)(nil)
var _ fyne.Draggable = (*ResizableRectangle)(nil)
var _ desktop.Cursorable = (*ResizableRectangle)(nil)

type PaintBoard struct {
	widget.BaseWidget
	board     *image.RGBA64
	color     color.Color
	brushSize int
}

func NewPaintBoard(size fyne.Size, col color.Color) *PaintBoard {

	// init the board
	var paintBoard *PaintBoard

	// init the main widget
	paintBoard = &PaintBoard{
		board:     image.NewRGBA64(image.Rect(0, 0, int(size.Width), int(size.Height))),
		color:     col,
		brushSize: 3.0,
	}

	paintBoard.ExtendBaseWidget(paintBoard)

	return paintBoard
}

func (pb *PaintBoard) DragEnd() {}
func (pb *PaintBoard) Dragged(ev *fyne.DragEvent) {
	pixel := ev.Position
	r, g, b, a := pb.color.RGBA()
	targetColor := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}

	// position a rectangle of the brush size as its side!
	startPos := fyne.Position{X: float32(pixel.X - float32(pb.brushSize)), Y: float32(pixel.Y - float32(pb.brushSize))}
	endPos := fyne.Position{X: float32(pixel.X + float32(pb.brushSize)), Y: float32(pixel.Y + float32(pb.brushSize))}

	for y := int(startPos.Y); y < int(endPos.Y); y++ {
		for x := int(startPos.X); x < int(endPos.X); x++ {
			if x < 0 || x >= pb.board.Bounds().Dx() || y < 0 || y >= pb.board.Bounds().Dy() {
				continue
			}
			pb.board.SetRGBA64(x, y, targetColor)
		}
	}
	pb.Refresh()
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

func (pb *PaintBoard) SetColor(col color.Color) {
	pb.color = col
}

func (pb *PaintBoard) GetColorcol(col color.Color) {
	pb.color = col
}

func (pb *PaintBoard) SetBrushSize(size int) {
	pb.brushSize = size
}
func (pb *PaintBoard) GetBrushSize() int {
	return pb.brushSize
}

func (pb *PaintBoard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(canvas.NewImageFromImage(pb.board))
}
