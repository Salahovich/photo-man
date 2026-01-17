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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TopToolbar(st *state.AppState) *fyne.Container {

	collapseItem := widget.NewToolbarAction(theme.MenuIcon(), func() {
		go fyne.Do(func() {
			event_actions.ToggleRightSideBarVisibility(st)
		})
	}).ToolbarObject()

	// box container
	hBoxMainDialogContainer := container.NewCenter()

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background1 := canvas.NewRectangle(bgColor)
	background2 := canvas.NewRectangle(bgColor)

	stack1 := container.NewStack(background1, container.NewPadded(hBoxMainDialogContainer))
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

func PaintBoardDialog(st *state.AppState, itemWidget *customUI.ActionItemWidget) *fyne.Container {
	label := widget.NewLabel("Paint Board:")
	label.Alignment = fyne.TextAlignLeading
	label.TextStyle = fyne.TextStyle{Bold: true}

	// color picker
	colorLabel := widget.NewLabel("color")
	colorPicker := customUI.NewCustomColorPicker(fyne.NewSize(60, 10), func(choosen color.Color) {
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetColor(choosen)
	})
	colorPickerWrapper := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 35)), colorPicker)

	// brush slider
	brushLabel := widget.NewLabel("brush size")
	brushSlider := widget.NewSlider(1, 10)
	brushSlider.OnChanged = func(value float64) {
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetBrushSize(int(value))
	}
	sliderWrapper := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 35)), brushSlider)

	// apply button
	apply := widget.NewButtonWithIcon("Apply", assets.Apply, func() {
		var newImg image.Image
		if st.CanvasState.GetPaintBoardState().CanPaint() {
			newImg = event_actions.PaintBoardAction(st.CanvasState.GetCurrentImage(), st.CanvasState.GetPaintBoardState())
			st.CanvasState.UpdateSceneImage(newImg)
			event_actions.RemovePaintBoardCanvas(st)
			st.RemoveToolDialog()
			itemWidget.Importance = widget.LowImportance
			itemWidget.Refresh()
		}
	})

	// clear widget
	clear := widget.NewToolbarAction(assets.Discard, func() {
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().ClearBoard()
	}).ToolbarObject()

	// eraser widget
	var eraser *widget.Button
	eraser = widget.NewButtonWithIcon("", assets.Eraser, func() {
		if eraser.Importance == widget.LowImportance {
			eraser.Importance = widget.HighImportance
		} else {
			eraser.Importance = widget.LowImportance
		}
		eraseState := st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().IsInEraserState()
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetInEraserState(!eraseState)
		eraser.Refresh()
	})
	eraser.Importance = widget.LowImportance

	// discard widget
	discard := widget.NewButtonWithIcon("Discard", assets.Discard, func() {
		event_actions.RemovePaintBoardCanvas(st)
		st.RemoveToolDialog()
		itemWidget.Importance = widget.LowImportance
		itemWidget.Refresh()
	})

	// spacers
	spacer1 := container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 35)), widget.NewToolbarSpacer().ToolbarObject())
	spacer2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(70, 35)), widget.NewToolbarSpacer().ToolbarObject())

	return container.NewHBox(
		label, spacer1,
		colorLabel, colorPickerWrapper,
		brushLabel, sliderWrapper,
		clear, eraser,
		spacer2,
		apply, discard)
}
