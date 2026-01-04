package ui

import (
	"image/color"
	"photo-man/assets"
	eventActions "photo-man/event-actions"
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
		eventActions.OpenImageAction(st)
	})

	copyItem := widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
		eventActions.CopyImageAction(st)
	})
	closeItem := widget.NewToolbarAction(theme.WindowCloseIcon(), func() {
		eventActions.CloseImage(st)
	})
	resetItem := widget.NewToolbarAction(theme.ContentUndoIcon(), func() {
		eventActions.ResetImage(st)
	})
	toolBarOne := widget.NewToolbar(openItem, copyItem, closeItem, resetItem)

	// functionality toolbar items
	rotateLeftItem := widget.NewToolbarAction(theme.MediaReplayIcon(), func() {
		st.Transformations.RotateAntiClockwise()
		go eventActions.RotateAntiClockwiseAction(st)
	})
	rotateRightItem := widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
		st.Transformations.RotateClockwise()
		go eventActions.RotateClockwiseAction(st)

	})
	flipHorizontallyItem := widget.NewToolbarAction(assets.FlipRight, func() {
		st.Transformations.FlipHorizontally()
		go eventActions.FlipHorizontallyAction(st)

	})
	flipVerticallyItem := widget.NewToolbarAction(assets.FlipDown, func() {
		st.Transformations.FlipHVertically()
		go eventActions.FlipVerticallyAction(st)
	})

	toolBarTwo := widget.NewToolbar(
		rotateLeftItem,
		rotateRightItem,
		flipHorizontallyItem,
		flipVerticallyItem)

	// export toolbar items
	exportItem := widget.NewToolbarAction(theme.DownloadIcon(), func() {
		eventActions.ExportImageAction(st)
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
