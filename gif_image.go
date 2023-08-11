package issa

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/gif"
)

const DefaultDownSampling = 16

// GIFImage converts a given image to a standard palette (see StdPalette) image,
// that is down-sampled (e.g. by a factor of 16x by default)
func GIFImage(img image.Image, plt color.Palette, sample int) *gif.GIF {

	size := img.Bounds().Size()
	gifRect := image.Rect(0, 0, size.X/sample, size.Y/sample)

	gifImage := image.NewPaletted(gifRect, plt)

	draw.CatmullRom.Scale(gifImage, gifRect, img, img.Bounds(), draw.Src, nil)

	return &gif.GIF{Delay: []int{0}, Image: []*image.Paletted{gifImage}}
}
