package customUI

import (
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

type ResizableRectangle struct {
	widget.BaseWidget
	Rect              *canvas.Rectangle
	canCrop           bool
	left              bool
	top               bool
	right             bool
	bottom            bool
	startPosition     fyne.Position
	endPosition       fyne.Position
	currStartPosition fyne.Position
	currEndPosition   fyne.Position
	minCropSize       fyne.Size
	previousXMove     float32
	previousYMove     float32
}

func NewResizableRectangle(size fyne.Size, col color.Color) *ResizableRectangle {

	// init the rectangle
	var item *ResizableRectangle
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(size)
	rect.StrokeWidth = 5
	rect.StrokeColor = color.White

	// init the main widget
	item = &ResizableRectangle{
		Rect:              rect,
		canCrop:           false,
		left:              false,
		top:               false,
		right:             false,
		bottom:            false,
		endPosition:       fyne.NewPos(rect.MinSize().Width, rect.MinSize().Height),
		currEndPosition:   fyne.NewPos(rect.MinSize().Width, rect.MinSize().Height),
		startPosition:     rect.Position(),
		currStartPosition: rect.Position(),
	}
	item.minCropSize = fyne.NewSize(item.currEndPosition.X/10, item.currEndPosition.Y/10)
	item.ExtendBaseWidget(item)

	return item
}

func (rr *ResizableRectangle) DragEnd() {

}
func (rr *ResizableRectangle) Dragged(ev *fyne.DragEvent) {
	// detect the move direction
	moveUp, moveDown, moveRight, moveLeft := false, false, false, false
	if ev.Position.X > rr.previousXMove {
		moveRight = true
	} else if ev.Position.X < rr.previousXMove {
		moveLeft = true
	}
	if ev.Position.Y > rr.previousYMove {
		moveDown = true
	} else if ev.Position.Y < rr.previousYMove {
		moveUp = true
	}

	if rr.right {
		// move left
		if moveLeft && (rr.currEndPosition.X > rr.currStartPosition.X+rr.minCropSize.Width && rr.currEndPosition.X <= rr.endPosition.X) {
			rr.currEndPosition.X += ev.Dragged.DX
		}
		// move right
		if moveRight && (rr.currEndPosition.X < rr.endPosition.X) {
			rr.currEndPosition.X += ev.Dragged.DX
		}

	} else if rr.left {
		// move right
		if moveRight && (rr.currStartPosition.X >= rr.startPosition.X && rr.currStartPosition.X < rr.endPosition.X-rr.minCropSize.Width) {
			rr.currStartPosition.X += ev.Dragged.DX
		}
		// move left
		if moveLeft && (rr.currStartPosition.X > rr.startPosition.X) {
			rr.currStartPosition.X += ev.Dragged.DX
		}

	} else if rr.top {
		// move down
		if moveDown && (rr.currStartPosition.Y >= rr.startPosition.Y && rr.currStartPosition.Y < rr.endPosition.Y-rr.minCropSize.Height) {
			rr.currStartPosition.Y += ev.Dragged.DY
		}
		// move up
		if moveUp && (rr.currStartPosition.Y > rr.startPosition.Y) {
			rr.currStartPosition.Y += ev.Dragged.DY
		}
	} else if rr.bottom {
		// move up
		if moveUp && (rr.currEndPosition.Y > rr.currStartPosition.Y+rr.minCropSize.Height && rr.currEndPosition.Y <= rr.endPosition.Y) {
			rr.currEndPosition.Y += ev.Dragged.DY
		}
		// move down
		if moveDown && (rr.currEndPosition.Y < rr.endPosition.Y) {
			rr.currEndPosition.Y += ev.Dragged.DY
		}
	}

	rr.previousXMove = ev.Position.X
	rr.previousYMove = ev.Position.Y

	rr.Rect.Move(rr.currStartPosition)
	rr.Rect.Resize(fyne.NewSize(rr.currEndPosition.X-rr.currStartPosition.X, rr.currEndPosition.Y-rr.currStartPosition.Y))
	rr.Refresh()
}

func (rr *ResizableRectangle) MouseMoved(ev *desktop.MouseEvent) {
	rr.detectEdges(ev)
}
func (rr *ResizableRectangle) MouseIn(ev *desktop.MouseEvent) {
	rr.detectEdges(ev)
}
func (rr *ResizableRectangle) MouseOut() {
	rr.left, rr.right, rr.top, rr.bottom = false, false, false, false
}
func (rr *ResizableRectangle) Tapped(ev *fyne.PointEvent) {}

func (rr *ResizableRectangle) detectEdges(ev *desktop.MouseEvent) {
	if ev.Position.X >= rr.currStartPosition.X && ev.Position.X <= (rr.currStartPosition.X+rr.Rect.StrokeWidth) {
		rr.left = true
	} else if ev.Position.X >= (rr.currEndPosition.X-rr.Rect.StrokeWidth) && ev.Position.X <= rr.currEndPosition.X {
		rr.right = true
	} else {
		rr.left = false
		rr.right = false
	}

	if ev.Position.Y >= rr.currStartPosition.Y && ev.Position.Y <= (rr.currStartPosition.Y+rr.Rect.StrokeWidth) {
		rr.top = true
	} else if ev.Position.Y >= (rr.currEndPosition.Y-rr.Rect.StrokeWidth) && ev.Position.Y <= rr.currEndPosition.Y {
		rr.bottom = true
	} else {
		rr.top = false
		rr.bottom = false
	}
}
func (rr *ResizableRectangle) Cursor() desktop.Cursor {
	if rr.left || rr.right {
		return desktop.HResizeCursor
	} else if rr.top || rr.bottom {
		return desktop.VResizeCursor
	} else {
		return desktop.DefaultCursor
	}
}

func (rr *ResizableRectangle) MinSize() fyne.Size {
	return rr.Rect.MinSize()
}

func (rr *ResizableRectangle) GetRect() *canvas.Rectangle {
	return rr.Rect
}
func (rr *ResizableRectangle) GetCurrStartPosition() fyne.Position {
	return fyne.NewPos(rr.currStartPosition.X+rr.Rect.StrokeWidth, rr.currStartPosition.Y+rr.Rect.StrokeWidth)
}
func (rr *ResizableRectangle) GetCurrEndPosition() fyne.Position {
	return fyne.NewPos(rr.currEndPosition.X-rr.Rect.StrokeWidth, rr.currEndPosition.Y-rr.Rect.StrokeWidth)
}
func (rr *ResizableRectangle) CanCrop() bool {
	return rr.currStartPosition.X != rr.startPosition.X ||
		rr.currStartPosition.Y != rr.startPosition.Y ||
		rr.currEndPosition.X != rr.endPosition.X ||
		rr.currEndPosition.Y != rr.endPosition.Y
}
func (rr *ResizableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(rr.Rect)
}
