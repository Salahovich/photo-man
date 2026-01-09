package ui

import (
	"image/color"
	"photo-man/assets"
	event_actions "photo-man/event-actions"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BottomToolbar(st *state.AppState) *fyne.Container {

	var mainContainer *fyne.Container

	// props toolbar items
	zoomInAction := widget.NewToolbarAction(theme.ZoomInIcon(), func() {

	}).ToolbarObject()

	zoomOutAction := widget.NewToolbarAction(theme.ZoomOutIcon(), func() {

	}).ToolbarObject()

	undoAction := widget.NewToolbarAction(assets.Undo, func() {

	}).ToolbarObject()

	redoAction := widget.NewToolbarAction(assets.Redo, func() {

	}).ToolbarObject()

	fullScreenAction := widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() {
		event_actions.EnterFullScreenMode(st)
	}).ToolbarObject()

	separatorOne := widget.NewToolbarSeparator().ToolbarObject()
	separatorTwo := widget.NewToolbarSeparator().ToolbarObject()

	toolBar := container.NewHBox(undoAction, redoAction, separatorOne, zoomInAction, zoomOutAction, separatorTwo, fullScreenAction)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(toolBar)
	mainContainer = container.NewStack(background, container.NewPadded(centerContainer))

	return mainContainer
}
