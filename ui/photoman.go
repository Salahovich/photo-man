package ui

import (
	"image/color"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func StartApp() {
	myApp := app.New()
	masterWindow := myApp.NewWindow("Photo-Man")
	masterWindow.SetMaster()
	masterWindow.Resize(fyne.NewSize(1400, 800))
	masterWindow.SetPadded(false)

	appState := state.NewAppState()

	middle := ViewPortContainer(appState)
	right := Sidebar(appState)
	top := Toolbar(appState)

	splitter := container.NewHSplit(middle, right)
	splitter.SetOffset(0.8)

	mainContainer := container.NewBorder(top, nil, nil, nil, splitter)

	bgColor := color.RGBA{R: 44, G: 44, B: 44, A: 255}
	background := canvas.NewRectangle(bgColor)

	masterWindow.SetContent(container.NewStack(background, mainContainer))
	masterWindow.ShowAndRun()
}
