package event_actions

import (
	"bytes"
	"image/png"
	"os"
	"photo-man/core/image_io"
	"photo-man/core/image_transform"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"golang.design/x/clipboard"
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
			st.SetImage(img, newImg)
		} else {
			st.SetImage(img, img)
		}

		if readerErr := reader.Close(); readerErr != nil {
			dialog.ShowError(readerErr, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

	}, fyne.CurrentApp().Driver().AllWindows()[0])

	fileDialog.Show()
}

func CopyImageAction(st *state.AppState) {
	err := clipboard.Init()
	if err != nil {
		return
	}
	imgBuffer := new(bytes.Buffer)
	if err := png.Encode(imgBuffer, st.GetCurrentImage()); err != nil {
		return
	}
	clipboard.Write(clipboard.FmtImage, imgBuffer.Bytes())
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
		st.ApplyAllModification()
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

func ResetImage(st *state.AppState) {
	st.UpdateSceneImage(st.GetScaledImage())
}

func CloseImage(st *state.AppState) {
	st.SetImage(nil, nil)
}

func RotateClockwiseAction(st *state.AppState) {
	rotImg := image_transform.RotateClockwise(st.GetCurrentImage())
	st.RegisterListener(image_transform.RotateClockwise)
	st.UpdateSceneImage(rotImg)
}
func RotateAntiClockwiseAction(st *state.AppState) {
	rotImg := image_transform.RotateAntiClockwise(st.GetCurrentImage())
	st.RegisterListener(image_transform.RotateAntiClockwise)
	st.UpdateSceneImage(rotImg)
}
func FlipHorizontallyAction(st *state.AppState) {
	flpImg := image_transform.FlipHorizontally(st.GetCurrentImage())
	st.RegisterListener(image_transform.FlipHorizontally)
	st.UpdateSceneImage(flpImg)
}
func FlipVerticallyAction(st *state.AppState) {
	flpImg := image_transform.FlipVertically(st.GetCurrentImage())
	st.RegisterListener(image_transform.FlipVertically)
	st.UpdateSceneImage(flpImg)
}
