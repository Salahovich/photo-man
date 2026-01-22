package state

type EyeDropState struct {
	isEyeDropState bool
}

func (ed *EyeDropState) ToggleEyeDropState() {
	ed.isEyeDropState = !ed.isEyeDropState
}

func (ed *EyeDropState) IsInEyeDropState() bool {
	return ed.isEyeDropState
}
