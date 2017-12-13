package drawing

import (
	"image/color"
)

// Clear ... Clears the entire surface.
func (s *Surface) Clear() {
	s.FillRect(s.Bounds, s.Background)
}

// Plot ... Places a point on the surface.
func (s *Surface) Plot(x, y int, c color.NRGBA) {
	s.Image.Set(x, y, c)
}

// Hline ... Draws a horizontal line, ignoring any Y deviation.
func (s *Surface) Hline(start, end Point, c color.NRGBA) {
	x1 := start.X
	x2 := end.X
	y := start.Y
	for x := x1; x <= x2; x++ {
		s.Plot(x, y, c)
	}
}

// Vline ... Draws a vertical line, ignoring any X deviation.
func (s *Surface) Vline(start, end Point, c color.NRGBA) {
	x := start.X
	y1 := start.Y
	y2 := end.Y
	for y := y1; y <= y2; y++ {
		s.Plot(x, y, c)
	}
}

// DrawRect ... Draws a rectangular outline.
func (s *Surface) DrawRect(rect Rect, c color.NRGBA) {
	s.Hline(rect.TopLeft, rect.TopRight, c)       // top
	s.Hline(rect.BottomLeft, rect.BottomRight, c) // bottom
	s.Vline(rect.TopLeft, rect.BottomLeft, c)     // left
	s.Vline(rect.TopRight, rect.BottomRight, c)   // right
}

// FillRect ... Draws a solid rectangle.
func (s *Surface) FillRect(rect Rect, c color.NRGBA) {
	for offs := 0; offs <= rect.Height; offs++ {
		start := NewPoint(rect.TopLeft.X, rect.TopLeft.Y+offs)
		end := NewPoint(rect.TopRight.X, rect.TopLeft.Y+offs)
		s.Hline(start, end, c)
	}
}
