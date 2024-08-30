package issa

import "encoding/base64"

// logicalScreenImageDescriptors provides 7 byte logical screen descriptor,
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

// palettePrefix is a reproducible portion of base64 encoded GIF image
// with a standard palette that contains GIF89a format header, logical screen
// descriptor, image descriptor, and standard palette (color or greyscale)
func palettePrefix(w, h int, paletteBase64Content string) string {
	return GIF89aBase64Header +
		base64.RawStdEncoding.EncodeToString(
			logicalScreenImageDescriptors(w, h)) +
		paletteBase64Content
}

func ColorPalettePrefix(w, h int) string {
	return palettePrefix(w, h, ColorPaletteBase64Content)
}

func GreyscalePalettePrefix(w, h int) string {
	return palettePrefix(w, h, GreyscalePaletteBase64Content)
}
