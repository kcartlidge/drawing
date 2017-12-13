package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	drawing "github.com/kcartlidge/drawing"
)

func main() {
	fmt.Println("surface example")

	const width, height = 250, 250
	bg := color.NRGBA{R: 0, G: 0, B: 0, A: 128}
	fg := color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	s := drawing.NewSurface(width, height, bg)

	s.DrawRect(s.Bounds, fg)

	f, err := os.Create("example.png")
	if err == nil {
		err = s.WritePNG(f)
		if err == nil {
			err = f.Close()
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
