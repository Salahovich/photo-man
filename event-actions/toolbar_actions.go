package event_actions

import (
	"bytes"
	"image"
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
		OpenImageWithPath(imagePath, st)

		if readerErr := reader.Close(); readerErr != nil {
			dialog.ShowError(readerErr, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

	}, fyne.CurrentApp().Driver().AllWindows()[0])

	fileDialog.Show()
}

func OpenImageWithPath(imagePath string, st *state.AppState) {
	img, format, imgErr := image_io.ReadImage(imagePath)
	if imgErr != nil {
		dialog.ShowError(imgErr, fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}
	ScaleAndViewImage(img, format, st)

}

func ScaleAndViewImage(img image.Image, format string, st *state.AppState) {
	st.CanvasState.SetImageInCanvs(true)
	st.CanvasState.SetFormat(format)
	if ok, newImg := image_io.Rescale(img); ok {
		st.CanvasState.SetImage(img, newImg)
	} else {
		st.CanvasState.SetImage(img, img)
	}
	st.Reset()
}

func CopyImageAction(st *state.AppState) {
	err := clipboard.Init()
	if err != nil {
		return
	}
	imgBuffer := new(bytes.Buffer)
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	if err := png.Encode(imgBuffer, st.CanvasState.GetCurrentImage()); err != nil {
		return
	}
	clipboard.Write(clipboard.FmtImage, imgBuffer.Bytes())
}

func PasteImageAction(st *state.AppState) {
	err := clipboard.Init()
	if err != nil {
		return
	}

	imgBuffer := clipboard.Read(clipboard.FmtImage)
	img, err := png.Decode(bytes.NewReader(imgBuffer))

	if err == nil {
		ScaleAndViewImage(img, "png", st)
	}
}

func ExportImageAction(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	saveFileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer == nil || err != nil {
			return
		}

		// get the save file path and delete the new file created
		fileSavePath := writer.URI().Path()
		if err := os.Remove(fileSavePath); err != nil {
			return
		}

		// save the image with all modifications applied
		go func() {
			img := st.ApplyAllModificationOnOriginalImage()
			if err := image_io.WriteImage(img, fileSavePath, st.CanvasState.GetFormat()); err != nil {
				dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			}
		}()

		if err := writer.Close(); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
	}, fyne.CurrentApp().Driver().AllWindows()[0])

	saveFileDialog.Show()
}

func ResetImage(st *state.AppState) {
	st.Reset()
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	st.CanvasState.UpdateSceneImage(st.CanvasState.GetScaledImage())
}

func CloseImage(st *state.AppState) {
	st.Reset()
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	st.CanvasState.SetImageInCanvs(false)
	st.CanvasState.SetImage(nil, nil)
}

func RotateClockwiseAction(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	rotImg := image_transform.RotateClockwise(st.CanvasState.GetCurrentImage())
	st.CanvasState.SetScaledImage(image_transform.RotateClockwise(st.CanvasState.GetScaledImage()))
	st.CanvasState.UpdateSceneImage(rotImg)
}

func RotateAntiClockwiseAction(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	rotImg := image_transform.RotateAntiClockwise(st.CanvasState.GetCurrentImage())
	st.CanvasState.SetScaledImage(image_transform.RotateAntiClockwise(st.CanvasState.GetScaledImage()))
	st.CanvasState.UpdateSceneImage(rotImg)

}

func FlipHorizontallyAction(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	flpImg := image_transform.FlipHorizontally(st.CanvasState.GetCurrentImage())
	st.CanvasState.SetScaledImage(image_transform.FlipHorizontally(st.CanvasState.GetScaledImage()))
	st.CanvasState.UpdateSceneImage(flpImg)
}

func FlipVerticallyAction(st *state.AppState) {
	if !st.CanvasState.IsImageInCanvas() {
		return
	}
	flpImg := image_transform.FlipVertically(st.CanvasState.GetCurrentImage())
	st.CanvasState.SetScaledImage(image_transform.FlipVertically(st.CanvasState.GetScaledImage()))
	st.CanvasState.UpdateSceneImage(flpImg)
}

func ToggleRightSideBarVisibility(st *state.AppState) {
	// fetch the right side bar
	rightSideBar := st.AppEdgeContainers[3]
	if rightSideBar.Hidden {
		rightSideBar.Show()
		rightSideBar.Refresh()
	} else {
		rightSideBar.Hide()
		rightSideBar.Refresh()
	}
	st.AppWindow.Content().Refresh()
}
