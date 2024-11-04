package issa

import (
	"encoding/hex"
	"image/color"
	"image/gif"
)

const minLuma = 192 << 8

func RepColor(image *gif.GIF) color.Color {

	if image == nil {
		return color.Black
	}

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
	if c == nil {
		return "#000000"
	}
	bts := make([]byte, 3)
	r, g, b, _ := c.RGBA()
	bts[0], bts[1], bts[2] = byte(r>>8), byte(g>>8), byte(b>>8)
	return "#" + hex.EncodeToString(bts)
}
