package issa

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"strings"
)

// dehydrate encodes a GIF image (base64 encoding), removes a reproducible portion (about 1Kb of content) and
// prefixes the result with image width and height (separated with base64 padding character)
func dehydrate(gifImage *gif.GIF, palettePrefix func(w, h int) string) (string, error) {

	if len(gifImage.Image) < 1 {
		return "", fmt.Errorf("GIF must contain at least 1 frame")
	}

	fi := gifImage.Image[0]
	size := fi.Bounds().Size()

	bts := make([]byte, 0, size.X*size.Y)
	buf := bytes.NewBuffer(bts)

	if err := gif.EncodeAll(buf, gifImage); err != nil {
		return "", err
	}

	b64s := base64.RawStdEncoding.EncodeToString(buf.Bytes())

	return DehydratedSizePrefix(size.X, size.Y) + strings.TrimPrefix(b64s, palettePrefix(size.X, size.Y)), nil
}

func DehydrateColor(gifImage *gif.GIF) (string, error) {
	return dehydrate(gifImage, ColorPalettePrefix)
}

func DehydrateGreyscale(gifImage *gif.GIF) (string, error) {
	return dehydrate(gifImage, GreyscalePalettePrefix)
}

func DehydratedSizePrefix(x, y int) string {
	return fmt.Sprintf("%d%c%d%c", x, base64.StdPadding, y, base64.StdPadding)
}
