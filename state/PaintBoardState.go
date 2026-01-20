package state

import (
	customUI "photo-man/ui/custom-ui"
)

type PaintBoardState struct {
	isPaintState     bool
	paintBoardCanvas *customUI.PaintBoard
}

func (pb *PaintBoardState) EnablePaintBoard() {
	pb.isPaintState = true
}

func (pb *PaintBoardState) DisablePaintBoard() {
	pb.isPaintState = false
}

func (pb *PaintBoardState) IsInPaintBoard() bool {
	return pb.isPaintState
}

func (pb *PaintBoardState) SetCropImageCanvas(board *customUI.PaintBoard) {
	pb.paintBoardCanvas = board
}

func (pb *PaintBoardState) CanPaint() bool {
	return pb.paintBoardCanvas.IsPainted()
}
func (pb *PaintBoardState) InitPaintBoardState() {
	pb.isPaintState = false
}

func (pb *PaintBoardState) GetPaintBoardCanvas() *customUI.PaintBoard {
	return pb.paintBoardCanvas
}
