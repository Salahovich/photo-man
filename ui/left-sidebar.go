package ui

import (
	"image"
	"image/color"
	"photo-man/assets"
	event_actions "photo-man/event-actions"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LeftSidebar(st *state.AppState) *fyne.Container {
	// props toolbar items

	// functionality toolbar items
	cropAction := customUI.NewActionItemWidget(assets.Crop, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
		if !st.CanvasState.GetCropState().IsInCropState() {
			event_actions.InitCropImageCanvas(st)
		} else {
			var newImg image.Image
			if st.CanvasState.GetCropState().CanCrop() {
				newImg = event_actions.CropImageAction(st.CanvasState.GetCurrentImage(), st.CanvasState.GetCropState())
				st.CanvasState.UpdateSceneImage(newImg)
			}
			event_actions.RemoveCropImageCanvas(st)
		}
	})
	rotateLeftItem := widget.NewToolbarAction(theme.MediaReplayIcon(), func() {
		st.Transformations.RotateAntiClockwise()
		go event_actions.RotateAntiClockwiseAction(st)
	}).ToolbarObject()
	rotateRightItem := widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
		st.Transformations.RotateClockwise()
		go event_actions.RotateClockwiseAction(st)

	}).ToolbarObject()
	flipHorizontallyItem := widget.NewToolbarAction(assets.FlipRight, func() {
		st.Transformations.FlipHorizontally()
		go event_actions.FlipHorizontallyAction(st)

	}).ToolbarObject()
	flipVerticallyItem := widget.NewToolbarAction(assets.FlipDown, func() {
		st.Transformations.FlipHVertically()
		go event_actions.FlipVerticallyAction(st)
	}).ToolbarObject()
	textAction := customUI.NewActionItemWidget(assets.Text, func() {

	})

	brushAction := customUI.NewActionItemWidget(assets.Brush, func() {

	})

	eraserAction := customUI.NewActionItemWidget(assets.Eraser, func() {

	})

	paletteAction := customUI.NewActionItemWidget(assets.Palette, func() {

	})

	// box container
	verticalActionItemList := customUI.NewVerticalActionItemList(true,
		cropAction,
		textAction,
		brushAction,
		eraserAction,
		paletteAction)

	verticalTransformations := container.NewVBox(rotateLeftItem, rotateRightItem, flipHorizontallyItem, flipVerticallyItem)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(
		container.NewVBox(verticalTransformations, widget.NewToolbarSeparator().ToolbarObject(), verticalActionItemList.VBox))

	return container.NewStack(background, container.NewPadded(centerContainer))
}
