package main

import (
	"fmt"
	"github.com/boggydigital/issa"
	"image"
	_ "image/jpeg"
	"os"
	"time"
)

func main() {
	//findOptimalSettings()
	findDominantColor()
}

func findOptimalSettings() {
	//below is a sample of using OptimalSampling on your images data set

	ni := 6
	images := make([]string, ni)

	for i := range ni {
		images[i] = fmt.Sprintf("image-%d.jpeg", i+1)
	}

	samples, tlen, err := issa.OptimalSampling(images)
	if err != nil {
		panic(err)
	}

	fmt.Println(samples, tlen/ni)
}

func findDominantColor() {

	fo, err := os.Create("test.html")
	if err != nil {
		panic(err)
	}

	if _, err = fo.WriteString("<body style='font-family:sans-serif'>"); err != nil {
		panic(err)
	}

	start := time.Now()

	ni := 26

	for i := range ni {
		imagePath := fmt.Sprintf("image-%d.jpeg", i+1)

		imageFile, err := os.Open(imagePath)
		if err != nil {
			panic(err)
		}

		jpegImage, _, err := image.Decode(imageFile)
		if err != nil {
			panic(err)
		}

		gifImage := issa.GIFImage(jpegImage, issa.ColorPalette(), issa.DefaultSampling)

		dc := issa.RepColor(gifImage)

		//r, g, b, _ := dc.RGBA()
		rgb := issa.ColorHex(dc)

		if _, err = fmt.Fprintf(fo, "<div style='background:%s;text-align:center'>", rgb); err != nil {
			panic(err)
		}

		if _, err = fmt.Fprintf(fo, "<img src='image-%d.jpeg' style='height:100px;padding:16px'>", i+1); err != nil {
			panic(err)
		}

		if _, err = fmt.Fprintf(fo, "<div style='color:white'>%s</div><br/>", rgb); err != nil {
			panic(err)
		}

		if _, err = fo.WriteString("</div>"); err != nil {
			panic(err)
		}
	}

	if _, err = fo.WriteString("</body>"); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed.Milliseconds(), "ms")

}
