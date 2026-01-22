package state

import (
	customUI "photo-man/ui/custom-ui"
)

type BlurBoardState struct {
	isBlurState     bool
	blurBoardCanvas *customUI.BlurBoard
}

func (bb *BlurBoardState) EnableBlurBoard() {
	bb.isBlurState = true
}
func (bb *BlurBoardState) DisableBlurBoard() {
	bb.isBlurState = false
}

func (bb *BlurBoardState) IsInBlurBoard() bool {
	return bb.isBlurState
}

func (bb *BlurBoardState) SetBlurBoardCanvas(board *customUI.BlurBoard) {
	bb.blurBoardCanvas = board
}

func (bb *BlurBoardState) CanBlur() bool {
	return bb.blurBoardCanvas.CanBlur()
}

func (bb *BlurBoardState) InitBlurBoardState() {
	bb.isBlurState = false
}

func (bb *BlurBoardState) GetBlurBoardCanvas() *customUI.BlurBoard {
	return bb.blurBoardCanvas
}
