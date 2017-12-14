package drawing

import (
	"image/color"
)

// GetColorGradient ... Returns colors mapped by integer percentage from start to end.
func (s *Surface) GetColorGradient(start, end color.RGBA) [101]color.RGBA {
	// Pre-convert to avoid conversion creep.
	sR := float64(start.R)
	sG := float64(start.G)
	sB := float64(start.B)
	eR := float64(end.R)
	eG := float64(end.G)
	eB := float64(end.B)

	// Calculate single percentage adjustment of RGB values.
	pcR := (eR - sR) / 100.0
	pcG := (eG - sG) / 100.0
	pcB := (eB - sB) / 100.0

	// Incrementally apply that movement for 1..99 percent.
	var steps [101]color.RGBA
	for i := 1; i < 100; i++ {
		r := sR + (pcR * float64(i))
		g := sG + (pcG * float64(i))
		b := sB + (pcB * float64(i))
		steps[i] = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: start.A}
	}

	// Enforce start and end.
	steps[0] = start
	steps[100] = end
	return steps
}
