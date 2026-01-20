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

	// brush slider
	brushLabel := widget.NewLabel("brush size")
	brushSlider := widget.NewSlider(1, 20)
	brushSlider.OnChanged = func(value float64) {
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetBrushSize(int(value))
	}
	sliderWrapper := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 35)), brushSlider)

	brushTypeLabel := widget.NewLabel("Brush Type")

	var squareBrush *customUI.ActionItemWidget
	squareBrush = customUI.NewActionItemWidget(assets.Square, func() {
		if st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().GetBrushType() == customUI.BRUSH_TYPE_RECTANGLE {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetNoBrush()
		} else {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetBrushType(customUI.BRUSH_TYPE_RECTANGLE)
		}
		squareBrush.Refresh()
	})

	var circleBrush *customUI.ActionItemWidget
	circleBrush = customUI.NewActionItemWidget(assets.Circle, func() {
		if st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().GetBrushType() == customUI.BRUSH_TYPE_CIRCLE {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetNoBrush()
		} else {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetBrushType(customUI.BRUSH_TYPE_CIRCLE)
		}
		circleBrush.Refresh()
	})

	// eraser widget
	var eraserBrush *customUI.ActionItemWidget
	eraserBrush = customUI.NewActionItemWidget(assets.Eraser, func() {
		if st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().GetBrushType() == customUI.BRUSH_TYPE_ERASER {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetNoBrush()
		} else {
			st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().SetBrushType(customUI.BRUSH_TYPE_ERASER)
		}
		eraserBrush.Refresh()
	})

	// clear widget
	clear := widget.NewButtonWithIcon("Clear", assets.Discard, func() {
		st.CanvasState.GetPaintBoardState().GetPaintBoardCanvas().ClearBoard()
	})

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

	brushTypeContainer := customUI.NewActionItemList(true, false, squareBrush, circleBrush, eraserBrush)

	return container.NewHBox(
		label, spacer1,
		brushLabel, sliderWrapper,
		brushTypeLabel, brushTypeContainer.Box,
		spacer2,
		clear, apply, discard)
}

func TransformationsDialog(st *state.AppState, itemWidget *customUI.ActionItemWidget) *fyne.Container {
	var rotateLeftItem *widget.Button
	rotateLeftItem = widget.NewButtonWithIcon("Rotate Left", theme.MediaReplayIcon(), func() {
		st.Transformations.RotateAntiClockwise()
		go event_actions.RotateAntiClockwiseAction(st)
	})
	rotateLeftItem.Importance = widget.MediumImportance

	var rotateRightItem *widget.Button
	rotateRightItem = widget.NewButtonWithIcon("Rotate Right", theme.ViewRefreshIcon(), func() {
		rotateRightItem.Refresh()
		st.Transformations.RotateClockwise()
		go event_actions.RotateClockwiseAction(st)

	})
	rotateRightItem.Importance = widget.MediumImportance

	var flipHorizontallyItem *widget.Button
	flipHorizontallyItem = widget.NewButtonWithIcon("Flip Horizontally", assets.FlipRight, func() {
		flipHorizontallyItem.Refresh()
		st.Transformations.FlipHorizontally()
		go event_actions.FlipHorizontallyAction(st)

	})
	flipHorizontallyItem.Importance = widget.MediumImportance

	var flipVerticallyItem *widget.Button
	flipVerticallyItem = widget.NewButtonWithIcon("Flip Vertically", assets.FlipDown, func() {
		flipVerticallyItem.Refresh()
		st.Transformations.FlipHVertically()
		go event_actions.FlipVerticallyAction(st)
	})
	flipVerticallyItem.Importance = widget.MediumImportance

	return container.NewHBox(rotateLeftItem, rotateRightItem, flipHorizontallyItem, flipVerticallyItem)
}
