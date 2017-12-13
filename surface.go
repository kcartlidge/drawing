package drawing

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

// Surface ... Drawing surface with various primitives and helpers.
type Surface struct {
	Image         *image.NRGBA
	Bounds        Rect
	Width, Height int
	Background    color.NRGBA
}

// WritePNG ... Writes a PNG stream to the output (eg *io.File).
func (s *Surface) WritePNG(w io.Writer) error {
	return png.Encode(w, s.Image)
}
