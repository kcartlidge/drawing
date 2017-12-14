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

	// Drawing surface with a black background and yellow inset border.
	const width, height = 500, 500
	s := drawing.NewSurface(width, height, drawing.LogDuration)
	s.Clear(drawing.Black)
	b := drawing.NewRect(5, 5, width-5, height-5)
	s.DrawRect(b, drawing.Yellow)
	midX := (width / 2)

	// Get the background color from the drawn inset border.
	_ = s.GetPoint(5, 5)

	// Draw some bars with sample colors (internally also uses Hline).
	s.FillRect(drawing.NewRect(10, 10, width-10, 12), drawing.Blue)
	s.FillRect(drawing.NewRect(10, 15, width-10, 17), drawing.Red)
	s.FillRect(drawing.NewRect(10, 20, width-10, 22), drawing.Magenta)
	s.FillRect(drawing.NewRect(10, 25, width-10, 27), drawing.Green)
	s.FillRect(drawing.NewRect(10, 30, width-10, 32), drawing.Cyan)
	s.FillRect(drawing.NewRect(10, 35, width-10, 37), drawing.Yellow)
	s.FillRect(drawing.NewRect(10, 40, width-10, 42), drawing.Gray)
	s.FillRect(drawing.NewRect(10, 45, width-10, 47), drawing.LightBlue)
	s.FillRect(drawing.NewRect(10, 50, width-10, 52), drawing.LightRed)
	s.FillRect(drawing.NewRect(10, 55, width-10, 57), drawing.LightMagenta)
	s.FillRect(drawing.NewRect(10, 60, width-10, 62), drawing.LightGreen)
	s.FillRect(drawing.NewRect(10, 65, width-10, 67), drawing.LightCyan)
	s.FillRect(drawing.NewRect(10, 70, width-10, 72), drawing.LightYellow)

	// Draw a vertical separator.
	s.Vline(drawing.NewPoint(midX, 85), drawing.NewPoint(midX, height-10), drawing.Blue)

	// Draw sample circles.
	s.Circle(midX/2, ((height-85)/2)+85, (midX/2)-10, drawing.LightRed)

	// Draw sample filled circle.
	s.FillCircle((midX/2)+midX, ((height-85)/2)+85, (midX/2)-10, drawing.Red)

	// Show some standard (fast) lines.
	s.Line(drawing.NewPoint(10, 85), drawing.NewPoint(midX-5, height-10), drawing.Cyan)
	s.Line(drawing.NewPoint(midX-5, 85), drawing.NewPoint(10, height-10), drawing.Green)

	// Show some anti-aliased lines.
	s.LineA(drawing.NewPoint(midX+5, 85), drawing.NewPoint(width-10, height-10), drawing.Cyan)
	s.LineA(drawing.NewPoint(width-10, 85), drawing.NewPoint(midX+5, height-10), drawing.Green)

	// Draw a color gradient.
	cg := s.GetColorGradient(drawing.LightRed, drawing.LightGreen)
	for x := 0; x <= 100; x++ {
		s.Vline(drawing.NewPoint(x+(midX/4), 90), drawing.NewPoint(x+(midX/4), 100), cg[x])
	}

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
