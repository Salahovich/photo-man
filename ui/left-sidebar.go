package ui

import (
	"image/color"
	"photo-man/assets"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func LeftSidebar(st *state.AppState) *fyne.Container {
	// props toolbar items

	// functionality toolbar items
	cropAction := customUI.NewActionItemWidget(assets.Crop, func() {

	})

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

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(verticalActionItemList.VBox)

	return container.NewStack(background, container.NewPadded(centerContainer))
}
