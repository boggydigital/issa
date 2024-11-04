package issa

import (
	"image"
	"math"
	"os"
)

const maxSample = 64

func OptimalSampling(imageSources []string) (int, int, error) {

	images := make([]image.Image, len(imageSources))
	for i, imgSrc := range imageSources {

		fi, err := os.Open(imgSrc)
		if err != nil {
			return -1, 0, err
		}

		di, _, err := image.Decode(fi)
		if err != nil {
			return -1, 0, err
		}

		images[i] = di
	}

	minLen := math.MaxInt
	minSample := -1
	totalLen := 0

	for s := 2; s < maxSample; s++ {

		totalLen = 0
		for _, img := range images {

			ig := GIFImage(img, ColorPalette(), s)

			dhi, err := DehydrateColor(ig)
			if err != nil {
				return minSample, 0, err
			}

			totalLen += len(dhi)
		}

		if totalLen < minLen {
			minLen = totalLen
			minSample = s
		}

	}

	return minSample, totalLen, nil
}
