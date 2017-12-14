package drawing

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"time"
)

// Surface ... Drawing surface with various primitives and helpers.
type Surface struct {
	log TimingLogFunc

	Image         *image.RGBA
	Bounds        Rect
	Width, Height int
	Background    color.RGBA
}

// WritePNG ... Writes a PNG stream to the output (eg *io.File).
func (s *Surface) WritePNG(w io.Writer) error {
	return png.Encode(w, s.Image)
}

func (s *Surface) trackDuration(operation string, started time.Time) {
	sofar := time.Since(started)
	s.log(operation, sofar)
}
