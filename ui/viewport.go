package ui

import (
	"photo-man/assets"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ViewPortContainer(st *state.AppState) *fyne.Container {
	vp := NewViewport(st)
	return vp.mainContainer
}

type Viewport struct {
	mainContainer *fyne.Container
	imageView     *canvas.Image
	placeholder   *fyne.Container
}

func NewViewport(st *state.AppState) *Viewport {
	imgCanvas := canvas.NewImageFromResource(nil)
	imgCanvas.FillMode = canvas.ImageFillOriginal

	noPhotoIcon := canvas.NewImageFromResource(assets.NoPhoto)
	noPhotoIcon.FillMode = canvas.ImageFillCover
	noPhotoIcon.SetMinSize(fyne.NewSize(200, 200))
	placeholderContent := container.NewVBox(
		noPhotoIcon,
		widget.NewLabelWithStyle("Click 'Open' to start editing", fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	placeholderWrapper := container.NewCenter(placeholderContent)
	mainContainer := container.NewCenter(placeholderWrapper, imgCanvas)

	viewPortState := &Viewport{
		mainContainer: mainContainer,
		imageView:     imgCanvas,
		placeholder:   placeholderWrapper,
	}

	viewPortState.UpdateViewPortImage(st)
	return viewPortState
}

func (v *Viewport) UpdateViewPortImage(st *state.AppState) {
	go func() {
		for image := range st.GetChannel() {
			if image == nil {
				fyne.Do(func() {
					v.Clear()
				})
			} else {
				w, h := float32(image.Bounds().Dx()), float32(image.Bounds().Dy())
				fyne.Do(func() {
					v.imageView.Image = image
					v.imageView.SetMinSize(fyne.Size{Width: w, Height: h})

					v.imageView.Refresh()
					v.placeholder.Hide()
				})
			}

		}
	}()
}

func (v *Viewport) Clear() {
	v.imageView.Image = nil
	v.imageView.Refresh()
	v.placeholder.Show()
}
