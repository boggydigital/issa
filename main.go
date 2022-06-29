package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	_ "image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Palette() color.Palette {
	plt := color.Palette{}

	n := uint8(6)
	m := uint8(255) / n // 42

	//adding the first 6 * 6 * 6 = 216 colors in 42 increments
	//from (42,42,42) to (252,252,252)
	for r := uint8(0); r < n; r++ {
		for g := uint8(0); g < n; g++ {
			for b := uint8(0); b < n; b++ {
				c := color.RGBA{
					R: (r + 1) * m,
					G: (g + 1) * m,
					B: (b + 1) * m,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	//adding 8 more colors from (0,0,0) to (255,255,255)
	//representing combinations of 0 and 255
	for r := uint8(0); r <= 1; r++ {
		for g := uint8(0); g <= 1; g++ {
			for b := uint8(0); b <= 1; b++ {
				c := color.RGBA{
					R: uint8(255) * r,
					G: uint8(255) * g,
					B: uint8(255) * b,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	start, step := 255-m/2, m

	for r := start; r >= start-2*step; r -= step {
		for g := start; g >= start-2*step; g -= step {
			for b := start; b >= start-2*step; b -= step {
				c := color.RGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	return plt
}

func debugPaletteHTML(w io.Writer, plt color.Palette) {
	io.WriteString(w, "<style>")
	io.WriteString(w, "div { width: 1em; height: 1em; display: inline-block }\n")
	for i, c := range plt {
		r, g, b, _ := c.RGBA()
		io.WriteString(w, fmt.Sprintf(".i%d { background-color: rgb(%d,%d,%d) }\n", i+1, r>>8, g>>8, b>>8))
	}
	io.WriteString(w, "</style>")

	for i := 0; i < len(plt); i++ {
		io.WriteString(w, fmt.Sprintf("<div class='i%d'></div>", i+1))
	}
}

func PalettedGIF(img image.Image, plt color.Palette, sample int) *gif.GIF {

	size := img.Bounds().Size()
	gifWidth, gifHeight := size.X/sample, size.Y/sample
	gifRect := image.Rect(0, 0, gifWidth, gifHeight)

	gifImage := image.NewPaletted(gifRect, plt)

	for y := 0; y < gifHeight; y++ {
		for x := 0; x < gifWidth; x++ {
			if y > size.Y ||
				x > size.X {
				continue
			}
			c := img.At(x*sample, y*sample)
			gifImage.SetColorIndex(x, y, uint8(plt.Index(c)))
		}
	}

	return &gif.GIF{Delay: []int{0}, Image: []*image.Paletted{gifImage}}
}

func main() {

	start := time.Now()

	plt := Palette()

	dbgHtml, err := os.Create("palette.html")
	if err != nil {
		panic(err)
	}

	defer dbgHtml.Close()

	debugPaletteHTML(dbgHtml, plt)

	inputs := []string{
		"input1.jpeg",
		"input2.jpeg",
		"input3.jpeg",
		"input4.jpeg",
		"input5.jpeg",
		"input6.jpeg",
	}

	for _, input := range inputs {

		f, err := os.Open(input)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}

		gifImage := PalettedGIF(img, plt, 20)

		fn := strings.TrimSuffix(input, filepath.Ext(input)) + ".gif"
		gout, err := os.Create(fn)
		if err != nil {
			panic(err)
		}

		defer gout.Close()

		if err := gif.EncodeAll(gout, gifImage); err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
