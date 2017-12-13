package drawing

import (
	"image"
	"image/color"
)

// NewSurface ... Returns a new drawing surface.
func NewSurface(width, height int, background color.NRGBA) Surface {
	b := NewRect(0, 0, width-1, height-1)
	s := Surface{
		Bounds:     b,
		Image:      image.NewNRGBA(image.Rect(0, 0, width, height)),
		Width:      width,
		Height:     height,
		Background: background,
	}

	s.Clear()
	return s
}
