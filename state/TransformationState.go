package state

type TransformationState struct {
	Rotate                int
	FlipVertical          bool
	FlipHorizontal        bool
	inTransformationState bool
}

func (ts *TransformationState) RotateClockwise() {
	if ts.Rotate == 2 {
		ts.Rotate = -1
	} else {
		ts.Rotate++
	}
}

func (ts *TransformationState) RotateAntiClockwise() {
	if ts.Rotate == -1 {
		ts.Rotate = 2
	} else {
		ts.Rotate--
	}
}

func (ts *TransformationState) FlipHorizontally() {
	ts.FlipHorizontal = !ts.FlipHorizontal
}

func (ts *TransformationState) FlipHVertically() {
	ts.FlipVertical = !ts.FlipVertical
}

func (ts *TransformationState) ControlTransformationState() {
	ts.inTransformationState = !ts.inTransformationState
}
func (ts *TransformationState) IsInTransformationState() bool {
	return ts.inTransformationState
}

func (ts *TransformationState) InitTransformations() {
	ts.Rotate = 0
	ts.FlipVertical = false
	ts.FlipVertical = false
	ts.inTransformationState = false
}
