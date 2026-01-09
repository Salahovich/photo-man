package event_actions

import (
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func SetupShortcutsActions(st *state.AppState) {

	// file drop
	st.AppWindow.SetOnDropped(func(position fyne.Position, uris []fyne.URI) {
		imagePath := uris[0].Path()
		OpenImageWithPath(imagePath, st)
	})

	// single key actions
	st.AppWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyEscape:
			exitFullScreenMode(st)
			st.AppWindow.Content().Refresh()
		case fyne.KeyF:
			handleFullScreenMode(st)
			st.AppWindow.Content().Refresh()
		}
	})

	// multi key actions

	// CTRL+S --> Export Image
	ctrlS := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
	st.AppWindow.Canvas().AddShortcut(ctrlS, func(shortcut fyne.Shortcut) {
		ExportImageAction(st)
	})
	// CTRL+G --> Paste Image
	ctrlG := &desktop.CustomShortcut{KeyName: fyne.KeyG, Modifier: fyne.KeyModifierControl}
	st.AppWindow.Canvas().AddShortcut(ctrlG, func(shortcut fyne.Shortcut) {
		PasteImageAction(st)
	})

}

func handleFullScreenMode(st *state.AppState) {
	isFullScreen := fyne.CurrentApp().Driver().AllWindows()[0].FullScreen()
	if isFullScreen {
		exitFullScreenMode(st)
	} else {
		EnterFullScreenMode(st)
	}
}

func exitFullScreenMode(st *state.AppState) {
	fyne.CurrentApp().Driver().AllWindows()[0].SetFullScreen(false)

	for _, container := range st.AppContainers {
		container.Show()
		container.Refresh()
	}
}

func EnterFullScreenMode(st *state.AppState) {
	fyne.CurrentApp().Driver().AllWindows()[0].SetFullScreen(true)

	for _, container := range st.AppContainers {
		container.Hide()
		container.Refresh()
	}
}
