package colorBlending

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type BLEND_MODE int

const (
	STANDARD BLEND_MODE = 1
	MULTIPLY BLEND_MODE = 2
	SCREEN   BLEND_MODE = 3
	OVERLAY  BLEND_MODE = 4
)

func PerformBlending(img image.Image, blendColor color.Color, blendMode func(color.Color, color.Color) color.RGBA64) image.Image {
	newImg := image.NewRGBA64(img.Bounds())
	wg := sync.WaitGroup{}
	wg.Add(img.Bounds().Dy())
	for y := 0; y < img.Bounds().Dy(); y++ {
		go func(y int) {
			defer wg.Done()
			for x := 0; x < img.Bounds().Dx(); x++ {
				newColor := blendMode(img.At(x, y), blendColor)
				newImg.SetRGBA64(x, y, newColor)
			}
		}(y)
	}
	wg.Wait()
	return newImg
}

func StandardBlend(pixel color.Color, blend color.Color) color.RGBA64 {
	r, g, b, a := pixel.RGBA()
	rBlend, gBlend, bBlend, aBlend := blend.RGBA()

	// normalize to [0,1]
	r_, g_, b_, a_ := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)
	rBlend_, gBlend_, bBlend_, aBlend_ := float64(rBlend)/65535.0, float64(gBlend)/65535.0, float64(bBlend)/65535.0, float64(aBlend)/65535.0

	// color standard blending per channel
	r_ = math.Max(0, math.Min(65535, (r_*(1-aBlend_)+(rBlend_*aBlend_))*65535.0))
	g_ = math.Max(0, math.Min(65535, (g_*(1-aBlend_)+(gBlend_*aBlend_))*65535.0))
	b_ = math.Max(0, math.Min(65535, (b_*(1-aBlend_)+(bBlend_*aBlend_))*65535.0))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}

func MultiplyBlend(pixel color.Color, blend color.Color) color.RGBA64 {
	r, g, b, a := pixel.RGBA()
	rBlend, gBlend, bBlend, aBlend := blend.RGBA()

	// normalize to [0,1]
	r_, g_, b_, a_ := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)
	rBlend_, gBlend_, bBlend_, aBlend_ := float64(rBlend)/65535.0, float64(gBlend)/65535.0, float64(bBlend)/65535.0, float64(aBlend)/65535.0

	// color multiply blending per channel
	rBlend_ = r_ * rBlend_
	gBlend_ = g_ * gBlend_
	bBlend_ = b_ * bBlend_

	// transparency handling
	r_ = math.Max(0, math.Min(65535, (r_*(1-aBlend_)+(rBlend_*aBlend_))*65535.0))
	g_ = math.Max(0, math.Min(65535, (g_*(1-aBlend_)+(gBlend_*aBlend_))*65535.0))
	b_ = math.Max(0, math.Min(65535, (b_*(1-aBlend_)+(bBlend_*aBlend_))*65535.0))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}

func ScreenBlend(pixel color.Color, blend color.Color) color.RGBA64 {
	r, g, b, a := pixel.RGBA()
	rBlend, gBlend, bBlend, aBlend := blend.RGBA()

	// normalize to [0,1]
	r_, g_, b_, a_ := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)
	rBlend_, gBlend_, bBlend_, aBlend_ := float64(rBlend)/65535.0, float64(gBlend)/65535.0, float64(bBlend)/65535.0, float64(aBlend)/65535.0

	// color screen blending per channel
	rBlend_ = 1 - ((1 - r_) * (1 - rBlend_))
	gBlend_ = 1 - ((1 - g_) * (1 - gBlend_))
	bBlend_ = 1 - ((1 - b_) * (1 - bBlend_))

	// transparency handling
	r_ = math.Max(0, math.Min(65535, (r_*(1-aBlend_)+(rBlend_*aBlend_))*65535.0))
	g_ = math.Max(0, math.Min(65535, (g_*(1-aBlend_)+(gBlend_*aBlend_))*65535.0))
	b_ = math.Max(0, math.Min(65535, (b_*(1-aBlend_)+(bBlend_*aBlend_))*65535.0))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}

func OverlayBlend(pixel color.Color, blend color.Color) color.RGBA64 {
	r, g, b, a := pixel.RGBA()
	rBlend, gBlend, bBlend, aBlend := blend.RGBA()

	// normalize to [0,1]
	r_, g_, b_, a_ := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)
	rBlend_, gBlend_, bBlend_, aBlend_ := float64(rBlend)/65535.0, float64(gBlend)/65535.0, float64(bBlend)/65535.0, float64(aBlend)/65535.0

	// color overlay blending per channel, if blend color > 0.5 use screen else use multiply
	if r_ >= 0.5 {
		rBlend_ = 1 - 2*(1-r_)*(1-rBlend_)
	} else {
		rBlend_ = 2 * r_ * rBlend_
	}
	if g_ >= 0.5 {
		gBlend_ = 1 - 2*(1-g_)*(1-gBlend_)
	} else {
		gBlend_ = 2 * g_ * gBlend_
	}
	if b_ >= 0.5 {
		bBlend_ = 1 - 2*(1-b_)*(1-bBlend_)
	} else {
		bBlend_ = 2 * b_ * bBlend_
	}

	// transparency handling
	r_ = math.Max(0, math.Min(65535, (r_*(1-aBlend_)+(rBlend_*aBlend_))*65535.0))
	g_ = math.Max(0, math.Min(65535, (g_*(1-aBlend_)+(gBlend_*aBlend_))*65535.0))
	b_ = math.Max(0, math.Min(65535, (b_*(1-aBlend_)+(bBlend_*aBlend_))*65535.0))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}

func ExclusionBlend(pixel color.Color, blend color.Color) color.RGBA64 {
	r, g, b, a := pixel.RGBA()
	rBlend, gBlend, bBlend, aBlend := blend.RGBA()

	// normalize to [0,1]
	r_, g_, b_, a_ := float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0, float64(a)
	rBlend_, gBlend_, bBlend_, aBlend_ := float64(rBlend)/65535.0, float64(gBlend)/65535.0, float64(bBlend)/65535.0, float64(aBlend)/65535.0

	// color screen blending per channel
	rBlend_ = 0.5 - 2*((r_-0.5)*(rBlend_-0.5))
	gBlend_ = 0.5 - 2*((g_-0.5)*(gBlend_-0.5))
	bBlend_ = 0.5 - 2*((b_-0.5)*(bBlend_-0.5))

	// transparency handling
	r_ = math.Max(0, math.Min(65535, (r_*(1-aBlend_)+(rBlend_*aBlend_))*65535.0))
	g_ = math.Max(0, math.Min(65535, (g_*(1-aBlend_)+(gBlend_*aBlend_))*65535.0))
	b_ = math.Max(0, math.Min(65535, (b_*(1-aBlend_)+(bBlend_*aBlend_))*65535.0))

	return color.RGBA64{R: uint16(r_), G: uint16(g_), B: uint16(b_), A: uint16(a_)}
}
