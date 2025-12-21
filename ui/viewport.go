package ui

import (
	"photo-man/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ViewPortContainer() *fyne.Container {
	vp := NewViewport()
	return vp.Container
}

type Viewport struct {
	Container   *fyne.Container // The main UI object to add to the window
	imageView   *canvas.Image   // The actual photo
	placeholder *fyne.Container // The "Open Photo" message
}

func NewViewport() *Viewport {
	// 1. Create the Actual Image holder (Initially empty/nil)
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain

	// 2. Create the Placeholder (The "Empty State")
	// We use a VBox to stack the Icon above the Text
	noPhotoIcon := canvas.NewImageFromResource(assets.NoPhoto)
	noPhotoIcon.FillMode = canvas.ImageFillContain
	noPhotoIcon.SetMinSize(fyne.NewSize(200, 200))
	placeholderContent := container.NewVBox(
		noPhotoIcon,
		widget.NewLabelWithStyle("Click 'Open' to start editing", fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	// Center the placeholder in the middle of the screen
	placeholderWrapper := container.NewCenter(placeholderContent)

	containerWrapper := container.NewStack(placeholderWrapper, img)

	return &Viewport{
		Container:   containerWrapper,
		imageView:   img,
		placeholder: placeholderWrapper,
	}
}

// SetImage updates the viewport with a new photo
func (v *Viewport) SetImage(res fyne.Resource) {
	// Update the image source
	v.imageView.Resource = res
	v.imageView.Refresh()

	// Hide the placeholder text so only the photo is visible
	v.placeholder.Hide()
}

// Clear removes the photo and shows the placeholder again
func (v *Viewport) Clear() {
	v.imageView.Resource = nil
	v.imageView.Refresh()
	v.placeholder.Show()
}
