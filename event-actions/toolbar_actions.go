package event_actions

import (
	"photo-man/core/image_io"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func OpenImageAction(st *state.AppState) {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if reader == nil {
			return
		}

		imagePath := reader.URI().Path()
		img, _, err := image_io.ReadImage(imagePath)
		if err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		if ok, newImg := image_io.Rescale(img); ok {
			st.SetOriginalImage(img)
			st.SetImage(newImg)
		} else {
			st.SetImage(img)
		}

		if readerErr := reader.Close(); readerErr != nil {
			dialog.ShowError(readerErr, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

	}, fyne.CurrentApp().Driver().AllWindows()[0])

	fileDialog.Show()
}
