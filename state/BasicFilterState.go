package state

import (
	"math"

	"fyne.io/fyne/v2/data/binding"
)

type BasicFilterState struct {
	Blur       binding.Float
	Emboss     binding.Float
	Outline    binding.Float
	Sharpening binding.Float
	Sobel      binding.Float
}

func (bfs *BasicFilterState) SetBlur(blur float64) {
	bfs.Blur.Set(math.Max(0, math.Min(100, blur)))
}
func (bfs *BasicFilterState) SetEmboss(emboss float64) {
	bfs.Emboss.Set(math.Max(0, math.Min(100, emboss)))
}
func (bfs *BasicFilterState) SetOutline(outline float64) {
	bfs.Outline.Set(math.Max(0, math.Min(100, outline)))
}
func (bfs *BasicFilterState) SetSharpening(sharpening float64) {
	bfs.Sharpening.Set(math.Max(0, math.Min(100, sharpening)))
}
func (bfs *BasicFilterState) SetSobel(sobel float64) {
	bfs.Sobel.Set(math.Max(0, math.Min(100, sobel)))
}
func (bfs *BasicFilterState) InitBasicFilterState() {
	bfs.Blur.Set(50)
	bfs.Emboss.Set(50)
	bfs.Outline.Set(50)
	bfs.Sharpening.Set(50)
	bfs.Sobel.Set(50)
}
