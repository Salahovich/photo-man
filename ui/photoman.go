package ui

import (
	"image/color"
	event_actions "photo-man/event-actions"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func StartApp() {
	myApp := app.New()
	masterWindow := myApp.NewWindow("PhotoMan")
	masterWindow.SetMaster()
	masterWindow.Resize(fyne.NewSize(1400, 800))
	masterWindow.SetPadded(false)

	appState := state.NewAppState(masterWindow)

	middle := ViewPortContainer(appState)
	right := RightSidebar(appState)
	left := LeftSidebar(appState)
	top := TopToolbar(appState)
	bottom := BottomToolbar(appState)

	mainCanvas := container.NewBorder(top, bottom, nil, nil, middle)
	mainContainer := container.NewBorder(nil, nil, left, right, mainCanvas)

	appState.SetAppContainers([]*fyne.Container{left, right, top, bottom})
	event_actions.SetupShortcutsActions(appState)

	bgColor := color.RGBA{R: 44, G: 44, B: 44, A: 255}
	background := canvas.NewRectangle(bgColor)

	masterWindow.SetContent(container.NewStack(background, mainContainer))
	masterWindow.ShowAndRun()
}
