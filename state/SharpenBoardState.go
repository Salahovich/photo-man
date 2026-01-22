package state

import (
	customUI "photo-man/ui/custom-ui"
)

type SharpenBoardState struct {
	isSharpenState     bool
	sharpenBoardCanvas *customUI.SharpenBoard
}

func (sb *SharpenBoardState) EnableSharpenBoard() {
	sb.isSharpenState = true
}
func (sb *SharpenBoardState) DisableSharpenBoard() {
	sb.isSharpenState = false
}

func (sb *SharpenBoardState) IsInSharpenBoard() bool {
	return sb.isSharpenState
}

func (sb *SharpenBoardState) SetSharpenBoardCanvas(board *customUI.SharpenBoard) {
	sb.sharpenBoardCanvas = board
}

func (sb *SharpenBoardState) CanSharpen() bool {
	return sb.sharpenBoardCanvas.CanSharpen()
}

func (sb *SharpenBoardState) InitSharpenBoardState() {
	sb.isSharpenState = false
}

func (sb *SharpenBoardState) GetSharpenBoardCanvas() *customUI.SharpenBoard {
	return sb.sharpenBoardCanvas
}
