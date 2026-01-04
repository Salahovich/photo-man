package ui

import (
	"fmt"
	"image/color"
	"photo-man/core/image_filters"
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
	blurContainer := initBlurArea(st)
	embossContainer := initEmbossArea(st)
	outlineContainer := initOutlineArea(st)
	sharpeningContainer := initSharpeningArea(st)
	sobelContainer := initSobelArea(st)

	basicFiltersContainer := container.NewVBox(
		blurContainer,
		widget.NewSeparator(),
		embossContainer,
		widget.NewSeparator(),
		outlineContainer,
		widget.NewSeparator(),
		sharpeningContainer,
		widget.NewSeparator(),
		sobelContainer)

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
		go event_actions.PerformEdit(st)
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
		go event_actions.PerformEdit(st)
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
		go event_actions.PerformEdit(st)
	}

	return saturationContainer, saturationSlider
}

func initBlurArea(st *state.AppState) *fyne.Container {
	blurText := canvas.NewText("  Blur", color.White)
	lowButton := widget.NewButton("LOW", func() {
		st.BasicFilterState.SetBlurQuality(image_filters.LOW_BLUR)
		go event_actions.PerformEdit(st)
	})
	medButton := widget.NewButton("MED", func() {
		st.BasicFilterState.SetBlurQuality(image_filters.MEDIUM_BLUR)
		go event_actions.PerformEdit(st)
	})
	highButton := widget.NewButton("HIGH", func() {
		st.BasicFilterState.SetBlurQuality(image_filters.HIGH_BLUR)
		go event_actions.PerformEdit(st)
	})

	container := container.NewBorder(nil, nil, blurText, container.NewHBox(lowButton, medButton, highButton))
	return container
}

func initEmbossArea(st *state.AppState) *fyne.Container {
	embossText := canvas.NewText("  Emboss", color.White)
	lightButton := widget.NewButton("LIGHT", func() {
		st.BasicFilterState.SetEmbossQuality(image_filters.LIGHT_EMBOSS)
		go event_actions.PerformEdit(st)
	})
	darkButton := widget.NewButton("DARK", func() {
		st.BasicFilterState.SetEmbossQuality(image_filters.DARK_EMBOSS)
		go event_actions.PerformEdit(st)
	})
	heavyButton := widget.NewButton("HEAVY", func() {
		st.BasicFilterState.SetEmbossQuality(image_filters.HEAVY_EMBOSS)
		go event_actions.PerformEdit(st)
	})
	container := container.NewBorder(nil, nil, embossText, container.NewHBox(lightButton, heavyButton, darkButton))
	return container
}

func initOutlineArea(st *state.AppState) *fyne.Container {
	outlineText := canvas.NewText("  Outline", color.White)
	standardButton := widget.NewButton("STANDARD", func() {
		st.BasicFilterState.SetOutlineQuality(image_filters.STANDARD_OUTLINE)
		go event_actions.PerformEdit(st)
	})

	container := container.NewBorder(nil, nil, outlineText, container.NewHBox(standardButton))
	return container
}

func initSharpeningArea(st *state.AppState) *fyne.Container {
	blurText := canvas.NewText("  Sharpening", color.White)
	lowButton := widget.NewButton("LOW", func() {
		st.BasicFilterState.SetSharpeningQuality(image_filters.LOW_SHARP)
		go event_actions.PerformEdit(st)
	})
	medButton := widget.NewButton("MED", func() {
		st.BasicFilterState.SetSharpeningQuality(image_filters.MEDIUM_SHARP)
		go event_actions.PerformEdit(st)
	})
	highButton := widget.NewButton("HIGH", func() {
		st.BasicFilterState.SetSharpeningQuality(image_filters.HIGH_SHARP)
		go event_actions.PerformEdit(st)
	})

	container := container.NewBorder(nil, nil, blurText, container.NewHBox(lowButton, medButton, highButton))
	return container
}

func initSobelArea(st *state.AppState) *fyne.Container {
	blurText := canvas.NewText("  Sobel", color.White)
	leftButton := widget.NewButton("LEFT", func() {
		st.BasicFilterState.SetSobelQuality(image_filters.LEFT_SOBEL)
		go event_actions.PerformEdit(st)
	})
	topButton := widget.NewButton("TOP", func() {
		st.BasicFilterState.SetSobelQuality(image_filters.TOP_SOBEL)
		go event_actions.PerformEdit(st)
	})
	rightButton := widget.NewButton("RIGHT", func() {
		st.BasicFilterState.SetSobelQuality(image_filters.RIGHT_SOBEL)
		go event_actions.PerformEdit(st)
	})
	bottomButton := widget.NewButton("BOTTOM", func() {
		st.BasicFilterState.SetSobelQuality(image_filters.BOTTOM_SOBEL)
		go event_actions.PerformEdit(st)
	})
	container := container.NewBorder(nil, nil, blurText, container.NewHBox(leftButton, topButton, rightButton, bottomButton))
	return container
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
