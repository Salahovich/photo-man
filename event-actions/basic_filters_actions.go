package event_actions

import (
	"image"
	"photo-man/core/image_filters"
	"photo-man/state"
)

func UpdateBaseFiltersAction(img image.Image, bfs *state.BasicFilterState) image.Image {

	if blur := bfs.Blur; blur != 0 {
		img = image_filters.GaussianBlur(img, blur)
	}
	if emboss := bfs.Emboss; emboss != 0 {
		img = image_filters.EmbossImage(img, emboss)
	}
	if outline := bfs.Outline; outline != 0 {
		img = image_filters.OutlineImage(img, outline)
	}
	if sharpening := bfs.Sharpening; sharpening != 0 {
		img = image_filters.SharpImage(img, sharpening)
	}
	if sobel := bfs.Sobel; sobel != 0 {
		img = image_filters.SobelImage(img, sobel)
	}

	return img
}
