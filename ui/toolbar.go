package ui

import (
	"image"
	"image/color"
	"photo-man/assets"
	eventActions "photo-man/event-actions"
	event_actions "photo-man/event-actions"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TopToolbar(st *state.AppState) *fyne.Container {

	collapseItem := widget.NewToolbarAction(theme.MenuIcon(), func() {
		go fyne.Do(func() {
			eventActions.ToggleRightSideBarVisibility(st)
		})
	}).ToolbarObject()

	// box container
	hBoxMainDialogContainer := container.NewHBox()

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background1 := canvas.NewRectangle(bgColor)
	background2 := canvas.NewRectangle(bgColor)

	stack1 := container.NewStack(background1, container.NewPadded(container.NewCenter(hBoxMainDialogContainer)))
	stack2 := container.NewStack(background2, collapseItem)

	st.SetToolDialogContainers(hBoxMainDialogContainer)
	return container.NewBorder(nil, nil, nil, stack2, stack1)
}

func CropImageDialog(st *state.AppState, itemWidget *customUI.ActionItemWidget) *fyne.Container {
	label := widget.NewLabel("Crop Image: ")
	apply := widget.NewButtonWithIcon("Apply", assets.Apply, func() {
		var newImg image.Image
		if st.CanvasState.GetCropState().CanCrop() {
			newImg = event_actions.CropImageAction(st.CanvasState.GetCurrentImage(), st.CanvasState.GetCropState())
			st.CanvasState.UpdateSceneImage(newImg)
			event_actions.RemoveCropImageCanvas(st)
			st.RemoveToolDialog()
			itemWidget.Importance = widget.LowImportance
			itemWidget.Refresh()
		}
	})

	discard := widget.NewButtonWithIcon("Discard", assets.Discard, func() {
		event_actions.RemoveCropImageCanvas(st)
		st.RemoveToolDialog()
		itemWidget.Importance = widget.LowImportance
		itemWidget.Refresh()
	})

	return container.NewHBox(label, apply, discard)
}


