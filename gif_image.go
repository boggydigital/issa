package issa

import (
	"image"
	"image/color"
	"image/gif"
)

const DefaultSampling = 16

//GIFImage converts a given image to a standard palette (see StdPalette) image,
//that is sampled every `sample` pixels (e.g. every 16 pixels by default)
func GIFImage(img image.Image, plt color.Palette, sample int) *gif.GIF {

	size := img.Bounds().Size()
	gifWidth, gifHeight := size.X/sample, size.Y/sample
	gifRect := image.Rect(0, 0, gifWidth, gifHeight)
	offset := sample / 2

	gifImage := image.NewPaletted(gifRect, plt)

	for y := 0; y < gifHeight; y++ {
		for x := 0; x < gifWidth; x++ {
			pltc := uint8(plt.Index(img.At(x*sample+offset, y*sample+offset)))
			gifImage.SetColorIndex(x, y, pltc)
		}
	}

	return &gif.GIF{Delay: []int{0}, Image: []*image.Paletted{gifImage}}
}
