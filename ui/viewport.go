package ui

import (
	"photo-man/assets"
	"photo-man/state"
	customUI "photo-man/ui/custom-ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ViewPortContainer(st *state.AppState) *fyne.Container {
	vp := NewViewport(st)
	return container.NewStack(vp.canvasStack)
}

type Viewport struct {
	canvasStack *fyne.Container
	imageView   *customUI.CustomImageCanvas
	placeholder *fyne.Container
}

func NewViewport(st *state.AppState) *Viewport {
	imgCanvas := customUI.NewCustomImageCanvas(st.CanvasState.SystemColor)

	noPhotoIcon := canvas.NewImageFromResource(assets.NoPhoto)
	noPhotoIcon.FillMode = canvas.ImageFillContain
	noPhotoIcon.SetMinSize(fyne.NewSize(200, 200))
	placeholderContent := container.NewVBox(
		noPhotoIcon,
		widget.NewLabelWithStyle(
			"Drop any image \n or press 'Ctrl+G' to paste image \n or click 'Open' to start editing",
			fyne.TextAlignCenter,
			fyne.TextStyle{}))

	placeholderWrapper := container.NewCenter(placeholderContent)
	canvasContainerStack := container.NewCenter(placeholderWrapper, imgCanvas)

	viewPortState := &Viewport{
		canvasStack: canvasContainerStack,
		imageView:   imgCanvas,
		placeholder: placeholderWrapper,
	}

	st.CanvasState.SetCanvasStack(canvasContainerStack)
	st.CanvasState.SetCanvas(imgCanvas)

	// start the View Port channel.
	viewPortState.UpdateViewPortImage(st)
	return viewPortState
}

func (v *Viewport) UpdateViewPortImage(st *state.AppState) {
	go func() {
		for image := range st.CanvasState.GetChannel() {
			if image == nil {
				fyne.Do(func() {
					v.Clear()
				})
			} else {
				fyne.Do(func() {
					v.imageView.SetImageInCanvas(image)
					if v.placeholder.Visible() {
						v.placeholder.Hide()
					}
				})
			}

		}
	}()
}

func (v *Viewport) Clear() {
	v.imageView.ClearCanvas()
	v.placeholder.Show()
}
