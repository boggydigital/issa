package issa

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"strconv"
	"strings"
)

const (
	MIMETypeBase64Prefix = "data:image/gif;base64,"
	GIF89aBase64Header   = "R0lGODlh" // GIF89a
)

//logicalScreenImageDescriptors provides 7 byte logical screen descriptor,
// and 10 bytes image descriptor according to the https://www.w3.org/Graphics/GIF/spec-gif89a.txt
func logicalScreenImageDescriptors(width, height int) []byte {
	return []byte{
		// Logical Screen Descriptor
		uint8(width % 256),  // Logical Screen Width
		uint8(width / 256),  // Logical Screen Width
		uint8(height % 256), // Logical Screen Height
		uint8(height / 256), // Logical Screen Height
		0,                   // <Packed Fields>
		0,                   // Background Color Index
		0,                   // Pixel Aspect Ratio
		// Image Descriptor
		0x2C,                // Image Separator
		0,                   // Image Left Position
		0,                   // Image Left Position
		0,                   // Image Top Position
		0,                   // Image Top Position
		uint8(width % 256),  // Image Width
		uint8(width / 256),  // Image Width
		uint8(height % 256), // Image Height
		uint8(height / 256), // Image Height
		0x87,                // <Packed Fields>
	}
}

//formatDescriptorsStdPalette is a reproducible portion of base64 encoded GIF image
//with a standard palette that contains GIF89a format header, logical screen
//descriptor, image descriptor, and standard palette
func formatDescriptorsStdPalette(width, height int) string {
	return GIF89aBase64Header +
		base64.RawStdEncoding.EncodeToString(
			logicalScreenImageDescriptors(width, height)) +
		StdPaletteBase64Content
}

//Dehydrate encodes a GIF image (base64 encoding), removes a reproducible portion (about 1Kb of content) and
//prefixes the result with image width and height (separated with base64 padding character)
func Dehydrate(gifImage *gif.GIF) (string, error) {

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

	return fmt.Sprintf("%d%c%d%c", size.X, base64.StdPadding, size.Y, base64.StdPadding) +
		strings.TrimPrefix(b64s, formatDescriptorsStdPalette(size.X, size.Y)), nil
}

//Hydrate takes dehydrated string, restores it and prefixes with base64 GIF image MIME type
func Hydrate(dehydrated string) string {
	if dehydrated == "" {
		return ""
	}

	parts := strings.Split(dehydrated, "=")
	if len(parts) != 3 {
		return ""
	}

	if width, err := strconv.Atoi(parts[0]); err == nil {
		if height, err := strconv.Atoi(parts[1]); err == nil {
			return MIMETypeBase64Prefix +
				formatDescriptorsStdPalette(width, height) +
				parts[2]
		}
	}

	return ""
}
