package event_actions

import (
	"os"
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
		img, format, imgErr := image_io.ReadImage(imagePath)
		if imgErr != nil {
			dialog.ShowError(imgErr, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		st.SetFormat(format)
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

func ExportImageAction(st *state.AppState) {
	saveFileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer == nil || err != nil {
			return
		}

		// get the save file path and delete the new file created
		fileSavePath := writer.URI().Path()
		if err := os.Remove(fileSavePath); err != nil {
			return
		}

		if err := image_io.WriteImage(st.GetOriginalImage(), fileSavePath, st.GetFormat()); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		}

		if err := writer.Close(); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
	}, fyne.CurrentApp().Driver().AllWindows()[0])

	saveFileDialog.Show()
}
