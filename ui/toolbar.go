package ui

import (
	"image/color"
	eventActions "photo-man/event-actions"
	"photo-man/state"

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
	hBoxMainToolsContainer := container.NewHBox()

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background1 := canvas.NewRectangle(bgColor)
	background2 := canvas.NewRectangle(bgColor)

	stack1 := container.NewStack(background1, container.NewPadded(container.NewCenter(hBoxMainToolsContainer)))
	stack2 := container.NewStack(background2, collapseItem)

	return container.NewBorder(nil, nil, nil, stack2, stack1)
}
