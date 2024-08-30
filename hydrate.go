package issa

import (
	"strconv"
	"strings"
)

const (
	MIMETypeBase64Prefix = "data:image/gif;base64,"
	GIF89aBase64Header   = "R0lGODlh" // GIF89a
)

// hydrate takes dehydrated string, restores it and prefixes with base64 GIF image MIME type
func hydrate(dehydrated string, palettePrefix func(w, h int) string) string {
	if dehydrated == "" {
		return ""
	}

	parts := strings.Split(dehydrated, "=")
	if len(parts) != 3 {
		return ""
	}

	if width, err := strconv.Atoi(parts[0]); err == nil {
		if height, err := strconv.Atoi(parts[1]); err == nil {

			hydratedImage := ""

			if strings.HasPrefix(parts[2], GIF89aBase64Header) {
				hydratedImage = MIMETypeBase64Prefix + parts[2]
			} else {
				hydratedImage = MIMETypeBase64Prefix + palettePrefix(width, height) +
					parts[2]
			}

			return hydratedImage
		}
	}

	return ""
}

func HydrateColor(dehydrated string) string {
	return hydrate(dehydrated, ColorPalettePrefix)
}

func HydrateGreyscale(dehydrated string) string {
	return hydrate(dehydrated, GreyscalePalettePrefix)
}
