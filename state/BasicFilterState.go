package state

import (
	"photo-man/core/image_filters"
)

type BasicFilterState struct {
	Blur       image_filters.BLUR_QUALITY
	Emboss     image_filters.EMBOSS_QUALITY
	Outline    image_filters.OUTLINE_QUALITY
	Sharpening image_filters.SHARPENING_QUALITY
	Sobel      image_filters.SOBEL_DIRECTION
}

func (bfs *BasicFilterState) SetBlurQuality(quality image_filters.BLUR_QUALITY) {
	// reset filters value to apply one filter at a time
	bfs.InitBasicFilterState()
	bfs.Blur = quality
}
func (bfs *BasicFilterState) SetEmbossQuality(quality image_filters.EMBOSS_QUALITY) {
	bfs.InitBasicFilterState()
	bfs.Emboss = quality
}
func (bfs *BasicFilterState) SetOutlineQuality(quality image_filters.OUTLINE_QUALITY) {
	bfs.InitBasicFilterState()
	bfs.Outline = quality
}
func (bfs *BasicFilterState) SetSharpeningQuality(quality image_filters.SHARPENING_QUALITY) {
	bfs.InitBasicFilterState()
	bfs.Sharpening = quality
}
func (bfs *BasicFilterState) SetSobelQuality(direction image_filters.SOBEL_DIRECTION) {
	bfs.InitBasicFilterState()
	bfs.Sobel = direction
}

func (bfs *BasicFilterState) InitBasicFilterState() {
	bfs.Blur = 0
	bfs.Emboss = 0
	bfs.Outline = 0
	bfs.Sharpening = 0
	bfs.Sobel = 0
}
