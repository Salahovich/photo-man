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

func TopToolbar(st *state.AppState) *fyne.Container {

	// props toolbar items
	openItem := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		go eventActions.OpenImageAction(st)
	})

	copyItem := widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
		go eventActions.CopyImageAction(st)
	})

	pasteItem := widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
		go eventActions.PasteImageAction(st)
	})

	toolBarOne := widget.NewToolbar(openItem, copyItem, pasteItem)

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
		go eventActions.ExportImageAction(st)
	})
	closeItem := widget.NewToolbarAction(theme.WindowCloseIcon(), func() {
		go eventActions.CloseImage(st)
	})
	resetItem := widget.NewToolbarAction(theme.ContentUndoIcon(), func() {
		go eventActions.ResetImage(st)
	})
	collapseItem := widget.NewToolbarAction(theme.MenuIcon(), func() {
		go fyne.Do(func() {
			eventActions.ToggleRightSideBarVisibility(st)
		})
	}).ToolbarObject()

	toolBarThree := widget.NewToolbar(exportItem, closeItem, resetItem)

	// separators
	separatorOne := widget.NewToolbarSeparator().ToolbarObject()
	separatorTwo := widget.NewToolbarSeparator().ToolbarObject()

	// box container
	hBoxMainToolsContainer := container.NewHBox(
		toolBarOne,
		separatorOne,
		toolBarTwo,
		separatorTwo,
		toolBarThree)

	// collapseItemhBoxContainer := container.NewHBox(collapseItem)

	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background1 := canvas.NewRectangle(bgColor)
	background2 := canvas.NewRectangle(bgColor)

	stack1 := container.NewStack(background1, container.NewPadded(container.NewCenter(hBoxMainToolsContainer)))
	stack2 := container.NewStack(background2, collapseItem)

	return container.NewBorder(nil, nil, nil, stack2, stack1)
}
