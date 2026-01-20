package ui

import (
	"image/color"
	"photo-man/assets"
	event_actions "photo-man/event-actions"
	"photo-man/state"

	// "photo-man/ui"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LeftSidebar(st *state.AppState) *fyne.Container {

	// functionality toolbar items
	var transformations *customUI.ActionItemWidget
	transformations = customUI.NewActionItemWidget(assets.Transformation, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
		if !st.Transformations.IsInTransformationState() {
			st.ShowToolDialog(TransformationsDialog(st, transformations))
		} else {
			st.RemoveToolDialog()
		}
		st.Transformations.ControlTransformationState()
	})

	var cropAction *customUI.ActionItemWidget
	cropAction = customUI.NewActionItemWidget(assets.Crop, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
		if !st.CanvasState.GetCropState().IsInCropState() {
			st.ShowToolDialog(CropImageDialog(st, cropAction))
			event_actions.InitCropImageCanvas(st)
		} else {
			event_actions.RemoveCropImageCanvas(st)
			st.RemoveToolDialog()
		}
	})

	var brushAction *customUI.ActionItemWidget
	brushAction = customUI.NewActionItemWidget(assets.Brush, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
		if !st.CanvasState.GetPaintBoardState().IsInPaintBoard() {
			event_actions.InitPaintBoardCanvas(st)
			st.ShowToolDialog(PaintBoardDialog(st, brushAction))
		} else {
			event_actions.RemovePaintBoardCanvas(st)
			st.RemoveToolDialog()
		}
	})

	var blurItem *customUI.ActionItemWidget
	blurItem = customUI.NewActionItemWidget(assets.Blur, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
	})

	var sharpenItem *customUI.ActionItemWidget
	sharpenItem = customUI.NewActionItemWidget(assets.Sharpen, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
	})

	var cloneItem *customUI.ActionItemWidget
	cloneItem = customUI.NewActionItemWidget(assets.Clone, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
	})

	var eyeDropItem *customUI.ActionItemWidget
	eyeDropItem = customUI.NewActionItemWidget(assets.EyeDrop, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
	})

	var bucketItem *customUI.ActionItemWidget
	bucketItem = customUI.NewActionItemWidget(assets.Bucket, func() {
		if !st.CanvasState.IsImageInCanvas() {
			return
		}
	})

	textAction := customUI.NewActionItemWidget(assets.Text, func() {

	})

	paletteItem := customUI.NewCustomColorPicker(fyne.NewSize(30, 30), func(choosen color.Color) {
		st.SystemColor.Color = choosen
	})

	// box container
	verticalActionItemList := customUI.NewActionItemList(
		true,
		true,
		transformations,
		cropAction,
		brushAction,
		blurItem,
		sharpenItem,
		cloneItem,
		bucketItem,
		eyeDropItem,
		textAction,
	)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(
		container.NewVBox(
			verticalActionItemList.Box,
			widget.NewToolbarSeparator().ToolbarObject(),
			container.NewGridWrap(fyne.NewSize(30, 5), widget.NewToolbarSpacer().ToolbarObject()),
			container.NewCenter(container.NewGridWrap(fyne.NewSize(30, 30), paletteItem))))

	return container.NewStack(background, container.NewPadded(centerContainer))
}
