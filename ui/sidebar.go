package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Sidebar() *fyne.Container {

	// adjustment accordion item
	brightnessText := canvas.NewText("  Brightness", color.White)
	brightnessValue := canvas.NewText("0  ", color.White)
	brightnessContainer := container.NewBorder(nil, nil, brightnessText, brightnessValue, nil)
	brightnessSlider := widget.NewSlider(0, 100)
	brightnessSlider.OnChanged = func(value float64) {
		brightnessValue.Text = fmt.Sprintf("%d  ", int(value))
		brightnessValue.Refresh()
	}

	contrastText := canvas.NewText("  Contrast", color.White)
	contrastValue := canvas.NewText("0  ", color.White)
	contrastContainer := container.NewBorder(nil, nil, contrastText, contrastValue, nil)
	contrastSlider := widget.NewSlider(0, 100)
	contrastSlider.OnChanged = func(value float64) {
		contrastValue.Text = fmt.Sprintf("%d  ", int(value))
		contrastValue.Refresh()
	}

	saturationText := canvas.NewText("  Saturation", color.White)
	saturationValue := canvas.NewText("0  ", color.White)
	saturationContainer := container.NewBorder(nil, nil, saturationText, saturationValue, nil)
	saturationSlider := widget.NewSlider(0, 100)
	saturationSlider.OnChanged = func(value float64) {
		saturationValue.Text = fmt.Sprintf("%d  ", int(value))
		saturationValue.Refresh()
	}

	adjustmentsContainer := container.NewVBox(
		brightnessContainer,
		brightnessSlider,
		contrastContainer,
		contrastSlider,
		saturationContainer,
		saturationSlider)
	adjustments := widget.NewAccordionItem("Adjustments", adjustmentsContainer)

	// filter accordion item
	rec1 := canvas.NewRectangle(color.White)
	rec1.SetMinSize(fyne.NewSize(50, 50))
	rec2 := canvas.NewRectangle(color.White)
	rec2.SetMinSize(fyne.NewSize(50, 50))
	rec3 := canvas.NewRectangle(color.White)
	rec3.SetMinSize(fyne.NewSize(50, 50))
	rec4 := canvas.NewRectangle(color.White)
	rec4.SetMinSize(fyne.NewSize(50, 50))

	filtersContainer := container.NewGridWithColumns(2, rec1, rec2, rec3, rec4)
	filters := widget.NewAccordionItem("Filters", filtersContainer)

	// metadata accordion item
	fileType := canvas.NewText(".png", color.White)
	metaSeparatorOne := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	metaSeparatorOne.SetMinSize(fyne.NewSize(1, 1))
	fileSize := canvas.NewText("1.2MB", color.White)
	metaSeparatorTwo := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	metaSeparatorTwo.SetMinSize(fyne.NewSize(1, 1))
	imageDimensions := canvas.NewText("1920x1080", color.White)

	metadataContainer := container.NewCenter(container.NewHBox(fileType, metaSeparatorOne, fileSize, metaSeparatorTwo, imageDimensions))
	metadata := widget.NewAccordionItem("Metadata", metadataContainer)

	separator := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	separator.SetMinSize(fyne.NewSize(1, 1))

	sideTitle := widget.NewLabelWithStyle("Tools", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	accordion := widget.NewAccordion(adjustments, filters, metadata)
	accordion.MultiOpen = true

	contentVBox := container.NewVBox(sideTitle, separator, accordion)
	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	return container.NewStack(background, container.NewPadded(contentVBox))
}
