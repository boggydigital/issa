package issa

import (
	"encoding/hex"
	"image/color"
	"image/gif"
)

const minLuma = 192 << 8

func RepColor(image *gif.GIF) color.Color {

	colors := make(map[color.Color]int)

	for _, frame := range image.Image {
		for _, p := range frame.Pix {
			c := frame.Palette[p]
			if r, g, b, _ := c.RGBA(); (r + g + b) < minLuma {
				continue
			}
			colors[c] += 1
		}
	}

	maxCount := 0
	var maxColor color.Color

	for c, n := range colors {
		if n > maxCount {
			maxCount = n
			maxColor = c
		}
	}

	return maxColor
}

func ColorHex(c color.Color) string {
	bts := make([]byte, 0, 3)
	r, g, b, _ := c.RGBA()
	bts = append(bts, byte(r>>8), byte(g>>8), byte(b>>8))
	return "#" + hex.EncodeToString(bts)
}
