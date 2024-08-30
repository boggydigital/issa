package main

import (
	"fmt"
	"github.com/boggydigital/issa"
	_ "image/jpeg"
)

func main() {
	findOptimalSettings()
}

func findOptimalSettings() {
	//below is a sample of using FindOptimalSampling on your images data set

	ni := 6
	images := make([]string, ni)

	for i := 0; i < ni; i++ {
		images[i] = fmt.Sprintf("image-%d.jpeg", i+1)
	}

	samples, tlen, err := issa.FindOptimalSampling(images)
	if err != nil {
		panic(err)
	}

	fmt.Println(samples, tlen/ni)
}
