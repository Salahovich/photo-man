package ui

import (
	"fmt"
	"image/color"
	event_actions "photo-man/event-actions"
	"photo-man/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Sidebar(st *state.AppState) *fyne.Container {

	// adjustment accordion item
	brightnessContainer, brightnessSlider := initBrightnessArea(st)
	contrastContainer, contrastSlider := initContrastArea(st)
	saturationContainer, saturationSlider := initSaturationArea(st)

	adjustmentsContainer := container.NewVBox(
		brightnessContainer,
		brightnessSlider,
		contrastContainer,
		contrastSlider,
		saturationContainer,
		saturationSlider)
	adjustments := widget.NewAccordionItem("Adjustments", adjustmentsContainer)

	// basic filter accordion item
	blurContainer, blurSlider := initBlurArea(st)
	embossContainer, embossSlider := initEmbossArea(st)
	outlineContainer, outlineSlider := initOutlineArea(st)
	sharpeningContainer, sharpeningSlider := initSharpeningArea(st)
	sobelContainer, sobelSlider := initSobelArea(st)

	basicFiltersContainer := container.NewVBox(
		blurContainer,
		blurSlider,
		embossContainer,
		embossSlider,
		outlineContainer,
		outlineSlider,
		sharpeningContainer,
		sharpeningSlider,
		sobelContainer,
		sobelSlider)
	basicFilters := widget.NewAccordionItem("Basic Filters", basicFiltersContainer)

	// predefined-filters accordion item
	predefinedFiltersContainer := initFiltersArea()
	predefinedFilters := widget.NewAccordionItem("Predefined Filters", predefinedFiltersContainer)

	// metadata accordion item
	metadataContainer := initMetadataArea()
	metadata := widget.NewAccordionItem("Metadata", metadataContainer)

	separator := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	separator.SetMinSize(fyne.NewSize(1, 1))

	sideTitle := widget.NewLabelWithStyle("Tools", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	accordion := widget.NewAccordion(adjustments, basicFilters, predefinedFilters, metadata)
	accordion.MultiOpen = false

	contentVBox := container.NewVBox(sideTitle, separator, accordion)
	bgColor := color.RGBA{R: 62, G: 62, B: 62, A: 255}
	background := canvas.NewRectangle(bgColor)

	return container.NewStack(background, container.NewPadded(contentVBox))
}

func initBrightnessArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	brightnessText := canvas.NewText("  Brightness", color.White)
	brightnessValue := canvas.NewText("50  ", color.White)
	brightnessContainer := container.NewBorder(nil, nil, brightnessText, brightnessValue, nil)
	brightnessSlider := widget.NewSliderWithData(0, 100, st.AdjustmentState.Brightness)
	brightnessSlider.SetValue(50)
	brightnessSlider.OnChanged = func(value float64) {
		brightnessValue.Text = fmt.Sprintf("%d ", int(value))
		st.AdjustmentState.SetBrightness(value)
		brightnessValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return brightnessContainer, brightnessSlider
}

func initContrastArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	contrastText := canvas.NewText("  Contrast", color.White)
	contrastValue := canvas.NewText("50  ", color.White)
	contrastContainer := container.NewBorder(nil, nil, contrastText, contrastValue, nil)
	contrastSlider := widget.NewSliderWithData(0, 100, st.AdjustmentState.Contrast)
	contrastSlider.SetValue(50)
	contrastSlider.OnChanged = func(value float64) {
		contrastValue.Text = fmt.Sprintf("%d ", int(value))
		st.AdjustmentState.SetContrast(value)
		contrastValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}
	return contrastContainer, contrastSlider
}

func initSaturationArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	saturationText := canvas.NewText("  Saturation", color.White)
	saturationValue := canvas.NewText("50  ", color.White)
	saturationContainer := container.NewBorder(nil, nil, saturationText, saturationValue, nil)
	saturationSlider := widget.NewSliderWithData(0, 100, st.AdjustmentState.Saturation)
	saturationSlider.SetValue(50)
	saturationSlider.OnChanged = func(value float64) {
		saturationValue.Text = fmt.Sprintf("%d ", int(value))
		st.AdjustmentState.SetSaturation(value)
		saturationValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return saturationContainer, saturationSlider
}

func initBlurArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	blurText := canvas.NewText("  Blur", color.White)
	blurValue := canvas.NewText("0  ", color.White)
	blurContainer := container.NewBorder(nil, nil, blurText, blurValue, nil)
	blurSlider := widget.NewSliderWithData(0, 100, st.BasicFilterState.Blur)
	blurSlider.SetValue(0)
	blurSlider.OnChanged = func(value float64) {
		blurValue.Text = fmt.Sprintf("%d ", int(value))
		st.BasicFilterState.SetBlur(value)
		blurValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return blurContainer, blurSlider
}

func initEmbossArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	embossText := canvas.NewText("  Emboss", color.White)
	embossValue := canvas.NewText("0  ", color.White)
	embossContainer := container.NewBorder(nil, nil, embossText, embossValue, nil)
	embossSlider := widget.NewSliderWithData(0, 100, st.BasicFilterState.Emboss)
	embossSlider.SetValue(0)
	embossSlider.OnChanged = func(value float64) {
		embossValue.Text = fmt.Sprintf("%d ", int(value))
		st.BasicFilterState.SetEmboss(value)
		embossValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return embossContainer, embossSlider
}

func initOutlineArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	outlineText := canvas.NewText("  Outline", color.White)
	outlineValue := canvas.NewText("0  ", color.White)
	outlineContainer := container.NewBorder(nil, nil, outlineText, outlineValue, nil)
	outlineSlider := widget.NewSliderWithData(0, 100, st.BasicFilterState.Outline)
	outlineSlider.SetValue(0)
	outlineSlider.OnChanged = func(value float64) {
		outlineValue.Text = fmt.Sprintf("%d ", int(value))
		st.BasicFilterState.SetOutline(value)
		outlineValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return outlineContainer, outlineSlider
}

func initSharpeningArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	sharpeningText := canvas.NewText("  Sharpening", color.White)
	sharpeningValue := canvas.NewText("0  ", color.White)
	sharpeningContainer := container.NewBorder(nil, nil, sharpeningText, sharpeningValue, nil)
	sharpeningSlider := widget.NewSliderWithData(0, 100, st.BasicFilterState.Sharpening)
	sharpeningSlider.SetValue(0)
	sharpeningSlider.OnChanged = func(value float64) {
		sharpeningValue.Text = fmt.Sprintf("%d ", int(value))
		st.BasicFilterState.SetSharpening(value)
		sharpeningValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return sharpeningContainer, sharpeningSlider
}

func initSobelArea(st *state.AppState) (*fyne.Container, *widget.Slider) {
	sobelText := canvas.NewText("  Sobel", color.White)
	sobelValue := canvas.NewText("0  ", color.White)
	sobelContainer := container.NewBorder(nil, nil, sobelText, sobelValue, nil)
	sobelSlider := widget.NewSliderWithData(0, 100, st.BasicFilterState.Sobel)
	sobelSlider.SetValue(0)
	sobelSlider.OnChanged = func(value float64) {
		sobelValue.Text = fmt.Sprintf("%d ", int(value))
		st.BasicFilterState.SetSobel(value)
		sobelValue.Refresh()
		go event_actions.UpdateAdjustments(st)
	}

	return sobelContainer, sobelSlider
}

func initFiltersArea() *fyne.Container {
	rec1 := canvas.NewRectangle(color.White)
	rec1.SetMinSize(fyne.NewSize(50, 50))
	rec2 := canvas.NewRectangle(color.White)
	rec2.SetMinSize(fyne.NewSize(50, 50))
	rec3 := canvas.NewRectangle(color.White)
	rec3.SetMinSize(fyne.NewSize(50, 50))
	rec4 := canvas.NewRectangle(color.White)
	rec4.SetMinSize(fyne.NewSize(50, 50))

	return container.NewGridWithColumns(2, rec1, rec2, rec3, rec4)
}

func initMetadataArea() *fyne.Container {
	fileType := canvas.NewText(".png", color.White)
	metaSeparatorOne := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	metaSeparatorOne.SetMinSize(fyne.NewSize(1, 1))
	fileSize := canvas.NewText("1.2MB", color.White)
	metaSeparatorTwo := canvas.NewRectangle(color.RGBA{R: 74, G: 85, B: 101, A: 255})
	metaSeparatorTwo.SetMinSize(fyne.NewSize(1, 1))
	imageDimensions := canvas.NewText("1920x1080", color.White)

	return container.NewCenter(
		container.NewHBox(
			fileType,
			metaSeparatorOne,
			fileSize,
			metaSeparatorTwo,
			imageDimensions))
}
