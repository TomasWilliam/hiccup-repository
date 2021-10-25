package internal

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func generateGif(positions [][]location) {
	var images []*image.Paletted
	var delays []int

	var palette = []color.Color{
		color.RGBA{R: 0, G: 0, B: 0, A: 255},       // black
		color.RGBA{R: 255, G: 255, B: 255, A: 255}, // white
	}

	for i := range positions {
		img := image.NewPaletted(image.Rect(0, 0, int(width), int(width)), palette)
		images = append(images, img)
		delays = append(delays, 1)
		for j := range positions[i] {
			img.Set(int(positions[i][j].x), int(positions[i][j].y), palette[1])
		}
	}

	f, _ := os.OpenFile("sim.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
