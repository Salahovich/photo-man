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

	openItem := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		go event_actions.OpenImageAction(st)
	}).ToolbarObject()

	copyItem := widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
		go event_actions.CopyImageAction(st)
	}).ToolbarObject()

	pasteItem := widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
		go event_actions.PasteImageAction(st)
	}).ToolbarObject()

	exportItem := widget.NewToolbarAction(theme.DownloadIcon(), func() {
		go event_actions.ExportImageAction(st)
	}).ToolbarObject()

	closeItem := widget.NewToolbarAction(theme.WindowCloseIcon(), func() {
		go event_actions.CloseImage(st)
	}).ToolbarObject()
	resetItem := widget.NewToolbarAction(theme.ContentUndoIcon(), func() {
		go event_actions.ResetImage(st)
	}).ToolbarObject()

	separatorOne := widget.NewToolbarSeparator().ToolbarObject()
	separatorTwo := widget.NewToolbarSeparator().ToolbarObject()
	separatorThree := widget.NewToolbarSeparator().ToolbarObject()

	toolBar := container.NewHBox(
		openItem, copyItem, pasteItem, exportItem, separatorOne,
		undoAction, redoAction, closeItem, resetItem, separatorTwo,
		zoomInAction, zoomOutAction, separatorThree,
		fullScreenAction)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(toolBar)
	mainContainer = container.NewStack(background, container.NewPadded(centerContainer))

	return mainContainer
}
