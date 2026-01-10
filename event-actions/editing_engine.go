package event_actions

import "photo-man/state"

func PerformEdit(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}

	scaledImg := st.CanvasState.GetScaledImage()

	// Run all edits sequentially
	img := UpdateAdjustmentsAction(scaledImg, st.AdjustmentState, st.AdjustmentFactors)
	img = UpdateBaseFiltersAction(img, st.BasicFilterState)
	img = UpdateColorBlendingAction(img, st.ColorBlendState)

	st.CanvasState.UpdateSceneImage(img)
}
