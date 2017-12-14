package drawing

import (
	"image/color"
	"math"
	"time"
)

// Clear ... Clears the entire surface.
func (s *Surface) Clear(c color.RGBA) {
	defer s.trackDuration("Clear", time.Now())

	s.Background = c
	s.FillRect(s.Bounds, s.Background)
}

// Plot ... Places a point on the surface.
func (s *Surface) Plot(x, y int, c color.RGBA) {
	s.Image.Set(x, y, c)
}

// PlotA ... Plots, with a naive antialiased weighted (0..1) second pixel.
func (s *Surface) PlotA(x1, y1, x2, y2 float64, c color.RGBA, weight float64) {
	c2 := color.RGBA{
		R: uint8(float64(c.R) * weight),
		G: uint8(float64(c.G) * weight),
		B: uint8(float64(c.B) * weight),
		A: c.A,
	}
	s.Image.Set(int(x1), int(y1), c)
	s.Image.Set(int(x2), int(y2), c2)
}

// PlotPoint ... Places a point on the surface.
func (s *Surface) PlotPoint(point Point, c color.RGBA) {
	s.Image.Set(point.X, point.Y, c)
}

// Hline ... Draws a horizontal line, ignoring any Y deviation.
func (s *Surface) Hline(start, end Point, c color.RGBA) {
	x1 := start.X
	x2 := end.X
	y := start.Y
	for x := x1; x <= x2; x++ {
		s.Plot(x, y, c)
	}
}

// Vline ... Draws a vertical line, ignoring any X deviation.
func (s *Surface) Vline(start, end Point, c color.RGBA) {
	x := start.X
	y1 := start.Y
	y2 := end.Y
	for y := y1; y <= y2; y++ {
		s.Plot(x, y, c)
	}
}

// Line ... Draws a line (using Bresenham, no anti-aliasing).
func (s *Surface) Line(start, end Point, c color.RGBA) {
	p := NewPoint(start.X, start.Y)

	// Calculate movement and inherent deviation.
	dX := intAbs(end.X - start.X)
	dY := intAbs(end.Y - start.Y)
	err := dX - dY

	// Calculate step in each axis.
	stepX, stepY := -1, -1
	if start.X < end.X {
		stepX = 1
	}
	if start.Y < end.Y {
		stepY = 1
	}

	// Iterate over the line's intermediate pixels.
	deviation := 0
	for {
		s.PlotPoint(p, c)
		if p.X == end.X && p.Y == end.Y {
			return
		}
		deviation = 2 * err
		if deviation > -dY {
			err, p.X = err-dY, p.X+stepX
		}
		if deviation < dX {
			err, p.Y = err+dX, p.Y+stepY
		}
	}
}

// LineA ... Draws an arbitrary naive antialiased line (Xiaolin Wu's algorithm).
func (s *Surface) LineA(start, end Point, c color.RGBA) {
	defer s.trackDuration("LineA", time.Now())

	x0 := float64(start.X)
	y0 := float64(start.Y)
	x1 := float64(end.X)
	y1 := float64(end.Y)

	// Quadrant switch based on if the line more horizontal than vertical.
	tall := math.Abs(y1-y0) > math.Abs(x1-x0)
	if tall {
		x0, y0 = swapFloats(x0, y0)
		x1, y1 = swapFloats(x1, y1)
	}

	// Quadrant switch based on the line's horizontal direction.
	if x0 > x1 {
		x0, x1 = swapFloats(x0, x1)
		y0, y1 = swapFloats(y0, y1)
	}

	// Calculate the slope.
	dx := x1 - x0
	dy := y1 - y0
	gradient := 1.0
	if dx != 0.0 {
		gradient = dy / dx
	}

	// First endpoint.
	xend := round(x0)
	yend := y0 + gradient*(xend-x0)
	xgap := inverseMantissa(x0 + 0.5)
	xpxl1 := xend
	ypxl1 := math.Floor(yend)
	if tall {
		s.PlotA(ypxl1, xpxl1, ypxl1+1, xpxl1, c, mantissa(yend)*xgap)
	} else {
		s.PlotA(xpxl1, ypxl1, xpxl1, ypxl1+1, c, mantissa(yend)*xgap)
	}
	intersect := yend + gradient

	// Second endpoint.
	xend = round(x1)
	yend = y1 + gradient*(xend-x1)
	xgap = mantissa(x1 + 0.5)
	xpxl2 := xend //this will be used in the main loop
	ypxl2 := math.Floor(yend)
	if tall {
		s.PlotA(ypxl2, xpxl2, ypxl2+1, xpxl2, c, mantissa(yend)*xgap)
	} else {
		s.PlotA(xpxl2, ypxl2, xpxl2, ypxl2+1, c, mantissa(yend)*xgap)
	}

	// Line body.
	if tall {
		for x := xpxl1 + 1; x <= xpxl2-1; x++ {
			s.PlotA(math.Floor(intersect), x, math.Floor(intersect)+1, x, c, mantissa(intersect))
			intersect = intersect + gradient
		}
	} else {
		for x := xpxl1 + 1; x <= xpxl2-1; x++ {
			s.PlotA(x, math.Floor(intersect), x, math.Floor(intersect)+1, c, mantissa(intersect))
			intersect = intersect + gradient
		}
	}
}

// DrawRect ... Draws a rectangular outline.
func (s *Surface) DrawRect(rect Rect, c color.RGBA) {
	defer s.trackDuration("DrawRect", time.Now())

	s.Hline(rect.TopLeft, rect.TopRight, c)       // top
	s.Hline(rect.BottomLeft, rect.BottomRight, c) // bottom
	s.Vline(rect.TopLeft, rect.BottomLeft, c)     // left
	s.Vline(rect.TopRight, rect.BottomRight, c)   // right
}

// FillRect ... Draws a solid rectangle.
func (s *Surface) FillRect(rect Rect, c color.RGBA) {
	defer s.trackDuration("FillRect", time.Now())

	for offs := 0; offs <= rect.Height; offs++ {
		start := NewPoint(rect.TopLeft.X, rect.TopLeft.Y+offs)
		end := NewPoint(rect.TopRight.X, rect.TopLeft.Y+offs)
		s.Hline(start, end, c)
	}
}

// Circle ... Draws a fast circle.
func (s *Surface) Circle(centreX, centreY, radius int, c color.RGBA) {
	defer s.trackDuration("Circle", time.Now())

	x := radius
	y := 0
	err := 0

	for {
		// Plot all octants at the same time.
		s.Plot(centreX+x, centreY+y, c)
		s.Plot(centreX+x, centreY-y, c)
		s.Plot(centreX-x, centreY-y, c)
		s.Plot(centreX-x, centreY+y, c)
		s.Plot(centreX+y, centreY+x, c)
		s.Plot(centreX+y, centreY-x, c)
		s.Plot(centreX-y, centreY-x, c)
		s.Plot(centreX-y, centreY+x, c)

		if x <= y {
			return
		}

		// Move on.
		err += 2*y + 1
		y++
		if err > x {
			err += 1 - 2*x
			x--
		}
	}
}

// FillCircle ... Draws a filled circle.
func (s *Surface) FillCircle(centreX, centreY, radius int, c color.RGBA) {
	defer s.trackDuration("FillCircle", time.Now())

	r2 := radius * radius
	for cy := -radius; cy <= radius; cy++ {
		cx := (int)(math.Sqrt(float64(r2-cy*cy)) + 0.5)
		cyy := cy + centreY

		s.Line(NewPoint(centreX-cx, cyy), NewPoint(centreX+cx, cyy), c)
	}
}
