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

func Toolbar(st *state.AppState) *fyne.Container {

	// props toolbar items
	openItem := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		event_actions.OpenImageAction(st)
	})

	copyItem := widget.NewToolbarAction(theme.ContentCopyIcon(), func() {})
	toolBarOne := widget.NewToolbar(openItem, copyItem)

	// functionality toolbar items
	cropItem := widget.NewToolbarAction(theme.ContentCutIcon(), func() {})
	rotateLeftItem := widget.NewToolbarAction(theme.MediaPlayIcon(), func() {})
	rotateRightItem := widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {})
	flipLeftItem := widget.NewToolbarAction(assets.FlipLeft, func() {})
	flipRightItem := widget.NewToolbarAction(assets.FlipRight, func() {})
	flipUpItem := widget.NewToolbarAction(assets.FlipUp, func() {})
	flipDownItem := widget.NewToolbarAction(assets.FlipDown, func() {})
	toolBarTwo := widget.NewToolbar(
		cropItem,
		rotateLeftItem,
		rotateRightItem,
		flipLeftItem,
		flipRightItem,
		flipUpItem,
		flipDownItem)

	// export toolbar items
	exportItem := widget.NewToolbarAction(theme.DownloadIcon(), func() {
		event_actions.ExportImageAction(st)
	})
	toolBarThree := widget.NewToolbar(exportItem)

	// separators
	separatorOne := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	separatorOne.SetMinSize(fyne.NewSize(1, 1))
	separatorTwo := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	separatorTwo.SetMinSize(fyne.NewSize(1, 1))

	// box container
	hBoxContainer := container.NewHBox(
		toolBarOne,
		separatorOne,
		toolBarTwo,
		separatorTwo,
		toolBarThree)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	centerContainer := container.NewCenter(hBoxContainer)

	return container.NewStack(background, container.NewPadded(centerContainer))
}
