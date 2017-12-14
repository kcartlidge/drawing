package main

import (
	"fmt"
	"log"
	"os"

	drawing "github.com/kcartlidge/drawing"
)

func main() {
	fmt.Println()
	fmt.Println("Drawing Example")
	fmt.Println()

	// Drawing surface with a black background and yellow border.
	const width, height = 500, 500
	s := drawing.NewSurface(width, height, drawing.Black(255), drawing.LogDuration)
	b := drawing.NewRect(5, 5, width-5, height-5)
	s.DrawRect(b, drawing.Yellow(255))
	midX := (width / 2)

	// Draw some bars with sample colors (internally also uses Hline).
	s.FillRect(drawing.NewRect(10, 10, width-10, 12), drawing.Blue(255))
	s.FillRect(drawing.NewRect(10, 15, width-10, 17), drawing.Red(255))
	s.FillRect(drawing.NewRect(10, 20, width-10, 22), drawing.Magenta(255))
	s.FillRect(drawing.NewRect(10, 25, width-10, 27), drawing.Green(255))
	s.FillRect(drawing.NewRect(10, 30, width-10, 32), drawing.Cyan(255))
	s.FillRect(drawing.NewRect(10, 35, width-10, 37), drawing.Yellow(255))
	s.FillRect(drawing.NewRect(10, 40, width-10, 42), drawing.Gray(255))
	s.FillRect(drawing.NewRect(10, 45, width-10, 47), drawing.LightBlue(255))
	s.FillRect(drawing.NewRect(10, 50, width-10, 52), drawing.LightRed(255))
	s.FillRect(drawing.NewRect(10, 55, width-10, 57), drawing.LightMagenta(255))
	s.FillRect(drawing.NewRect(10, 60, width-10, 62), drawing.LightGreen(255))
	s.FillRect(drawing.NewRect(10, 65, width-10, 67), drawing.LightCyan(255))
	s.FillRect(drawing.NewRect(10, 70, width-10, 72), drawing.LightYellow(255))

	// Draw a vertical separator.
	s.Vline(drawing.NewPoint(midX, 85), drawing.NewPoint(midX, height-10), drawing.Blue(255))

	// Show some standard (fast) lines.
	s.Line(drawing.NewPoint(10, 85), drawing.NewPoint(midX-5, height-10), drawing.Cyan(255))
	s.Line(drawing.NewPoint(midX-5, 85), drawing.NewPoint(10, height-10), drawing.Green(255))

	// Show some anti-aliased lines.
	s.LineA(drawing.NewPoint(midX+5, 85), drawing.NewPoint(width-10, height-10), drawing.Cyan(255))
	s.LineA(drawing.NewPoint(width-10, 85), drawing.NewPoint(midX+5, height-10), drawing.Green(255))

	// Draw sample circles.
	s.Circle(midX/2, ((height-85)/2)+85, (midX/2)-10, drawing.LightRed(255))

	// Draw sample filled circle.
	s.FillCircle((midX/2)+midX, ((height-85)/2)+85, (midX/2)-10, drawing.Red(255))

	// Export.
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
